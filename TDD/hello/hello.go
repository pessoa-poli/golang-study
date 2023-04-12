package hello

const helloPrefix = "Hello, "

func choosePrefix(ln string) string {
	switch ln {
	case "Spanish":
		return "Hola, "
	case "French":
		return "Bonjour, "
	default:
		return "Hello, "
	}
}

func Hello(in, language string) string {
	if in == "" {
		in = "World"
	}
	prefix := choosePrefix(language)
	return prefix + in
}
