package isonebitcharacter

func IsOneBitCharacter(bits []int) bool {
	length := len(bits)
	count := 0
	for i := length - 2; i >= 0; i-- {
		if bits[i] != 1 {
			break
		} else {
			count++
		}
	}
	return count%2 == 0
}
