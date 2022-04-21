package main

import (
	"context"

	"github.com/johanych/adsimulation/pkg/behaviours/bypass_fodhelper"
	"github.com/johanych/adsimulation/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		bypass_fodhelper.NewBypassFodhelper(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
