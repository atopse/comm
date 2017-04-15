package kind

// Kind 值类型定义
type Kind uint8

const (
	Any     Kind = iota //未确定的任意值
	String              // 字符串
	Number              // 整数
	Float               // 浮点数
	Boolean             // 布尔值

	Map   //Map[string]interface{}
	Table //表格数据
	JSON  //JSON格式数据
	XML   //XML格式数据
	CSV   //CSV格式数据

	IPv4     //常规IP
	Password //密码
	Email    //邮箱地址

)

var items = []string{
	"Any",

	"String",
	"Number",
	"Float",
	"Boolean",

	"Map",
	"Table",
	"JSON",
	"XML",
	"CSV",

	"IPv4",
	"Password",
	"Email",
}

func (k Kind) String() string {
	return items[k]
}

//TODO: 支持 IPv6
