package main

import (
	"context"

	"github.com/johanych/adsimulation/pkg/behaviours/ransom_encrypt"
	"github.com/johanych/adsimulation/pkg/behaviours/ransom_note"
	"github.com/johanych/adsimulation/pkg/behaviours/ransom_wallpaper"
	"github.com/johanych/adsimulation/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransom_encrypt.NewRansomEncrypt(),
		ransom_note.NewRansomNote(),
		ransom_wallpaper.NewRansomWallpaper(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
