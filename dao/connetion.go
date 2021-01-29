package dao
import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"time"
	"reflect"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Active       uint8
	MemberNumber sql.NullString
	ActivedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

	// db.AutoMigrate(&User{})			迁移数据表，如果没有该表会自动的给予创建
	// creatRow1(db)
	return db
}

//	创建一条新数据
func CreatRow1(db *gorm.DB)  {
	user := User{
		Name :"张三",
		Age:19,
		Active:1,
		ActivedAt:time.Now(),
	}
	// INSERT INTO `users` (`name`,`email`,`age`,`active`,`member_number`,`actived_at`,`created_at`,`updated_at`) VALUES ('张三',NULL,19,1,NULL,'2021-01-22 11:41:39.68','2021-01-22 11:41:39.781','2021-01-22 11:41:39.781')
	result := db.Create(&user)    //通过数据的指针来创建
	fmt.Println(user.ID)		//返回插入数据的主键
	fmt.Println(result.Error)		//返回error
	fmt.Println(result.RowsAffected)		//返回插入记录的条数
	//这个的插入数据会把没有的字段自动按照要求填充
}
//	创建一条新数据
func CreatRow2(db *gorm.DB)  {
	user := User{
		Name :"李四sss",
		Age:19,
		Active:1,
		ActivedAt:time.Now(),
	}
	result := db.Select("Name", "Age", "Active","ActivedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`active`,`actived_at`,`created_at`,`updated_at`) VALUES ('李四',19,1,'2021-01-22 11:47:59.258','2021-01-22 11:47:59.311','2021-01-22 11:47:59.311')
	fmt.Println(user.ID)		//返回插入数据的主键
	fmt.Println(result.Error)		//返回error
	fmt.Println(result.RowsAffected)		//返回插入记录的条数
	// 虽然做了选择，但是没有的字段还是会依照设计库的时候的要求填上默认值
}