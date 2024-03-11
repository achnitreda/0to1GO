package piscine

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func Power(baselen, nbrPosition int) int {
	if nbrPosition == 0 {
		return 1
	}
	if nbrPosition == 1 {
		return baselen
	}
	return baselen * Power(baselen, nbrPosition-1)
}

func IndexOf(baseFrom string, char byte) int {
	for i := 0; i < len(baseFrom); i++ {
		if baseFrom[i] == char {
			return i
		}
	}
	return -1
}

func AtoiBase(nbr, baseFrom string) int {
	if !isValidBase(baseFrom) {
		return 0
	}
	var res int
	baseLen := len(baseFrom)

	for i := 0; i < len(nbr); i++ {
		n := IndexOf(baseFrom, nbr[i])
		res += n * Power(baseLen, len(nbr)-1-i)
	}
	return res
}
