package iteration

// Default repeat times
const default_repeat_times = 5

func Repeat(c string, t int) string {
	var result string

	if t < 0 {
		t = default_repeat_times
	}

	for i := 1; i <= t; i++ {
		result += c
	}

	return result
}
