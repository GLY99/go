package main

func kthCharacter(k int) byte {
	s := []byte{'a'}
	for len(s) < k {
		tmp_s := []byte{}
		for i := 0; i < len(s); i++ {
			tmp_s = append(tmp_s, 'a'+((s[i]-'a'+1)%26))
		}
		s = append(s, tmp_s...)
	}
	return s[k-1]
}