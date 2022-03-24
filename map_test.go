package go_collection

import (
	"testing"
)

func TestMap(t *testing.T) {
	type foo struct {
		Name string
	}
	data := []foo{
		{Name: "test"},
		{Name: "ben"},
		{Name: "tom"},
	}
	names := NewCollection(data).Map(func(item foo) any {
		return item.Name
	})
	if len(names) != 3 {
		t.Fail()
	}

	if len(Map(data, func(f foo) any {
		return f.Name
	})) != 3 {
		t.Fail()
	}
}
