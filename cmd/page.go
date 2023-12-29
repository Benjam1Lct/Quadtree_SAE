package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"

	"github.com/ebitenui/ebitenui/widget"
)

type page struct {
	title   string
	content widget.PreferredSizeLocateableWidget
}

func StartPage(res *uiResources) *page {
	c := newPageContentContainer()

	b := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Start", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Entered: " + args.Button.Text().Label) }),
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) { fmt.Println("Cursor Exited: " + args.Button.Text().Label) }),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
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
			record[restartIndex] = "true"

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

			// Déclarer la variable cmd à l'extérieur de la goroutine
			var cmd *exec.Cmd

			// Lancer la nouvelle instance dans une goroutine
			go func() {
				cmd := exec.Command("./cmd")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("Erreur lors du lancement du programme :", err)
				}
			}()

			// Attendre que la nouvelle instance se termine
			if err := cmd.Wait(); err != nil {
				fmt.Println("Erreur lors de l'attente de la fin du programme :", err)
			}

			// Terminer le programme actuel
			os.Exit(0)

		}),
	)
	c.AddChild(b)

	return &page{
		title:   "Start",
		content: c,
	}
}

func OptionsPage(res *uiResources) *page {

	c := newPageContentContainer()

	for i := 0; i < 16; i++ {
		cb1 := newCheckbox(fmt.Sprintf("Button %d", i+1), nil, res)
		c.AddChild(cb1)
	}

	return &page{
		title:   "Options",
		content: c,
	}
}
