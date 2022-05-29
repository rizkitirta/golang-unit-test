package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloTirta(t *testing.T) {
	result := SayHallo("Tirta")
	if result != "Hello Tirta" {
		// panic("Expected 'Hello Tirta', got " + result)
		// t.Fail()
		t.Error("Expected 'Hello Tirta', got " + result)
	}

	fmt.Println("Test Hello Tirta Done")
}
func TestHelloWorld(t *testing.T) {
	result := SayHallo("world")
	if result != "Hello world" {
		// panic("Expected 'Hello world', got " + result)
		// t.Fail()
		t.Error("Expected 'Hello world', got " + result)
	}

	fmt.Println("Test Hello world Done")
}

func TestSayHelloWithAssert(t *testing.T) {
	result := SayHallo("assert")
	assert.Equal(t, "Hello assert", result)

	fmt.Println("Test Hello assert Done")
}

func TestSayHelloWithRequire(t *testing.T) {
	result := SayHallo("require")
	require.Equal(t, "Hello require", result)

	fmt.Println("Test Hello require Done")
}
