package helper

import "testing"

func TestHelloTirta(t *testing.T) {
	result := SayHallo("Tirta")
	if result != "Hello Tirta" {
		panic("Expected 'Hello Tirta', got " + result)
	}
}
func TestHelloWorld(t *testing.T) {
	result := SayHallo("world")
	if result != "Hello world" {
		panic("Expected 'Hello world', got " + result)
	}
}
