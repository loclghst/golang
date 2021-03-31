package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	return "Hello There !!!"
}

// since we are not using sb and eb variables anywhere inside the getGreeting func body
// we can omit it as follows

// func (englishBot) getGreeting() {
// 	return "Hello There !!!"
// }

func (spanishBot) getGreeting() string {
	return "Hola !!!"
}

// both the printGreeting functions have almost the same body
// we can use interface to make code more reusable

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreetingx(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// interface approach
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
