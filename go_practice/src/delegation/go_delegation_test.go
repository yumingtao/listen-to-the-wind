package delegation

import (
	"fmt"
	"testing"
)

func TestLabel_Paint(t *testing.T) {
	label := Label{Widget{10, 10}, "State", 100}

	fmt.Printf("X=%d, Y=%d, Text=%s, Widget.x=%d\n", label.X, label.Y, label.Text, label.Widget.X)
	fmt.Println()
	fmt.Printf("%+v\n%v\n", label, label)
	label.Paint()
}
