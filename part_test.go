package labo

import (
	"fmt"
	"testing"
)

func TestPartAmount(t *testing.T) {

	s := []string{
		"One Short Straps",
		"Two Short Strap",
		"Three Short Strap",
		"Four Short Strings",
		"Five Short Strings",
		"Six Short Sheet",
		"Seven Short Sheet",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve Grommets",
		"Thirteen"}

	for _, s := range s {
		fmt.Println(getPartAmount(s))
	}
}

func TestPartColor(t *testing.T) {

	s := []string{
		"Blue",
		"(Blue)",
		"Gray",
		"One Orange",
		"Red",
		"Yellow"}

	for _, s := range s {
		fmt.Println(getPartColor(s))
	}
}
