package comm

import (
	"fmt"
)

// InvalidDataError 数据非法
type InvalidDataError struct {
	data string
	msg  string
}

func (e InvalidDataError) Error() string {
	return fmt.Sprintf("(%s)%s", e.data, e.msg)
}

// ErrDataIsEmpty 数据不能为空错误
func ErrDataIsEmpty(data string) InvalidDataError {
	return InvalidDataError{data: data, msg: "数据不能为空"}
}
