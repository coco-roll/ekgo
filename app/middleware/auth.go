package middleware

import (
	"ekgo/api/app/model"
	"ekgo/api/app/service"
	"ekgo/api/boot/db"
	"ekgo/api/lib/conv"
	"ekgo/api/lib/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

var Interface service.AdminInterface

//后台管理中间件授权认证登录
func AdminAuth() gin.HandlerFunc {
	return func(this *gin.Context) {
		token := this.Request.Header.Get("Authorization")
		if token == "" || token == "null" || token == "undefined" {
			fmt.Println(token)
			this.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			this.Abort()
			return
		}
		//解析JWT
		result, err := jwt.Decode(token)
		if err == nil {
			model := model.Admin{Id: conv.Int(result["id"])}
			Interface = &service.Admin{Model: model, Db: db.Master}
			var row = Interface.GetCache()

			if row.Id == 0 || model.Status == "false" {
				this.JSON(200, gin.H{
					"code": 402,
					"msg":  "权限不足,请求失败",
				})
				this.Abort()
				return
			}

			if result["exp"].(float64) < conv.GetTimestamp() {
				this.JSON(200, gin.H{
					"code": 402,
					"msg":  "登陆已超时,重新登陆",
				})
				this.Abort()
				return
			}

			this.Set("user", row)
			this.Next()
		} else {
			this.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			this.Abort()
			return
		}
	}
}

//获取商家用户信息
func GetAdmin(this *gin.Context) model.Admin {
	user := this.MustGet("user").(model.Admin)
	return user
}
