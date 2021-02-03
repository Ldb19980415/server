package dao
import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"time"
	"reflect"
)

var LikeDB *gorm.DB
func init()  {
	LikeDB = DBpool()
}



type User struct {
	ID           uint  `json:"id"`
	Name         string  `json:"name"`
	Email        *string	`json:"email"`
	Age          uint8
	Active       uint8
	MemberNumber sql.NullString
	ActivedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserInfo struct {
	ID	uint	`json:"id"`
	UserName	string	`json:"username"`
	PassWord	string 	`json:"password"`
	CreatAt		time.Time	`gorm:"column:createtime;default:null",json:"creat_at"`
}

type WeightInfo struct {
	ID uint `json:"id"`
	UserId uint `json:"user_id"`
	CurrentWeight string `json:"current_weight"`
	RecordTime time.Time `gorm:"column:createtime;default:null",json:"Record_time"`
}

func DBpool() *gorm.DB {
	// 使用GORM建立数据库连接
	dsn := "like:xa4mmPmzyfGPRFts@tcp(39.98.146.215:3306)/like?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(reflect.TypeOf(db))
	fmt.Println()
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	fmt.Println("连接数据库成功")
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&WeightInfo{})
	// db.AutoMigrate(&User{})			迁移数据表，如果没有该表会自动的给予创建
	// creatRow1(db)
	return db
}
