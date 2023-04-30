package main

import "testing"

func TestSaudacao(t *testing.T) {
	resultado := saudacao()
	experado := "Hello World, William"
	if resultado != experado {
		t.Errorf("resultado '%s', experado '%s'", resultado, experado)
	}
}
