package character

type Teleport struct {
	enterX, enterY   int
	endX, endY       int
	Portal, onPortal bool
	Tpress           bool
}

func Init_Teleport() Teleport {
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

func (t *Teleport) create_teleport(characterX, characterY int) {
	if t.Portal {
		t.enterX = characterX
		t.enterY = characterY
	} else {
		t.endX = characterX
		t.endY = characterY
	}
	t.Portal = !t.Portal
	t.onPortal = true
}
