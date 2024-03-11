package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for {
		prime := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return nb
		}
		nb++
	}
}
