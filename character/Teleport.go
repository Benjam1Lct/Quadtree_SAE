package character

// Teleport est la structure permettant la création de portail où on a:
// 	- enterX: l'abscisse du premier portail créé
// 	- enterY: l'ordonné du premier portail créé
// 	- endX: l'abscisse du dernier portail créé
// 	- endY: l'ordonné du dernier portail créé
// 	- Portal: un booléen nous indiquant si nous créons un portails de sorti
// 	  ou bien un portail d'entré
// 	- onPortal: un booléen nous valant true tant que le personnage reste sur un
// 	  un portail après sa téléportation
// 	- Tpress: un booléen valant true tant que la touche T est pressé
type Teleport struct {
	enterX, enterY   int
	endX, endY       int
	Portal, onPortal bool
	Tpress           bool
}

// cette fonction permet l'initialisation de la variable de type Teleport que l'on utilise
// dans update
func Init_Teleport() Teleport {
	// Initialisation de la variable de type Teleport que l'on utilise dans update
	Tp := Teleport{
		enterX:   -1,
		enterY:   -1,
		endX:     -1,
		endY:     -1,
		Portal:   true,
		onPortal: false,
		Tpress:   true,
	}
	return Tp
}

// méthode permettant la création des portails
// elle prend en paramètre
// 	- characterX: la position en abscisse du personnage
// 	- characterY: la position en ordonné du personnage
func (t *Teleport) create_teleport(characterX, characterY int) {
	if t.Portal {
		// création d'un portail d'entré
		t.enterX = characterX
		t.enterY = characterY
	} else {
		// création d'un portail de sorti
		t.endX = characterX
		t.endY = characterY
	}
	t.Portal = !t.Portal
	t.onPortal = true
}
