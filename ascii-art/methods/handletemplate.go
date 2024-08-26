package methods

func Template(str string) string {
	switch str {
	case "shadow", "standard", "thinkertoy":
		return str + ".txt"
	}
	return str
}

func IsValidTemplate(str string) bool {
	switch (str) {
		case "shadow", "standard", "thinkertoy":
			return true
	}
	return false
}

func IsValidType(str string) bool {
	switch str {
	case "right", "left", "center", "justify":
		return true
	}
	return false
}
