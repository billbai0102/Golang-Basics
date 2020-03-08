package main

import "fmt"

type rect struct {
	width int
	height int
}

func (r rect) getArea() int{
	return r.width * r.height
}

func (r *rect) getPerim() int{
	return 2*r.height + 2*r.height
}

func main()  {
	r := rect {width: 25, height: 10}
	fmt.Println("area:", r.getArea())
	fmt.Println("perim:", r.getPerim())

	rptr := &r
	fmt.Println("area:", rptr.getArea())
	fmt.Println("perim:", rptr.getPerim())
}
