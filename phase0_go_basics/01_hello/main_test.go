package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	got := greet("сеть")
	want := "Привет, сеть!"
	if got != want {
		t.Fatalf("greet() = %q, want %q", got, want)
	}
}

func TestGreetInternational(t *testing.T) {
	name := "мир"
	got := greetInternational(name)
	if got == "" {
		t.Fatal("greetInternational() вернула пустую строку")
	}
	if !strings.Contains(got, name) {
		t.Fatalf("greetInternational() = %q не содержит %q", got, name)
	}
}