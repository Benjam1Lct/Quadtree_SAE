/*
character est le paquet qui permet de définir la manière dont le personnage se
comporte. Pour le moment ceci prend en compte :
  - l'affichage
  - les contrôles
  - l'animation

# TYPES

	type Character struct {
	        X, Y int

	        // Has unexported fields.
	}

	Character définit les charactéristiques du personnage. Pour le moment seules
	les coordonnées absolues de l'endroit où il se trouve sont exportées,
	mais vous pourrez ajouter des choses au besoins lors de votre développement.

	Les champs non exportés sont :
	  - orientation : l'orientation du personnage (haut, bas, gauche, droite).
	  - animationStep : l'étape d'animation (-1 ou 1, représentant l'animation
	    d'un pas à gauche ou à droite).
	  - xInc, yInc : les incréments en X et Y à réaliser après la prochaine
	    animation.
	  - moving : l'information de si une animation est en cours ou pas.
	  - shift : la position actuelle en pixels du personnage relativement à ses
	    coordonnées absolues.
	  - animationFrameCount : le nombre d'appels à update (ou de 1/60 de
	    seconde) qui ont eu lieu depuis la dernière étape d'animation.
	  - tp : objet de type Teleport représentant le téléporteur associé au
	    personnage.
	  - animationFlag : le drapeau d'activation/désactivation de l'animation.
	  - animationCounter : le compteur d'animations.

func (c Character) Draw(screen *ebiten.Image, camX, camY int)

	Draw permet d'afficher le personnage dans une *ebiten.Image (en pratique,
	celle qui représente la fenêtre de jeu) en fonction des caractéristiques du
	personnage (position, orientation, étape d'animation, etc) et de la position
	de la caméra (le personnage est affiché relativement à la caméra).

func (c *Character) Init()

	Init met en place un personnage. Pour le moment cela consiste simplement
	à initialiser une variable responsable de définir l'étape d'animation
	courante. initialise egalement les points de teleportation et le sprite qui
	y est lié

func (c *Character) Update(blocking [4]bool)

	Update met à jour la position du personnage, son orientation et son étape
	d'animation (si nécessaire) à chaque pas de temps, c'est-à-dire tous les
	1/60 secondes.

	type Teleport struct {
	        Tpress bool
	        // Has unexported fields.
	}

	Teleport est la structure permettant la création de portail où on a:
	  - enterX: l'abscisse du premier portail créé
	  - enterY: l'ordonné du premier portail créé
	  - endX: l'abscisse du dernier portail créé
	  - endY: l'ordonné du dernier portail créé
	  - Portal: un booléen nous indiquant si nous créons un portails de sorti ou
	    bien un portail d'entré
	  - onPortal: un booléen nous valant true tant que le personnage reste sur
	    un un portail après sa téléportation
	  - Tpress: un booléen valant true tant que la touche T est pressé

func Init_Teleport() Teleport

	cette fonction permet l'initialisation de la variable de type Teleport que
	l'on utilise dans update
*/
package character
