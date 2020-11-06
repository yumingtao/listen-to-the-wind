package delegation

import "fmt"

type Button struct {
	Label
}

func NewButton(x, y int, text string) Button {
	return Button{Label{Widget{x, y}, text, x}}
}

func (button Button) Paint() {
	fmt.Printf("[%p] - Button.paint()\n", &button)
}

func (button Button) Click()  {
	fmt.Printf("[%p] - Button.Click()\n", &button)
}
