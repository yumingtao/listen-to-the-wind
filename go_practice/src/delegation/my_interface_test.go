package delegation

import (
	"fmt"
	"testing"
)

func TestButtonAndListBoxInterface(t *testing.T) {
	label := Label{Widget{10, 10}, "State", 100}
	button1 := Button{Label{Widget{10, 70}, "OK", 10}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	fmt.Println()
	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	fmt.Println()

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
	}
}