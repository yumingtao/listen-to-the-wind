package delegation

import "fmt"

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}
func (listBox ListBox) Paint() {
	fmt.Printf("[%p] - ListBox.Paint(%q)\n",
		&listBox, listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("[%p] - ListBox.Click()\n", &listBox)
}
