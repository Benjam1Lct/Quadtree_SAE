package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/game"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	_ "image/png"
)

// Menu object used by ebiten
type menu struct {
	ui *ebitenui.UI
	// Add other menu-specific fields if needed
}

type pageContainer struct {
	widget    widget.PreferredSizeLocateableWidget
	titleText *widget.Text
	flipBook  *widget.FlipBook
}

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")
	flag.Parse()

	configuration.Load(configFileName)
	assets.Load()

	// Ouvrir le fichier CSV
	file, err := os.Open("config.csv")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	// Créer un lecteur CSV
	reader := csv.NewReader(file)

	// Lire la première ligne du fichier CSV (entête)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'en-tête CSV:", err)
		return
	}

	// Vérifier si "Restart" est présent dans l'en-tête
	restartIndex := -1
	for i, header := range headers {
		if header == "Restart" {
			restartIndex = i
			break
		}
	}

	if restartIndex == -1 {
		fmt.Println("La colonne 'Restart' n'a pas été trouvée dans l'en-tête.")
		return
	}

	// Lire la valeur de "Restart"
	record, err := reader.Read()
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la valeur 'Restart' dans le CSV:", err)
		return
	}

	// Convertir la valeur en booléen
	println(record[restartIndex])

	if record[restartIndex] == "false" && configuration.Global.StartingMenu {
		ebiten.SetWindowSize(720, 480)
		ebiten.SetWindowTitle("Game Quadtree")
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
		ebiten.SetScreenClearedEveryFrame(false)
		ebiten.SetVsyncEnabled(false)

		ui, closeUI, err := createUI()
		if err != nil {
			log.Fatal(err)
		}

		defer closeUI()

		game := menu{
			ui: ui,
		}

		err = ebiten.RunGame(&game)
		if err != nil {
			log.Print(err)
		}
	} else {
		// Ouvrir le fichier CSV en lecture
		file, err := os.OpenFile("config.csv", os.O_RDWR, os.ModePerm)
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier en lecture/écriture :", err)
			return
		}
		defer file.Close()

		// Créer un lecteur CSV
		reader := csv.NewReader(file)

		// Lire l'en-tête
		headers, err := reader.Read()
		if err != nil {
			fmt.Println("Erreur lors de la lecture de l'en-tête CSV :", err)
			return
		}

		// Trouver l'index de la colonne "Restart"
		restartIndex := -1
		for i, header := range headers {
			if header == "Restart" {
				restartIndex = i
				break
			}
		}

		if restartIndex == -1 {
			fmt.Println("La colonne 'Restart' n'a pas été trouvée dans l'en-tête.")
			return
		}

		// Lire la ligne de données actuelle
		record, err := reader.Read()
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la ligne de données actuelle :", err)
			return
		}

		// Modifier la valeur de "Restart"
		record[restartIndex] = "false"

		// Déplacer le curseur au début du fichier pour réécrire les données
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Println("Erreur lors du déplacement du curseur au début du fichier :", err)
			return
		}

		// Créer un écrivain CSV pour écrire les modifications
		writer := csv.NewWriter(file)

		// Écrire l'en-tête modifié
		err = writer.Write(headers)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture de l'en-tête modifié :", err)
			return
		}

		// Écrire la ligne de données modifiée
		err = writer.Write(record)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture de la ligne de données modifiée :", err)
			return
		}

		// Vider le tampon pour s'assurer que toutes les données sont écrites dans le fichier
		writer.Flush()

		// Vérifier les erreurs d'écriture
		err = writer.Error()
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}

		fmt.Println("La valeur de 'Restart' a été modifiée avec succès.")

		runGame()
	}
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

func createUI() (*ebitenui.UI, func(), error) {
	res, err := newUIResources()
	if err != nil {
		return nil, nil, err
	}

	//This creates the root container for this UI.
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			// It is using a GridLayout with a single column
			widget.GridLayoutOpts.Columns(1),
			// It uses the Stretch parameter to define how the rows will be layed out.
			// - a fixed sized header
			// - a content row that stretches to fill all remaining space
			// - a fixed sized footer
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			// Padding defines how much space to put around the outside of the grid.
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			// Spacing defines how much space to put between each column and row
			widget.GridLayoutOpts.Spacing(0, 20))),
		widget.ContainerOpts.BackgroundImage(res.background))

	rootContainer.AddChild(headerContainer(res))

	var ui *ebitenui.UI
	rootContainer.AddChild(demoContainer(res, func() *ebitenui.UI {
		return ui
	}))

	footerContainer := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Padding(widget.Insets{
			Left:  25,
			Right: 25,
		}),
	)))
	rootContainer.AddChild(footerContainer)

	footerContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("https://gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-08_lecomte-benjamin_lerouley-clement/-/tree/main", res.text.smallFace, res.text.disabledColor)))

	ui = &ebitenui.UI{
		Container: rootContainer,
	}

	return ui, func() {
		res.close()
	}, nil
}

func headerContainer(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(15))),
	)

	c.AddChild(header("Game Quadtree", res,
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
	))

	c2 := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Left:  25,
				Right: 25,
			}),
		)),
	)
	c.AddChild(c2)

	c2.AddChild(widget.NewText(
		widget.TextOpts.Text("Starting Menu & Configuration", res.text.face, res.text.idleColor)))

	return c
}

func header(label string, res *uiResources, opts ...widget.ContainerOpt) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(append(opts, []widget.ContainerOpt{
		widget.ContainerOpts.BackgroundImage(res.header.background),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(res.header.padding))),
	}...)...)

	c.AddChild(widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionStart,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
		widget.TextOpts.Text(label, res.header.face, res.header.color),
		widget.TextOpts.Position(widget.TextPositionStart, widget.TextPositionCenter),
	))

	return c
}

func demoContainer(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {

	demoContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Padding(widget.Insets{
				Left:  25,
				Right: 25,
			}),
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(20, 0),
		)))

	pages := []interface{}{
		StartPage(res),
		OptionsPage(res),
	}

	pageContainer := newPageContainer(res)

	pageList := widget.NewList(
		widget.ListOpts.Entries(pages),
		widget.ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(*page).title
		}),
		widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(res.list.image)),
		widget.ListOpts.SliderOpts(
			widget.SliderOpts.Images(res.list.track, res.list.handle),
			widget.SliderOpts.MinHandleSize(res.list.handleSize),
			widget.SliderOpts.TrackPadding(res.list.trackPadding),
		),
		widget.ListOpts.EntryColor(res.list.entry),
		widget.ListOpts.EntryFontFace(res.list.face),
		widget.ListOpts.EntryTextPadding(res.list.entryPadding),
		widget.ListOpts.HideHorizontalSlider(),

		widget.ListOpts.EntrySelectedHandler(func(args *widget.ListEntrySelectedEventArgs) {
			pageContainer.setPage(args.Entry.(*page))
		}))
	demoContainer.AddChild(pageList)

	demoContainer.AddChild(pageContainer.widget)

	pageList.SetSelectedEntry(pages[0])

	return demoContainer
}

func newPageContainer(res *uiResources) *pageContainer {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.panel.image),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(res.panel.padding),
			widget.RowLayoutOpts.Spacing(15))),
	)

	titleText := widget.NewText(
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.TextOpts.Text("", res.text.titleFace, res.text.idleColor))
	c.AddChild(titleText)

	flipBook := widget.NewFlipBook(
		widget.FlipBookOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		}))),
	)
	c.AddChild(flipBook)

	return &pageContainer{
		widget:    c,
		titleText: titleText,
		flipBook:  flipBook,
	}
}

func (p *pageContainer) setPage(page *page) {
	p.titleText.Label = page.title
	p.flipBook.SetPage(page.content)
	p.flipBook.RequestRelayout()
}

func newCheckbox(label string, changedHandler widget.CheckboxChangedHandlerFunc, res *uiResources) *widget.LabeledCheckbox {
	return widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.Spacing(res.checkbox.spacing),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(widget.ButtonOpts.Image(res.checkbox.image)),
			widget.CheckboxOpts.Image(res.checkbox.graphic),
			widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
				if changedHandler != nil {
					changedHandler(args)
				}
			})),
		widget.LabeledCheckboxOpts.LabelOpts(widget.LabelOpts.Text(label, res.label.face, res.label.text)))
}

func newPageContentContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)))
}

func (g *menu) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *menu) Update() error {
	g.ui.Update()
	return nil
}

func (g *menu) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f", ebiten.ActualFPS()))
}
