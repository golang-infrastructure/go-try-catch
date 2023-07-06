package try_catch

import (
	"fmt"
)

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

// TryCatch 带捕获错误的执行函数
func TryCatch(f func()) (err error) {

	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("panic: %+v", r)
			}
		}
	}()

	// 执行函数
	f()

	return
}

// ------------------------------------------------ TryCatchReturn -----------------------------------------------------

// TryCatchReturn 带捕获错误的执行返回结果的函数
func TryCatchReturn[R any](f func() R) (result R, err error) {
	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// 执行函数
	result = f()

	return
}

func TryCatchReturn2[R1, R2 any](f func() (R1, R2)) (r1 R1, r2 R2, err error) {
	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// 执行函数
	r1, r2 = f()

	return
}

func TryCatchReturn3[R1, R2, R3 any](f func() (R1, R2, R3)) (r1 R1, r2 R2, r3 R3, err error) {
	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// 执行函数
	r1, r2, r3 = f()

	return
}

// ---------------------------------------------------------------------------------------------------------------------
