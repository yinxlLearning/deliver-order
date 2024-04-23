package sqlinit

import "gorm.io/gorm"

// QueryOrdersByUserId 根据用户id查询该用户的订单详情
func QueryOrdersByUserId(db *gorm.DB, userId uint) (orders []Order) {
	db.Where("uid = ?", userId).Order("updated_at desc").Find(&orders)
	return
}
