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
	if e.data == "" {
		return e.msg
	}
	return fmt.Sprintf("(%s)%s", e.data, e.msg)
}

// ErrDataIsEmpty 数据不能为空错误
func ErrDataIsEmpty(data ...string) InvalidDataError {
	err := InvalidDataError{msg: "不能为空"}
	if len(data) > 0 {
		err.data = data[0]
	}
	return err
}
