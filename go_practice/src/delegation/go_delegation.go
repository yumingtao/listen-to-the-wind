package delegation

import "fmt"

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
	X int
}

func (label Label) Paint()  {
	fmt.Printf("[%p] - Label.Paint(%q)\n", &label, label.Text)
}
