package main

import "testing"

func assert(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("Got '%s', want '%s'", got, want)
    }
}

func TestHello(t *testing.T) {
    t.Run("Saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        assert(t, got, want)
    })

    t.Run("When empty string is passed, output world", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        assert(t, got, want)
    })
}

// Table driven tests
var helloTests = []struct {
    input string
    want  string
}{
    {"Hello", "Hello, Hello"},
    {"Adrian", "Hello, Adrian"},
    {"", "Hello, World"},
    {"四月份平民", "Hello, 四月份平民"},
}

func TestHelloTables(t *testing.T) {
    for _, tt := range helloTests {
        got := Hello(tt.input)
        if got != tt.want {
            t.Errorf("Hello(%s): got '%s', want: '%s'", tt.input, got, tt.want)
        }
    }
}
