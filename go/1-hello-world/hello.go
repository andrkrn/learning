package helloworld

const (
	// Hello in english
	helloEn = "Hello, "

	// Hello in indonesian
	helloId = "Hai, "
)

func Hello(n string, lang string) string {
	if n == "" {
		n = "World"
	}

	return Greeting(lang) + n
}

// Get greeting for selected languange
func Greeting(lang string) string {
	switch lang {
	case "id":
		return helloId
	default:
		return helloEn
	}
}
