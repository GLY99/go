package main

func finalPrices(prices []int) []int {
	// 维护一个单调递增的栈
	length := len(prices)
	stack := []int{0} // 最后一个商品一定没有折扣，这里提前填个0
	ans := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		p := prices[i]
		for len(stack) > 1 && stack[len(stack)-1] > p {
			stack = stack[:len(stack)-1]
		}
		ans[i] = p - stack[len(stack)-1]
		stack = append(stack, p)
	}
	return ans
}
