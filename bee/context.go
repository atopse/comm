package bee

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Context 扩展Beego的Context
type Context struct {
	context.Context
}

// JSON 输出JSON数据到Http Response。
// 当data为Error类型时，将直接输出{error:"Error信息"}
func (ctx *Context) JSON(data interface{}, httpCode ...int) {
	var (
		hasEncoding = false
		hasIndent   = beego.BConfig.RunMode == "dev"
	)
	if len(httpCode) > 0 {
		ctx.Output.SetStatus(httpCode[0])
	}
	if err, ok := data.(error); ok {
		if len(httpCode) == 0 {
			ctx.Output.SetStatus(500)
		}
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.Output.JSON(data, hasIndent, hasEncoding)
}
