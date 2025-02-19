package main

import (
	"context"

	discoveryperipheral "github.com/johanych/adsimulation/pkg/behaviours/discovery_peripheral"
	discoveryprocess "github.com/johanych/adsimulation/pkg/behaviours/discovery_process"
	discoverysoftware "github.com/johanych/adsimulation/pkg/behaviours/discovery_software"
	"github.com/johanych/adsimulation/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		discoveryprocess.NewDiscoveryProcess(),
		discoveryperipheral.NewDiscoveryPeripheral(),
		discoverysoftware.NewDiscoverySoftware(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
