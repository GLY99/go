package main

func numUniqueEmails(emails []string) int {
	mySet := make(map[string]struct{})
	for _, email := range emails {
		newEmail := ""
		ignore := false
		atIdx := len(email)
		for i := 0; i < len(email); i++ {
			if email[i] == '@' {
				atIdx = i
			} else if email[i] == '+' && !ignore {
				ignore = true
				continue
			} else if email[i] == '.' && i < atIdx {
				continue
			}
			if ignore {
				if i >= atIdx {
					newEmail += string(email[i])
				}
			} else {
				newEmail += string(email[i])
			}
		}
		mySet[newEmail] = struct{}{}
	}
	return len(mySet)
}
