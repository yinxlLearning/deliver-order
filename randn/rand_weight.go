package randn

import (
	"math/rand"
	"sdInterview/constant"
)

// InitRandWeight 逆变换采样思想实现伪随机
// 详细说明请看根目录下的markdown文件
func InitRandWeight() func() float64 {
	// 权重总和
	var weightSums float64

	// 计算权重的总和
	initWeightParams(&weightSums)

	// weightProbabilityPrefixSum[i]表示随机出现的重量小于等于i时的概率前缀和
	weightProbabilityPrefixSum := make([]int, constant.WeightMax+1)

	// 初始化前缀和
	initWeightProbabilityPrefixSum(weightProbabilityPrefixSum, weightSums)

	// 二分查找，判断mid位置表示的重量的区间是否为随机数所在的概率区间
	isMatch := func(mid int, randNum int) bool {
		return randNum > weightProbabilityPrefixSum[mid-1] && randNum <= weightProbabilityPrefixSum[mid]
	}

	// 返回概率生成函数，每次调用生成一个随机数
	return func() float64 {
		// 随机数生成的上边界
		randUpBound := weightProbabilityPrefixSum[constant.WeightMax]
		// 生成随机数
		randNum := rand.Intn(randUpBound) + 1

		left, right := 1, constant.WeightMax+1

		var res = -1.0

		// 二分查找概率对应的重量
		// 左闭右开，出口为left==right
		for left < right {
			mid := left + (right-left)/2
			if isMatch(mid, randNum) {
				res = float64(mid)
				break
			} else if randNum > weightProbabilityPrefixSum[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 未找到
		if res < 0 {
			return res
		}
		res -= float64(rand.Intn(100)) * 0.01
		return res
	}
}

func initWeightProbabilityPrefixSum(weightProbabilityPrefixSum []int, weightSums float64) {
	for i := 1; i <= constant.WeightMax; i++ {
		weightProbabilityPrefixSum[i] = weightProbabilityPrefixSum[i-1] + getWeightMapping(1.0/float64(i)/weightSums)
	}
}

// 将浮点数概率映射为整数,方便计算
func getWeightMapping(r float64) int {
	return int(r * constant.ScaleFactor)
}

func initWeightParams(weightSums *float64) {
	for i := 1; i <= constant.WeightMax; i++ {
		t := 1 / float64(i)
		*weightSums += t
	}
}
