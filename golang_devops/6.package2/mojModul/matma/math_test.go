package matma

import "testing"

func TestAdd(t *testing.T){
	result := Add(2, 5)
	if result != 7{
		t.Error("Oczekiwano 7, a jest: ", result)
	}
}

func TestMultiply(t *testing.T){
	result := Multiply(3, 5)
	if result != 15{
		t.Error("Oczekiwano 15, a jest: ", result)
	}
}