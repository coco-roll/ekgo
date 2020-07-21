package plugins

import (
	"ekgo/api/lib/response"
	"github.com/gin-gonic/gin"
)

//定义一个接口，里面有两个方法
type pluginfunc interface {
	Before() *response.Write //在什么之前执行
	After()                  //在什么之后执行
}

//定义一个类，来存放我们的插件
type New struct {
	List map[string]pluginfunc
}

//初始化插件
func (p *New) Init() {
	p.List = make(map[string]pluginfunc)
}

//注册插件
func (p *New) Register(this *gin.Context, name string, plugin pluginfunc) {
	if this.Request.URL.Path == name {
		p.List[name] = plugin
	}
}
