package server

import (
	"log"
	"sync"

	pb "github.com/achelabov/systat/proto"
	"github.com/achelabov/systat/server/widgets"
	"google.golang.org/protobuf/types/known/emptypb"
)

type statsServer struct {
	pb.UnimplementedStatsServer
}

func (s *statsServer) GetBatteries(in *emptypb.Empty, srv pb.Stats_GetBatteriesServer) error {
	var wg sync.WaitGroup
	battWidget := *widgets.NewBatteryWidget()

	cancel := make(chan struct{})
	defer close(cancel)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

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
				if err := srv.Send(battsResp); err != nil {
					log.Printf("send error %v", err)
				}
			}
		}()
	}

	wg.Wait()
	return nil
}

func (s *statsServer) GetCpus(in *emptypb.Empty, srv pb.Stats_GetCpusServer) error {
	var wg sync.WaitGroup
	cpuWidget := *widgets.NewCpuWidget()

	cancel := make(chan struct{})
	defer close(cancel)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

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

				if err := srv.Send(cpusResp); err != nil {
					log.Printf("send error %v", err)
				}
			}
		}()
	}
	wg.Wait()
	return nil
}
