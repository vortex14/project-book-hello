package main

import "testing"

func TestHello(t *testing.T) {
	resultHelloFn := hello()
	if len(resultHelloFn) != 5 {
		t.Fatal("неправильная длина вернувшейся строки")
	}
	if resultHelloFn != "hello" {
		t.Fatal("это не hello")
	}
}
