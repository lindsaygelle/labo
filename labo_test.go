package labo_test

import (
	"fmt"
	"testing"

	"github.com/gellel/labo"
)

func TestLabo(t *testing.T) {

	robotKit, err := labo.NewRobotKit()

	if err != nil {
		panic(err)
	}
	fmt.Println(robotKit.BoxImage.Src)
}
