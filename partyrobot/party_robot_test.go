package partyrobot

import (
	"testing"
)

type Struct struct {
	name string
	want string
}

type Birthday struct {
	name string
	age  int
	want string
}

func TestWelcome(t *testing.T) {

	tt := Struct{
		name: "Chihiro",
		want: "Welcome to my party, Chihiro!",
	}

	if got := Welcome(tt.name); got != tt.want {
		t.Errorf("Welcome(%s) = %s, want %s", tt.name, got, tt.want)
	}

}
