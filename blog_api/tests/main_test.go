package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(5, 3) != 9 {
		t.Error("not passed")
	}
}
