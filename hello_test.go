package main

import "testing"

func TestHelloLn(t *testing.T) {
	resultHelloFn := hello()
	if len(resultHelloFn) != 5 {
		t.Fatal("неправильная длина вернувшейся строки")
	}

}
func TestHelloData(t *testing.T) {
	resultHelloFn := hello()
	if resultHelloFn != "Hello" {
		t.Fatal("это не hello")
	}
}
