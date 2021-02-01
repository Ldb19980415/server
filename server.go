package main

import (
	"fmt"
	// "gorm.io/gorm"
	// "goserver/utils"
	// "goserver/dao"
	// "github.com/gin-gonic/gin"
	"goserver/router"


	
)
type jhdkj  interface{

}
var port string = "0.0.0.0:3005"
// var db *gorm.DB
// func init()  {
// 	db = dao.DBpool()
// }

func main()  {
	//生成密钥对，保存到文件
	// utils.GenerateRSAKey(248)
	//加密
	// message:=[]byte("123456")
	// cipherText := utils.RSA_Encrypt(message,"public.pem")
	// fmt.Println("加密后为：",string(cipherText))
	//解密
	// plainText := utils.RSA_Decrypt(cipherText, "private.pem")
	// fmt.Println("解密后为：",string(plainText))
	fmt.Println("====================")
	// 操作数据库
	// dao.CreatRow1(db)
	// dao.CreatRow2(db)
	router1 := router.SetupRouter()

	router1.Run(port)
}


















