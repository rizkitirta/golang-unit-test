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

func TestSkip(t *testing.T) {
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

func TestSubTest(t *testing.T) {

	result := SayHallo("subtest")
	assert.Equal(t, "Hello subtest", result)

	t.Run("Hello", func(t *testing.T) {
		result := SayHallo("hello")
		assert.Equal(t, "Hello hello", result)
	})

	t.Run("World", func(t *testing.T) {
		result := SayHallo("World")
		assert.Equal(t, "Hello World", result)
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Tirta",
			request:  "Tirta",
			expected: "Hello Tirta",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
		{
			name:     "Arif",
			request:  "Arif",
			expected: "Hello Arif",
		},
		{
			name:     "Hanin",
			request:  "Hanin",
			expected: "Hello Hanin",
		},
	}

	for _, val := range tests {
		t.Run(val.name, func(t *testing.T) {
			result := SayHallo(val.request)
			require.Equal(t, val.expected, result)
		})
	}
}

// benchmark
func BenchmarkTirta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHallo("Tirta")
	}
}
func BenchmarkWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHallo("Joko Purnomo")
	}
}