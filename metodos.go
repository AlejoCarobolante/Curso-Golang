package main

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
} //Este metodo (area) recibe un dato de tipo *rect

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
