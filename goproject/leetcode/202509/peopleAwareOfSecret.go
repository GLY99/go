package main

// https://leetcode.cn/problems/number-of-people-aware-of-a-secret/?envType=daily-question&envId=2025-09-09

const MOD = 1000000007

type pair struct {
	day   int // 知道秘密的天数
	total int // 这天中知道秘密的人数
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	know := make([]pair, 0)
	share := make([]pair, 0)
	know = append(know, pair{1, 1})
	knowTotal, shareTotal := 1, 0 // 知道不分享和知道会分享
	for i := 2; i <= n; i++ {
		// 第i - delay天知道秘密的人会在第i天变成知道且分享
		if len(know) > 0 && know[0].day == i-delay {
			first := know[0]
			know = know[1:]
			knowTotal = (knowTotal - first.total + MOD) % MOD // 减去状态发生变化的一批人
			shareTotal = (shareTotal + first.total) % MOD     // 加上这批人
			share = append(share, first)
		}
		// 如果有人是在第i - forget天知道的秘密，他会在第i天忘记
		if len(share) > 0 && share[0].day == i-forget {
			first := share[0]
			share = share[1:]
			// 减去忘记的这批人，因为这批人已经不在knowTotal中了所以不存在重复统计的问题
			shareTotal = (shareTotal - first.total + MOD) % MOD
		}
		// 排除会在第i天忘记秘密的人后，share中的人会在这天分享秘密
		if len(share) > 0 {
			knowTotal = (knowTotal + shareTotal) % MOD
			know = append(know, pair{i, shareTotal})
		}
	}
	return (knowTotal + shareTotal) % MOD
}
