package main

import "testing"

func TestSaudacao(t *testing.T) {

	validar := func(t *testing.T, resultado, experado string) {
		t.Helper()
		if resultado != experado {
			t.Errorf("resultado '%s', experado '%s'", resultado, experado)
		}
	}

	t.Run("Saudacao sem nome", func(t *testing.T) {
		resultado := saudacao("")
		experado := "Hello World 1"
		validar(t, resultado, experado)
	})

	t.Run("Saudacao com nome", func(t *testing.T) {
		resultado := saudacao("William")
		experado := "Hello World, William 1"
		validar(t, resultado, experado)
	})
}
