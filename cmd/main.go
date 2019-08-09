package main

import (
	"alertmanager_dingtalk_webhook/alert"
	"alertmanager_dingtalk_webhook/types"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	dingtalkRobot string
)

func init() {
	flag.StringVar(&dingtalkRobot, "dingtalkRobot", "", "dingtalk webhook robot url")
}

func main() {
	flag.Parse()
	if dingtalkRobot == "" {
		log.Fatal("dingtalkRobot can not be empty!!!")
	}
	log.Printf("dingtalkRobot----->%s", dingtalkRobot)

	router := gin.Default()
	router.POST("/webhook", func(ctx *gin.Context) {
		var alertMessage types.AlertMessage
		if err1 := ctx.BindJSON(&alertMessage); err1 != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error": err1.Error(),
			})
			return
		}
		if err2 := alert.Send(alertMessage, dingtalkRobot); err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error": err2.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.Run(":8060")
}
