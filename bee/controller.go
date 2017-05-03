package bee

import "github.com/astaxie/beego"

// Controller 可扩展Controller
type Controller struct {
	beego.Controller
}

// JSON 输出JSON数据到Http Response。
// 当data为Error类型时，将直接输出{error:"Error信息"}
func (c *Controller) JSON(data interface{}, httpCode ...int) {
	var (
		hasEncoding = false
		hasIndent   = beego.BConfig.RunMode == "dev"
	)
	if len(httpCode) > 0 {
		c.Ctx.Output.SetStatus(httpCode[0])
	}
	if err, ok := data.(error); ok {
		if len(httpCode) == 0 {
			c.Ctx.Output.SetStatus(500)
		}
		c.JSON(map[string]string{"error": err.Error()})
		return
	}
	c.Ctx.Output.JSON(data, hasIndent, hasEncoding)
}
