package main

import "fmt"

func saudacao(nome string) string {
	if nome == "" {
		return "Hello World"
	}
	return "Hello World, " + nome
}

func main() {
	fmt.Println(saudacao(""))
}
