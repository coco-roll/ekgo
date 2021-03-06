package middleware

import (
	"ekgo/api/lib/plugins"
	"github.com/gin-gonic/gin"
)

//注册Hook
func RegisterHook(this *gin.Context, plugin *plugins.New) {
	/*plugin.Register(this, "/admin/menu", &hook.Test{})*/
}

//钩子
func Hook(this *gin.Context) {
	if this.Request.Method != "OPTIONS" {
		plugin := &plugins.New{}
		plugin.Init()
		RegisterHook(this, plugin)
		//处理之前请求Hook
		for _, plugin := range plugin.List {
			var result = plugin.Before()
			if result != nil {
				this.JSON(200, result)
				this.Abort()
				return
			}
		}
		//处理请求
		this.Next()
		//处理之后请求Hook
		go func() {
			for _, plugin := range plugin.List {
				plugin.After()
			}
		}()
	}
}
