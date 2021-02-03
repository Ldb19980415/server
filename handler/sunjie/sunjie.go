package sunjie

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goserver/dao"
	"fmt"
	"time"
)

func CreateWeightHandler (ctx *gin.Context) {
	var params CreateWeight
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"message":err.Error,
		})
	}
	fmt.Println(params.CurrentWeight)
	var weightInfo dao.WeightInfo
	weightInfo.CurrentWeight = params.CurrentWeight
	weightInfo.RecordTime = time.Now()
	result := dao.LikeDB.Create(&weightInfo)
	if result.Error != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"success":false,
			"message":result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"success":true,
		"record":weightInfo,
	})

}

func SearchWeightHandler( ctx *gin.Context)  {
	var weightInfos  []dao.WeightInfo
	result := dao.LikeDB.Limit(10).Order("createtime desc").Find(&weightInfos)
	if result.Error != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"success":false,
			"message":result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"success":true,
		"record":weightInfos,
	})

}