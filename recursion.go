package algo

func reverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	return reverse(s[1:]) + string(s[0])
}

func reverse2(s string) string { // do some test with runes instead

	end := len(s) - 1

	var recurse func(s string, idx int) string
	recurse = func(s string, idx int) string {
		if idx < 0 {
			return ""
		}
		return string(s[idx]) + recurse(s, idx-1)
	}

	return recurse(s, end)
}

func countX(str string) int {
	if len(str) == 0 {
		return 0
	}

	if str[0] == 'x' {
		return 1 + countX(str[1:])
	}

	return countX(str[1:])
}
