package piscine

func Pow(baselen, nbrPosition int) int {
	if nbrPosition == 0 {
		return 1
	}
	if nbrPosition == 1 {
		return baselen
	}
	return baselen * Pow(baselen, nbrPosition-1)
}

func index(baseFrom string, char byte) int {
	for i := 0; i < len(baseFrom); i++ {
		if baseFrom[i] == char {
			return i
		}
	}
	return -1
}

func Atoibase(nbr, baseFrom string) int {
	var res int
	baseLen := len(baseFrom)

	for i := 0; i < len(nbr); i++ {
		n := index(baseFrom, nbr[i])
		res += n * Pow(baseLen, len(nbr)-1-i)
	}
	return res
}

func ReturnNbrBase(nbr int, base string) string {
	baselen := len(base)
	return ReturnRecursively(nbr, baselen, base)
}

func ReturnRecursively(nbr, baselen int, base string) string {
	var s []rune
	if nbr >= baselen {
		s = append(s, []rune(ReturnRecursively(nbr/baselen, baselen, base))...)
	}
	i := nbr % baselen
	if i < 0 {
		i += baselen
	}
	s = append(s, rune(base[i]))
	return string(s)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := Atoibase(nbr, baseFrom)
	return ReturnNbrBase(res, baseTo)
}
