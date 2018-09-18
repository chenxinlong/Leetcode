package algs

func FindTheDifference(s string, t string) byte {
	var xorS, xorT byte
	for i:=0; i < len(s); i++ {
		xorS ^= s[i]
	}

	for j:=0; j < len(t); j++ {
		xorT ^= t[j]
	}

	return xorS ^ xorT
}

