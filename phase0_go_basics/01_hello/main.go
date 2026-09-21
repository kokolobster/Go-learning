package main

import "fmt"

func greet(name string) string {
	return "Привет, " + name + "!"
}

func greetInternational(name string) string {
	// ЗАДАНИЕ (день 1): дополните приветствие, чтобы в терминале появилась
	// и "Hello", и "Привет". Подсказка: верните "Hello, " + "Привет, " + name + "!"
	return "Hello, " + name + "!"
}

func main() {
	dear := "сеть"
	fmt.Println(greet(dear))
	fmt.Println(greetInternational(dear))
	fmt.Printf("байтов в слове %q: %d\n", dear, len(dear))
}