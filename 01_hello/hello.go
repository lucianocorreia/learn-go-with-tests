package main

const helloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix + ", " + name
}

func main() {
	println(Hello("Luciano"))
}
