package helloworld

import "strings"

func Greeting(name string, language string) string {
	greeting := ""
	switch strings.ToLower(strings.TrimSpace(language)) {
	case "spanish":
		greeting = "Hola Amigo, "
	case "japanese":
		greeting = "こんにちは、"
	default:
		greeting = "Hello World, "
	}
	if name == "" {
		name = "Kind Stranger"
	}

	return greeting + name

}
