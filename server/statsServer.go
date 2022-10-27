package server

import (
	"log"
	"sync"
	"time"

	pb "github.com/achelabov/systat/proto"
	"github.com/achelabov/systat/server/widgets"
	"google.golang.org/protobuf/types/known/emptypb"
)

type statsServer struct {
	pb.UnimplementedStatsServiceServer
}

func (s *statsServer) GetStats(in *emptypb.Empty, srv pb.StatsService_GetStatsServer) error {
	var wg sync.WaitGroup

	batts := make(chan *pb.BatteriesResponse)
	cpus := make(chan *pb.CpusResponse)

	cancel := make(chan struct{})
	defer close(cancel)

	wg.Add(2)
	recieve := func() {
		defer close(batts)
		defer close(cpus)

		defer wg.Done()

		for {
			select {
			case batts <- <-GetBatteries(cancel):
			case cpus <- <-GetCpus(cancel):
			}
		}
	}

	ticker := time.NewTicker(time.Second)
	send := func() {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				statsResp := new(pb.StatsResponse)
				statsResp.Batteries = <-batts
				statsResp.Cpus = <-cpus

				if err := srv.Send(statsResp); err != nil {
					log.Printf("send error %v", err)
				}
			case <-cancel:
				return
			}
		}
	}

	go recieve()
	time.Sleep(time.Second)
	go send()

	wg.Wait()
	return nil
}

func GetBatteries(cancel <-chan struct{}) <-chan *pb.BatteriesResponse {
	battWidget := *widgets.NewBatteryWidget()
	out := make(chan *pb.BatteriesResponse)

	for i := 0; i < 5; i++ {
		go func() {

			battsResp := &pb.BatteriesResponse{
				Batteries: make([]*pb.Battery, battWidget.BattsCount()),
			}
			for i := 0; i < battWidget.BattsCount(); i++ {
				battsResp.Batteries[i] = new(pb.Battery)
			}

			for batts := range battWidget.GetBatteries(cancel) {
				for i, v := range batts {
					battsResp.Batteries[i].BatteryLoad = v.BatteryLoad
					battsResp.Batteries[i].State = v.State
				}
				out <- battsResp
			}
		}()
	}

	return out
}

func GetCpus(cancel <-chan struct{}) <-chan *pb.CpusResponse {
	cpuWidget := *widgets.NewCpuWidget()
	out := make(chan *pb.CpusResponse)

	for i := 0; i < 5; i++ {
		go func() {
			cpusResp := &pb.CpusResponse{
				Cpus: make([]*pb.Cpu, cpuWidget.CpusCount()),
			}
			for i := 0; i < cpuWidget.CpusCount(); i++ {
				cpusResp.Cpus[i] = new(pb.Cpu)
			}

			for cpus := range cpuWidget.GetCpus(cancel) {
				for i, v := range cpus {
					cpusResp.Cpus[i].CpuLoad = v.CpuLoad
				}
				cpusResp.AverageLoad = <-cpuWidget.GetAverageLoad(cancel)

				out <- cpusResp
			}
		}()
	}

	return out
}
