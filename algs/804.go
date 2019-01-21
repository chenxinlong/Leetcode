package algs

func UniqueMorseRepresentations(words []string) int {
	arr := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	resMaps := make(map[string]bool)
	for _,  word := range words {
		wordMorse := ""
		for _, v := range word {
			wordMorse += arr[v-'a']
		}
		resMaps[wordMorse] = true
	}


	return len(resMaps)
}

/**
 * [analysis]
 * easy
 */