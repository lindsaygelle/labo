package labo_test

import (
	"fmt"
	"testing"

	"github.com/gellel/labo"
)

func TestLabo(t *testing.T) {
	kit, _ := labo.GetKit(labo.RobotKitURL)
	fmt.Println(kit.BoxImage)
	fmt.Println(kit.Href)
	fmt.Println(kit.Materials)
	fmt.Println(kit.Name)
	fmt.Println(kit.Overview)
	fmt.Println(kit.Price)
	fmt.Println(kit.Projects)
	fmt.Println(kit.Retailers)
	fmt.Println(kit.Software)
	fmt.Println(kit.ToyCons)
}
