package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Unit test with go test
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

// Unit test with testify
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

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "windows" {
		t.Skip("Skip on windows")
	}

	result := SayHallo("skip")
	require.Equal(t, "Hello skip", result)

	fmt.Println("Test Hello skip Done")
}

func TestMain(m *testing.M) {
	fmt.Println("SEBELUM UNIT TEST")
	m.Run()
	fmt.Println("SESUDAH UNIT TEST")
}