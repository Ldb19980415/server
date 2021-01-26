package main

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

func main()  {
	fmt.Println("数据库连接")
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
	creatRow2(db)
}


//	创建一条新数据
func creatRow1(db *gorm.DB)  {
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
func creatRow2(db *gorm.DB)  {
	user := User{
		Name :"李四",
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



// 建立数据库连接
// func setupConnection() *sql.DB {
// 	db,err := sql.Open("mysql","like:xa4mmPmzyfGPRFts@tcp(39.98.146.215:3306)/like?charset=utf8")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return db
// }


var CreatUserTable = "CREATE TABLE user(" +
	"id INT(10) NOT NULL AUTO_INCREMENT," +
	"name VARCHAR(64)  DEFAULT NULL," +
	"phone VARCHAR(64)  DEFAULT NULL," +
	"password VARCHAR(64)  DEFAULT NULL," +
	"active INT(1)  DEFAULT NULL, PRIMARY KEY (id))" +
	"ENGINE=InnoDB DEFAULT CHARSET=utf8;" 

// 创建表
// func creatTable(db *sql.DB,sql string) {
// 	_, err := db.Exec(sql)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

var INSERT_DATA = `INSERT INTO user(name,phone,password,active) VALUES(?,?,?,?);`

// 插入数据
// func Insert(db *sql.DB,name string,phone string, password string) {
//    info,err := db.Exec(INSERT_DATA, name, phone, password , 1)
// 	fmt.Println(err)
// }