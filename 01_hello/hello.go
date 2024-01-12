package main

const (
	helloPrefix        = "Hello"
	spanishHelloPrefix = "Hola"
	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + ", " + name
}

func greetingsPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
	println(Hello("Luciano", "Spanish"))
}
