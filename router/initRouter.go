package router



import (
    "github.com/gin-gonic/gin"
	"goserver/handler"
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

	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.LoginHandler)
		auth.POST("/logout", handler.LogoutHandler)
	}



	// 注册自定义验证器
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("bookabledate", bookableDate)
	// }
	// router.GET("/bookable", getBookable)





    // 添加 Get 请求路由
    // router.GET("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, handler.UserSave)
	// })
	// router.GET("/v0/", handler.UserSave)
	// router.POST("/v0/",handler.RetHelloGinAndMethod)
	// // router.GET("/user/:name",handler.Save)			//支持在路径中传参，通过context.Param来获取对应参数，如果当前的context.Param("name")
	// //一般用不到，我也灭整明白怎么用。
	// router.GET("/v0/save",handler.UserSaveByQuery)
	// router.GET("/v0/ldbtest",handler.UserSaveByQuery)


	// // 路由分组
	// v1 := router.Group("/v1")
	// {
	// 	v1.POST("/login", handler.UserSave)
	// 	// v1.POST("/submit", submitEndpoint)
	// 	// v1.POST("/read", readEndpoint)
	// }

	// v2 := router.Group("/v2")
	// {
	// 	v2.POST("/login", handler.UserSave)
	// 	// v2.POST("/submit", submitEndpoint)
	// 	// v2.POST("/read", readEndpoint)
	// }


	// router.POST("/form_post", func(c *gin.Context) {
	// 	message := c.PostForm("age")
	// 	nick := c.DefaultPostForm("name", "kim")

	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"nick":    nick,
	// 	})
	// })



	// router.POST("/test", func(context *gin.Context) {
	// 	var person Person
    //     // 这里我确定传过来的一定是JSON所以用ShouldBindJSON，否则可以用ShouldBind
	// 	if err := context.ShouldBindJSON(&person); err != nil { 
	// 		context.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"success": true,
	// 	})
	// })


	
    // // 添加 Post 请求路由
    // router.POST("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
    // })
    // // 添加 Put 请求路由 
    // router.PUT("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
    // })
    // // 添加 Delete 请求路由
    // router.DELETE("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
    // })
    // // 添加 Patch 请求路由
    // router.PATCH("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
    // })
    // // 添加 Head 请求路由
    // router.HEAD("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
    // })
    // // 添加 Options 请求路由
    // router.OPTIONS("/", func(context *gin.Context) {
    //     context.String(http.StatusOK, RetHelloGinAndMethod)
	// })
	


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
