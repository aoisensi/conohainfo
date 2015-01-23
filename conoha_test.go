package conohainfo

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestConoHa(t *testing.T) {
	list, err := GetList()
	if err != nil {
		t.Fatal(err)
	}
	body, err := list[0].GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	pp.Println(body)
}
