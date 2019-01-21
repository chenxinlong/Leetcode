package algs

func ToLowerCase(str string) string {
	res := []byte{}
	for _, v := range str {
		if v >= 65 && v <= 90 {
			res = append(res, uint8(v+32))
		} else {
			res = append(res, uint8(v))
		}
	}

	return string(res)
}


/**
 * [analysis]
 * easy
 */