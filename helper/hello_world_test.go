package helper

import (
	"fmt"
	"testing"
)

func TestHelloTirta(t *testing.T) {
	result := SayHallo("Tirta")
	if result != "Hello Tirta" {
		// panic("Expected 'Hello Tirta', got " + result)
		// t.Fail()
		t.Error("Expected 'Hello Tirta', got " + result)
	}

	fmt.Println("Test Hello Tirta Success")
}
func TestHelloWorld(t *testing.T) {
	result := SayHallo("world")
	if result != "Hello world" {
		// panic("Expected 'Hello world', got " + result)
		// t.Fail()
		t.Error("Expected 'Hello world', got " + result)
	}

	fmt.Println("Test Hello world Success")
}
