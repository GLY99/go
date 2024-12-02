package main

// https://leetcode.cn/problems/check-if-two-chessboard-squares-have-the-same-color/?envType=daily-question&envId=2024-12-03
func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	// 奇数行的奇数列是黑色，奇数行的偶数列是白色
	// 偶数行的偶数列是黑色，偶数行的奇数列是白色
	ans1 := (coordinate1[0]-'a'+1)%2 == (coordinate1[1]-'0')%2
	ans2 := (coordinate2[0]-'a'+1)%2 == (coordinate2[1]-'0')%2
	return ans1 == ans2
}
