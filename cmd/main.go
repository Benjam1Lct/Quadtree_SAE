package main

import (
	"flag"
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/game"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")
	flag.Parse()

	configuration.Load(configFileName)
	assets.Load()

	runGame()
}

func runGame() {
	g := &game.Game{}
	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(720, 480)
	ebiten.SetWindowTitle("Game Quadtree")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
