package main

func reverseVowels(s string) string {
	var mapping map[byte]interface{}
	mapping = map[byte]interface{}{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}
	length := len(s)
	l := 0
	r := length - 1
	sSlien := []byte(s)
	for l < r {
		for l < r {
			_, ok := mapping[sSlien[l]]
			if !ok {
				l++
			} else {
				break
			}
		}
		for r > l {
			_, ok := mapping[sSlien[r]]
			if !ok {
				r--
			} else {
				break
			}
		}
		if l >= r {
			break
		}
		sSlien[l], sSlien[r] = sSlien[r], sSlien[l]
		l++
		r--
	}
	return string(sSlien)
}
