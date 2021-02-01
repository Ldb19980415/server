package router



import (
    "github.com/gin-gonic/gin"
	"goserver/handler/auth"
	// "goserver/handler"
	// "goserver/dao"
	// "gomod/middlewares"
	// "net/http"
	// "reflect"
	// "time"

	// "github.com/gin-gonic/gin/binding"
	// "gopkg.in/go-playground/validator.v8"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// router.Use(middlewares.Cors())

	// 路由分组
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", auth.LoginHandler)
		authRouter.POST("/logout", auth.LogoutHandler)
	}

	// 路由不分组
	// router.GET("/books", handler.GetBooks)

	
	


	// 注册自定义验证器
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("bookabledate", bookableDate)
	// }
	// router.GET("/bookable", getBookable)

	// 自定义一个参数校验器
    return router
}
// type Person struct {
// 	Name string `json:"name" binding:"required"` // json格式从name取值，并且该值为必须的
// 	Age  int    `json:"age" binding:"required,gt=20"` // json格式从age取值，并且该值为必须的，且必须大于20
// }
// type Booking struct {
// 	// 这里的验证方法为bookabledate
// 	  CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
// 	  // gtfield=CheckIn表示大于的字段为CheckIn
// 	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
// }
  
// func bookableDate(
// 	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
// 	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
// ) bool {
//   // 这里有两个知识点，映射和断言
//   // 在这里，field是一个reflect.Type的接口类型变量，通过Interface方法获得field接口类型变量的真实类型，可以理解为reflect.Value的逆操作
//   // 在这里，断言就是将一个接口类型的变量转化为time.Time，前提是后者必须实现了前者的接口
//   // 综上，这里就是将field进行了类型转换
// 	if date, ok := field.Interface().(time.Time); ok {
// 		today := time.Now()
// 		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
// 			return false
// 		}
// 	}
// 	return true
// }
// func getBookable(c *gin.Context) {
// 	var b Booking
// 	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
// 		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// }
