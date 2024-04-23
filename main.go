package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"sdInterview/constant"
	"sdInterview/db"
	"sdInterview/deliveryfee"
	"sdInterview/randn"
)

// 主流程
func main() {
	log.Println("#############开始执行程序#############")

	// 初始化sql连接
	db := sqlinit.InitDBConnection()

	// 初始化用户id随机生成函数
	// 用1-1000表示用户id
	randIdFunc := func() uint {
		return uint(rand.Intn(constant.IdNum) + 1)
	}

	// 初始化权重随机生成函数
	randWeightFunc := randn.InitRandWeight()

	// 生成测试数据并插入数据库中
	generateTestDataInDB(db, randIdFunc, randWeightFunc)

	// 根据用户id，查询并打印用户订单
	queryAndPrintUserOrder(db)

	log.Println("#############结束执行#############")
}

func queryAndPrintUserOrder(db *gorm.DB) {
	log.Println("本程序提供查询功能，为方便使用1-1000表示1000个用户的id")
	log.Println("请输入要查询的用户id: (输入 0 退出程序)")

	var id uint
	for {
		_, err := fmt.Scan(&id)
		if err != nil {
			log.Fatalln(err)
		}
		if id == 0 {
			break
		}
		if id < 0 || id > 1000 {
			log.Println("不支持的用户id")
			continue
		}

		log.Println("------------start--------------")

		// 根据用户id查询订单列表
		orders := sqlinit.QueryOrdersByUserId(db, id)

		// 计算订单总费用
		totalCost := calculateTotalCost(orders)

		// 打印查询的结果
		formatOutput(orders, totalCost)

		log.Println("------------end--------------")
	}
}

func formatOutput(orders []sqlinit.Order, totalCost int) {
	if len(orders) == 0 {
		log.Println("数据库中不存在该用户的相关数据。")
		return
	}
	fmt.Printf("[ 用户ID: %d\t订单总数: %d\t总消费: %8d 元]", orders[0].Uid, len(orders), totalCost)
	fmt.Printf("\n")
	for _, order := range orders {
		fmt.Printf("订单ID: %6d, \t重量: %10.2f KG, \t创建时间: %s, \t更新时间: %s\n",
			order.ID, order.Weight, order.CreatedAt.Format("2006-01-02 15:04:05"), order.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

// 计算订单总费用
func calculateTotalCost(orders []sqlinit.Order) int {
	if len(orders) == 0 {
		return 0
	}
	// 快递费用计算器
	calculator := deliveryfee.NewCacheCalculator()

	var sums int
	for _, order := range orders {
		sums += calculator.Calculate(order)
	}

	return sums
}

func generateTestDataInDB(db *gorm.DB, randIdFunc func() uint, randWeightFunc func() float64) {
	log.Println("开始生成测试数据")

	data := make([]sqlinit.Order, constant.OrderNum)
	for i := 0; i < constant.OrderNum; i++ {
		data[i].Uid = randIdFunc()
		data[i].Weight = randWeightFunc()
	}

	log.Println("测试数据生成成功")

	log.Println("开启数据库事务，开始执行插入操作")

	// 事务插入数据
	err := db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < constant.OrderNum; i++ {
			if err := tx.Create(&data[i]).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}
		// 提交事务
		return nil
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("数据库插入完成")
}
