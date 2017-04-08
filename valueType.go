package comm

// ValueType 值类型定义
type ValueType uint8

const (
	VTAny          ValueType = iota //未确定的任意值
	VTSingleString                  // 字符串
	VTSingleInt                     // 整数
	VTSingleFloat                   // 浮点数
	VTBoolean                       // 布尔值
	VTIPv4                          //常规IP
	VTPassword                      //密码
	VTEmail                         //邮箱地址
	VTMap                           //Map[string]interface{}
	VTTable                         //表格数据
	VTJSON                          //JSON格式数据
	VTXML                           //XML格式数据
	VTCSV                           //CSV格式数据
)

//TODO: 支持 VTIPv6
