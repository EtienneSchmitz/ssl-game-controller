package main

import (
	"fmt"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/config"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/gc"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
	"os"
	"path/filepath"
	"strings"
)

var basePath = "cmd/ssl-game-controller-c4-arch"
var targetPath = "doc/"

func main() {
	var cfg = config.DefaultControllerConfig()
	var gameController = gc.NewGameController(cfg)

	var structure = scrap(gameController, basePath+"/scraper.yml")
	err := filepath.Walk(basePath+"/views", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			createPuml(structure, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func scrap(gameController *gc.GameController, scraperConfig string) model.Structure {
	s, err := scraper.NewScraperFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	return s.Scrape(gameController)
}

func createPuml(structure model.Structure, scraperConfig string) {
	v, err := view.NewViewFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	var baseName = strings.TrimSuffix(filepath.Base(scraperConfig), filepath.Ext(scraperConfig))
	outFile, err := os.Create(targetPath + baseName + ".puml")
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(err)
		}
	}(outFile)

	err = v.RenderStructureTo(structure, outFile)
	if err != nil {
		panic(err)
	}
}
