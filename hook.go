package comm

//hook function to run
type hookfunc func() error

var (
	hooks      = make(map[int][]hookfunc, 0) //hook function slice to store the hookfunc
	maxPriotiy = 10
)

// RegAPPInitHook 将需要在程序启动时运行的方法加入，该Hook将在执行RunHook时按优先级执行。
// priority 范围 0-10；执行顺序为：0->10.
func RegAPPInitHook(hf hookfunc, priority ...int) {
	p := maxPriotiy
	if len(priority) > 0 {
		p = priority[0]
		if p < 0 {
			p = 0
		} else if p > maxPriotiy {
			p = maxPriotiy
		}
	}
	hooks[p] = append(hooks[p], hf)
}

// RunAPPInitHook 执行Hook
func RunAPPInitHook() error {
	if len(hooks) == 0 {
		return nil
	}
	for i := 0; i <= maxPriotiy; i++ {
		fs := hooks[i]
		for _, f := range fs {
			if err := f(); err != nil {
				return err
			}
		}
	}
	return nil
}
