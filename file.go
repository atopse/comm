package comm

import (
	"os"
	"path/filepath"
	"runtime"

	"qiniupkg.com/x/log.v7"
)

// SearchFile from the runtime
func SearchFile(name string) (string, error) {
	if name == "" {
		return "", os.ErrNotExist
	}
	dir, _ := os.Getwd()
	if dir == "" {
		_, file, _, _ := runtime.Caller(1)
		dir = filepath.Dir(file)
	}
	name = filepath.Base(name)
	var fileName string
	for {
		fileName = filepath.Join(dir, name)
		log.Debug("FindeFile:", fileName)
		if _, err := os.Stat(fileName); err == nil {
			return fileName, nil
		} else if os.IsNotExist(err) == false {
			return "", err
		}
		// 继续往上一级目录查找
		dir = filepath.Dir(dir)
		if dir == filepath.VolumeName(dir) {
			return "", os.ErrNotExist
		}
		if dir == "/" || dir == "" || dir == fileName {
			return "", os.ErrNotExist
		}
	}
}
