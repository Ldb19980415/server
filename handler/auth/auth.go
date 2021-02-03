package auth



import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goserver/dao"
)
// [{url:"jkhs",hander:LoginHandler}]
func LoginHandler (ctx *gin.Context) {
	var userInfo loginSerializer
	// 这里我确定传过来的一定是JSON所以用ShouldBindJSON，否则可以用ShouldBind
	if err := ctx.ShouldBindJSON(&userInfo); err != nil { 
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user dao.UserInfo
	result := dao.LikeDB.Where("user_name = ? AND pass_word = ?", userInfo.UserName, userInfo.PassWord).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"success": true,
		"data":user,
	})
}
func LogoutHandler(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}




/*
get参数解析：
	name := ctx.Query("name")
	name := ctx.DefaultQuery("name","defaultvalue")
	names := ctx.QueryArray("names")
	nameobj := ctx.QueryMap("nameMap")


post参数解析：
	name := ctx.DefaultPostForm("username","defaultvalue")
	age := ctx.PostForm("age")
	add := ctx.PostForm("add")	//不存在的参数
	arr_ck := ctx.PostFormArray("ctx")
	map_other := ctx.PostFormMap("other")



但是绑定了校验之后，直接取用于绑定校验的变量即可
*/


// func LoginHandler (ctx *gin.Context) {
// 	var userInfo loginSerializer
// 	// 这里我确定传过来的一定是JSON所以用ShouldBindJSON，否则可以用ShouldBind
// 	if err := ctx.ShouldBindJSON(&userInfo); err != nil { 
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	var user dao.UserInfo
// 	result := db.Where("user_name = ? AND pass_word >= ?", userInfo.UserName, userInfo.PassWord).First(&user)
// 	if result.Error != nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": result.Error,
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status":http.StatusOK,
// 		"success": true,
// 		"data":user,
// 	})
// }
// func LogoutHandler(context *gin.Context)  {
// 	context.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 	})
// }



/*

gorm使用笔记：

1.定义模型：
type User struct{ 对应到数据库的表为user
	ID	uint	对应到数据库的id，约定俗成自增
	Name	string	
	Email	*string
	Age		uint8
	Birthday	*time.Time
	MemberNumber	sql.NullString
	ActivedAt		sql.NullTime
	CreateAt		time.Time
	UpdateAt		time.Time
}
GORMq倾向于约定，而不是配置。默认情况下，GORM使用ID作为主键，使用结构体名的蛇形复数作为表名，字段名的蛇形作为列名。
并使用CreatedAt、UpdateAt字段追踪创建、更新时间。

---->约定
	1)使用ID作为主键
	默认情况下，GORM会使用ID作为表的主键。
	type User struct{
		ID	string		//默认情况下，名为`ID`的字段会作为表的主键
		Name	string	
	}
	你可以通过标签primaryKey将其他字段设为主键
	`gorm:"primaryKey"`	
	设置复合主键：同时给多列设置primarykey。默认情况下，主键是启动了自增，不需要的话可以通过aotoIncrement:false来关闭
	type Product struct {
		ID		string 	`gorm:"primaryKey"`
		LanguageCode	string	`gorm:"primaryKey"`
		Name	string
		Code	string
	}
	type Product struct {
		CategoryID	uint64	`gorm:"primaryKey";autoIncrement:false`
		TypeID		uint64	`gorm:"primaryKey";autoIncrement:false`
	}
	2)临时指定表名
	您可以使用Table方法临时指定表名，例如：
		db.Table("delete_users").AutoMigrate(&User{})
		//根据User的字段创建`deleted_users`表

		var deletedUsers []User
		db.Table("deleted_users").Find(&deleteUsers)
		select * from delete_users

		db.Table("deleted_users").Where("name"=?,"jinzhu").Delete(&User{})
		// delete from delete_users where name = "jinzhu"

2.连接到数据库		

	MYSQL：
	import (
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
	)
	func main(){
		dsn := "user:pass@tcp(host:port)/dbname?charset=utf8m64&parseTime=True&loc=Local"
		db ,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

		//设置连接池
		sqlDB ,err := db.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns	设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime	设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

3.使用现有结构体创建表(若表存在则不创建)
	// db.AutoMigrate(&User{})			迁移数据表，如果没有该表会自动的给予创建


CRUD使用：

Create：-------------->
	创建记录：
	user := User{Name : "jinzhu", Age:18, Birthday : time.Now()}
	result := db.Create(&user)		// 通过数据的指针来创建。user是属于User结构体的，创建的时候会丢到以User对应的表里面去。

	user.ID			//此时user就是插入到数据库中的刚才插入的这条数据的信息。通过user.ID能看到该条数据在表中的id
	result.Error	//如果有错误的话，result.Error为错误信息，否则为nil
	result.RowsAffected		//影响的行数，在这里指插入的记录数量


	//	使用指定的字段创建纪录
	创建记录并更新给出的字段。
	db.Select("Name","Age","CreatAt").Create(&user)
	Insert into users ("name","age","create_at") values ("zhangjs", 18, "2020-01-22")
		
	创建记录并更新未给出的字段。
	db.Omit("Name","Age","CreateAt").Create(&user)
	insert into users ("birthday","update_at") values ("2021-01-22 ..." , "......")

	// 批量插入
	通过一个切片slice传递给Create方法。
	var users = []User{{"Name":"zhangsan"},{"Name":"lisi"},{"Name":"wangba"}}
	db.Create(&users)

	//此时的users为一个插入到表并返回的插入数据信息。通常我们还是希望看到插入数据的ID
	for _, user := range users {
		user.ID		// 1,2,3    被放弃的字段为数组的索引，拿着没用，弃掉
	}
		//这里我们可以加入一个result，通过result.Error和result.RowsAffected查看插入数据的概要信息
	使用CreateBatches创建时，可以指定创建的数量，
	var users := []User{{Name:"khdka"},.....{"name":"hkjahdkja"}}
	db.CreateBatches(users,100)	//插入100条

	**创建钩子
	钩子方法有：BeforeCreate,AfterSave,AfterCreate。。。。。。。
	例如给User结构体绑定一个钩子方法：
	func (u *User) BeforeCreate(tx *gorm.DB) (err error){
		u.UUID = uuid.New()
		if u.Role == "admin" {
			return errors.New("invalid role")
		}
		return
	}

	默认值：在定义结构体字段的是时候加入default，那么当插入数据的时候没有该字段，则用默认值替代
	type User struct {
		ID int64
		Name string `gorm:"default:jslkajslka"`
		Age int64 	`gorm:"default:18"`
	}


查询：
	GORM提供了First、Take、Last方法，以便从数据库中检索单个对象。当查询数据库时它添加了LIMIT 1条件，且没有找到记录时，会返回ErrRecordNotFound错误。
	var user User

	db.First(&user)
	select * from users order by id limit 1;
	
	db.Take(&user)
	select * from users limit		随便乱来一条

	db.Last(&user)
	select * from users order by id Desc limit 1;

	result := db.First(&user)
	result.RowsAffected
	result.Error

	// 避开报ErrRecordNotFound：使用Find进行查询
	例如：db.Limit(1).Find(&user)

	根据主键检索
	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&users, []int{1,2,3})
	// SELECT * FROM users WHERE id IN (1,2,3);


	检索全部对象
	//获取全部记录
	var users []User
	result := db.Find(&users)
		----> select * from users
	result.RowsAffected			相当于len(users)
	result.Error		错误信息。
		//此时users是一个数组，通常会使用for range遍历或者是直接返回给前端(此时字段名任然是定义结构体时候的样子，通常加个json的tag来时返回的前端的字段名比较正常)
	
	按照条件查找：
		String条件
		







*/