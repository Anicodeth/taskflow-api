package models

import "testing"

func TestStatusValid(t *testing.T) {
	if !StatusTodo.Valid() {
		t.Fatal("todo should be valid")
	}
	if Status("bogus").Valid() {
		t.Fatal("bogus should be invalid")
	}
}