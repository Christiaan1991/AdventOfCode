package robot

type Robot struct {
	Direction string
	Position  [2]int
}

func (r *Robot) PaintPanel(num int) int {
	if num == 0 { //black
		return 0
	}
	return 1 //white
}

func (r *Robot) GetDirection() string {
	return (*r).Direction
}

func (r *Robot) GoLeft() {
	if r.Direction == "NORTH" {
		r.Direction = "WEST"
		return
	}

	if r.Direction == "WEST" {
		r.Direction = "SOUTH"
		return
	}

	if r.Direction == "SOUTH" {
		r.Direction = "EAST"
		return
	}

	if r.Direction == "EAST" {
		r.Direction = "NORTH"
		return
	}
}

func (r *Robot) GoRight() {
	if r.Direction == "NORTH" {
		r.Direction = "EAST"
		return
	}

	if r.Direction == "WEST" {
		r.Direction = "NORTH"
		return
	}

	if r.Direction == "SOUTH" {
		r.Direction = "WEST"
		return
	}

	if r.Direction == "EAST" {
		r.Direction = "SOUTH"
		return
	}
}

func (r *Robot) Move() {

	if r.Direction == "NORTH" {
		r.Position[1]++
	}

	if r.Direction == "WEST" {
		r.Position[0]--
	}

	if r.Direction == "SOUTH" {
		r.Position[1]--
	}

	if r.Direction == "EAST" {
		r.Position[0]++
	}
}

func (r *Robot) TurnRobot(num int) {
	if num == 1 {
		r.GoLeft()
		return
	}
	r.GoRight()
}
