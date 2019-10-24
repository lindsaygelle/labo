package labo_test

import (
	"net/http"
	"testing"

	"github.com/gellel/labo"
)

func Test(t *testing.T) {

	for _, l := range labo.GetAll(labo.KitsID) {

		k := labo.GetKit(l)

		if k.StatusCode == http.StatusOK {
			/*
				fmt.Println(string(labo.Marshal(l)))

				fmt.Println(string(labo.MarshalKit(k)))

				fmt.Println("-")
				fmt.Println("")
				fmt.Println("")
			*/
		}
	}

}
