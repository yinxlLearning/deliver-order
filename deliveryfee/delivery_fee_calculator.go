package deliveryfee

import (
	"log"
	"math"
	"sdInterview/constant"
	sqlinit "sdInterview/db"
)

// Calculator 快递费用计算接口
type Calculator interface {
	Calculate(order sqlinit.Order) int
}

// DefaultCalculator 普通的迭代实现
type DefaultCalculator struct{}

// NewDefaultCalculator 构造器
func NewDefaultCalculator() *DefaultCalculator {
	return &DefaultCalculator{}
}

// Calculate 计算订单的费用
func (dc DefaultCalculator) Calculate(order sqlinit.Order) int {
	if order.Weight < 0 {
		return 0
	}
	weightCharged := getWeightCharged(order.Weight)
	fee := calculateFee(weightCharged)
	return fee
}

// 根据重量计算计费时的重量
func getWeightCharged(weight float64) int {
	if weight <= 0 {
		return 0
	}
	return int(math.Ceil(weight))
}

// 根据计费重量，计算费用
func calculateFee(weightCharged int) int {
	if weightCharged <= 0 {
		return 0
	}
	if weightCharged > constant.WeightMax {
		log.Fatalln("不支持的重量", weightCharged)
	}

	var res int
	for i := 1; i <= weightCharged; i++ {
		if i == 1 {
			res += 18
			continue
		}
		// 四舍五入取整数值
		cur := int(math.Round(5.0 + 0.01*float64(res)))
		res += cur
	}
	return res
}

// CacheCalculator
// 当频繁计算的时候，可以通过缓存各个重量，避免频繁计算
type CacheCalculator struct {
	// 费用缓存
	cache map[int]int
}

func NewCacheCalculator() *CacheCalculator {
	cache := make(map[int]int, constant.WeightMax)
	for i := 1; i <= constant.WeightMax; i++ {
		cache[i] = calculateFee(i)
	}
	return &CacheCalculator{cache: cache}
}

// Calculate 计算费用
func (calculator CacheCalculator) Calculate(order sqlinit.Order) int {
	if order.Weight > constant.WeightMax || order.Weight < 0 {
		log.Fatalln("不支持的重量", order.Weight)
	}
	weightCharged := getWeightCharged(order.Weight)
	if weightCharged == 0 {
		return 0
	}
	return calculator.cache[weightCharged]
}
