/*
floor est le paquet qui gère la configuration du sol sur lequel le personnage
se déplace. Sa principale fonction est de fournir, chaque 1/60 de seconde
et en fonction de la position absolue de la caméra, une structure de données
représentant les cases de terrain visibles à l'écran.

# TYPES

	type Floor struct {
	        // Has unexported fields.
	}

	Floor représente les données du terrain. Pour le moment aucun champs n'est
	exporté.

	  - content : partie du terrain qui doit être affichée à l'écran
	  - fullContent : totalité du terrain (utilisé seulement avec le type
	    d'affichage du terrain "fromFileFloor")
	  - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
	    avec le type d'affichage du terrain "quadtreeFloor")

func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int)
(blocking [4]bool)

	Blocking retourne, étant donnée la position du personnage, un tableau de
	booléen indiquant si les cases au dessus (0), à droite (1), au dessous (2)
	et à gauche (3) du personnage sont bloquantes.

func (f Floor) Draw(screen *ebiten.Image)

	Draw affiche dans une image (en général, celle qui représente l'écran), la
	partie du sol qui est visible (qui doit avoir été calculée avec Get avant).

func (f *Floor) Init()

	Init initialise les structures de données internes de f

func (f *Floor) Update(camXPos, camYPos int)

	Update se charge de stocker dans la structure interne (un tableau) de f une
	représentation de la partie visible du terrain à partir des coordonnées
	absolues de la case sur laquelle se situe la caméra.

	On aurait pu se passer de cette fonction et tout faire dans Draw. Mais cela
	permet de découpler le calcul de l'affichage.

# CONSTANTS

const (

	SphereWorld int

)

	types d'affichage du terrain disponibles

# TYPES

	type Floor struct {
	        // Has unexported fields.
	}

	Floor représente les données du terrain. Pour le moment aucun champs n'est
	exporté.

	  - content : partie du terrain qui doit être affichée à l'écran
	  - fullContent : totalité du terrain (utilisé seulement avec le type
	    d'affichage du terrain "fromFileFloor")
	  - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
	    avec le type d'affichage du terrain "quadtreeFloor")

func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool)

	Blocking retourne, étant donnée la position du personnage, un tableau de
	booléen indiquant si les cases au dessus (0), à droite (1), au dessous (2)
	et à gauche (3) du personnage sont bloquantes.

func (f Floor) Draw(screen *ebiten.Image)

	Draw affiche dans une image (en général, celle qui représente l'écran), la
	partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
	du au changement des assets des textures du sol il s'agit maintenant d'un
	cadrillage de 32x32 donc il est necessere d'effectuer des operations afin de
	revenenir a une tuile simple

func (f *Floor) Init()

	Init initialise les structures de données internes de f Cette fonction
	initialise le contenu d'un étage (Floor) en fonction des paramètres de
	configuration. Possibilité de créer en aléatoire ou à partir d'un fichier
	avec ou sans quadtree

func (f *Floor) Update(camXPos, camYPos int)

	Update se charge de stocker dans la structure interne (un tableau) de f une
	représentation de la partie visible du terrain à partir des coordonnées
	absolues de la case sur laquelle se situe la caméra.

	On aurait pu se passer de cette fonction et tout faire dans Draw. Mais cela
	permet de découpler le calcul de l'affichage.
*/
package floor
