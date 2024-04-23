package deliveryfee

import (
	"log"
	sqlinit "sdInterview/db"
	"testing"
)

func TestDefaultCalculator_Calculate(t *testing.T) {
	type args struct {
		order sqlinit.Order
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test0", args: args{order: sqlinit.Order{Weight: 0.5}}, want: 18},
		{"test1", args{order: sqlinit.Order{Weight: 1.8}}, 23},
		{"test2", args{order: sqlinit.Order{Weight: 5.9}}, 43},
		{"test3", args{order: sqlinit.Order{Weight: 10.1}}, 71},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DefaultCalculator{}
			fee := dc.Calculate(tt.args.order)
			if int(fee) != tt.want {
				t.Errorf("Calculate() = %v, want %v", fee, tt.want)
			}
			log.Println("重量：", tt.args.order.Weight, "的费用是：", fee)
		})
	}

	dc := DefaultCalculator{}
	cc := NewCacheCalculator()

	for i := 1; i <= 100; i++ {
		order := sqlinit.Order{Weight: float64(i)}
		cost1 := dc.Calculate(order)
		cost2 := cc.Calculate(order)
		if cost2 != cost1 {
			log.Println(cost1, "--", cost2)
		}
		//log.Println("重量：", order.Weight, "的费用是：", cost1)
	}
}
