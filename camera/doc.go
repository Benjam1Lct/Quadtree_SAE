/*
camera est le paquet qui prend en charge la gestion d'une caméra abtraite.
L'idée est que la position de la caméra représente toujours le centre de
l'affichage à l'écran.

Pour le moment, la caméra peut soit suivre le personnage soit rester immobile.
On peut choisir entre ces deux modes à l'aide du fichier de configuration du jeu
(fourni avec la paquet quadtree/cmd).

# CONSTANTS

const (

	Static int = iota
	FollowCharacter
	NoVoid
	MovieCamera

)

	types de caméra disponibles

# TYPES

	type Camera struct {
	        X, Y               int
	        MapSizeX, MapSizeY int
	}

	Camera définit les caractéristiques de la caméra. Pour le moment il s'agit
	simplement des coordonnées absolues de l'endroit où elle se trouve mais vous
	pourrez ajouter des choses au besoin lors de votre développement.

func (c *Camera) Init()

	Init initialise la caméra en fonction de la configuration du jeu.

func (c *Camera) Update(characterPosX, characterPosY int)

	Update met à jour la position de la caméra à chaque pas de temps,
	c'est-à-dire tous les 1/60 secondes.
*/
package camera
