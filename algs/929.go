package algs

import "strings"

func numUniqueEmails(emails []string) int {
	uniqueEmails := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		rawEmail := strings.Replace(strings.Split(parts[0], "+")[0], ".", "", -1)+"@"+parts[1]
		if _, ok := uniqueEmails[rawEmail]; !ok {
			uniqueEmails[rawEmail] = true
		}
	}

	return len(uniqueEmails)
}


/**
 * [analysis]
 * string
 */