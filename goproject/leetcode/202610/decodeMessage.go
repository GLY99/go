package main

// https://leetcode.cn/problems/decode-the-message/?envType=problem-list-v2&envId=hash-table

func decodeMessage(key string, message string) string {
	mapping := [26]byte{}
	relpaceIdx := 0
	for i := 0; i < len(key); i++ {
		if key[i] == ' ' {
			continue
		}
		idx := int(key[i] - 'a')
		if mapping[idx] == 0 {
			mapping[idx] = byte(relpaceIdx) + 'a'
			relpaceIdx++
		}
	}
	m := []byte(message)
	for i := 0; i < len(m); i++ {
		if m[i] == ' ' {
			continue
		}
		m[i] = mapping[int(m[i]-'a')]
	}
	return string(m)
}
