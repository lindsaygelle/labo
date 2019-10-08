package labo

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for _, k := range getKits() {
		if len(k.Ref) > 0 {
			fmt.Println(k.Name)
			fmt.Println(k.Ref)
			fmt.Println("-")
		}
	}
}
