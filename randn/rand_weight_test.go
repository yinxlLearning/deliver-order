package randn

import (
	"fmt"
	"math"
	"sdInterview/constant"
	"testing"
)

func TestInitRandWeight(t *testing.T) {
	randWeightFunc := InitRandWeight()

	n := 100000

	cnt1 := 0
	cnt3 := 0
	cnt100 := 0

	for i := 0; i < n; i++ {
		tr := randWeightFunc()
		ti := int(math.Ceil(tr))
		if ti == 1 {
			cnt1++
		}
		if ti == 3 {
			cnt3++
		}
		if ti == 100 {
			cnt100++
		}
	}

	rate1 := float64(cnt1) / float64(n)
	fmt.Println("1出现的概率为：", rate1)
	rate3 := float64(cnt3) / float64(n)
	fmt.Println("3出现的概率为：", rate3)
	rate100 := float64(cnt100) / float64(n)
	fmt.Println("100出现的概率为：", rate100)

	sums := 0.0
	for i := 1; i <= constant.WeightMax; i++ {
		sums += 1 / float64(i)
	}
	fmt.Println("1出现的理论概率为：", 1.0/sums)
	fmt.Println("3出现的理论概率为：", 1/3.0/sums)
	fmt.Println("100出现的理论概率为：", 1/100.0/sums)

}
