package main

func calculate(s string) int {
	// "A" 运算：使 x = 2 * x + y；
	// "B" 运算：使 y = 2 * y + x。
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}

func calculate1(s string) int {
	// "A" 运算：使 x = 2 * x + y；
	// "B" 运算：使 y = 2 * y + x。
	// 出现一个"A"，有x+y=(2x+y)+y=2x+2y
	// 出现一个"B"，有x+y=x+(2y+x)=2x+2y
	// 所以每出现一个A/B，都使x+y的值翻倍
	return 1 << len(s)
}
