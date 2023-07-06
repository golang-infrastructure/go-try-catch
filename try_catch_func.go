package try_catch

import (
	"fmt"
)

// ------------------------------------------------- --------------------------------------------------------------------

func IgnoreLastError2[V1 any](v1 V1, err error) V1 {
	return v1
}

func IgnoreLastError3[V1, V2 any](v1 V1, v2 V2, err error) (V1, V2) {
	return v1, v2
}

func IgnoreLastError4[V1, V2, V3 any](v1 V1, v2 V2, v3 V3, err error) (V1, V2, V3) {
	return v1, v2, v3
}

// ------------------------------------------------- --------------------------------------------------------------------

func toError(v any) error {
	err, ok := v.(error)
	if !ok {
		err = fmt.Errorf("panic: %+v", v)
	}
	return err
}

// ------------------------------------------------- --------------------------------------------------------------------

// TryCatch 带捕获错误的执行函数
func TryCatch(f func()) (err error) {

	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			err = toError(r)
		}
	}()

	// 执行函数
	f()

	return
}

// TryCatchIgnore 捕捉panic，但是忽略错误
func TryCatchIgnore(f func()) {
	_ = TryCatch(f)
}

// ------------------------------------------------ TryCatchReturn -----------------------------------------------------

// TryCatchReturn 带捕获错误的执行返回结果的函数
func TryCatchReturn[R any](f func() R) (result R, err error) {
	// 来个try-CatchHandler
	defer func() {
		if r := recover(); r != nil {
			err = toError(r)
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
			err = toError(r)
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
			err = toError(r)
		}
	}()

	// 执行函数
	r1, r2, r3 = f()

	return
}

// ---------------------------------------------------------------------------------------------------------------------
