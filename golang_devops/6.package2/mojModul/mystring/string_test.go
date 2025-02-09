package mystring

import "testing"

func TestReverse(t *testing.T){
	result := Reverse("Golang")
	expected := "gnaloG"
	if result != expected{
		t.Errorf("golang odwrotnie to %s, a zwrócono %s", expected, result )
	}
}