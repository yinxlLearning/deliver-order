package sqlinit

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sdInterview/constant"
	"time"
)

// Order 订单数据 对应数据库订单表
type Order struct {
	gorm.Model
	Uid       uint      `gorm:"not null;index:idx_uid"`
	Weight    float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}

// InitDBConnection 初始化数据库连接
func InitDBConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(constant.DBPath), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().UTC() // 指定时间格式为 UTC 时间
	}})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	// 清除掉旧数据
	clearOrderData(db)

	err = db.AutoMigrate(&Order{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func clearOrderData(db *gorm.DB) {
	if hasTable := db.Migrator().HasTable("orders"); hasTable {
		log.Println("已经存在orders，删除旧表")
		err := db.Migrator().DropTable(&Order{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
