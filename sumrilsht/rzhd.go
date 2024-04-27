package sumrilsht

import (
	"fmt"
)

type BSTrain struct {
	Number string
}

func (t *BSTrain) Drive() {
	fmt.Printf("Drive train %q\n", t.Number)
}

func NewBSTrain(number string) *BSTrain {
	return &BSTrain{Number: number}
}
