package try_catch

import "errors"

// ------------------------------------------------ ---------------------------------------------------------------------

// TryHandler try块执行
type TryHandler struct {

	// 要执行的方法
	handler func()

	// 用来捕捉代码块的CatchHandler
	bindingCatchHandlerSlice []*CatchHandler

	// 会首先使用普通的来匹配，如果匹配不上则使用默认的匹配
	bindingDefaultCatchHandler *DefaultCatchHandler

	// 当没有发生异常的时候执行的代码
	bindingTryCatchElseHandler *TryCatchElseHandler

	// 无论是否发生异常都会执行的代码
	finallyHandler *FinallyHandler
}

// Try 创建一个TryHandler，可以看做是它的一个构造方法
func Try(funcToTry func()) *TryHandler {
	return &TryHandler{
		handler: funcToTry,
	}
}

func (x *TryHandler) Catch(err error, handler func(err error)) *CatchHandler {
	return NewCatchHandler(x, err, handler)
}

func (x *TryHandler) Finally(handler func()) *FinallyHandler {
	return NewFinallyHandler(x, handler)
}

func (x *TryHandler) DefaultCatch(handler func(err error)) *DefaultCatchHandler {
	return NewDefaultCatchHandler(x, handler)
}

func (x *TryHandler) Else(handler func()) *TryCatchElseHandler {
	return NewTryCatchElseHandler(x, handler)
}

// Do 开始执行整个流程
func (x *TryHandler) Do() {

	// 无论是否执行成功，在最后退出的时候都要执行finally
	defer func() {
		if x.finallyHandler != nil {
			x.finallyHandler.handle()
		}
	}()

	// 执行函数，尝试捕获错误
	err := TryCatch(x.handler)
	if err == nil {
		// 如果乜有捕获到，看下是否有设置else，设置了的话就调用下
		if x.bindingTryCatchElseHandler != nil {
			x.bindingTryCatchElseHandler.handle()
		}
		return
	}

	// 走到这里说明发生了错误了，则看下设置的CatchHandler是否能够捕获得到这个异常
	catchSuccess := false
	for _, catchHandler := range x.bindingCatchHandlerSlice {
		if catchHandler.match(err) {
			catchSuccess = true
			catchHandler.handle(err)
			break
		}
	}

	// 如果前面的捕获不到则走这里的Default，没设置的话就是不需要捕获了
	if !catchSuccess && x.bindingDefaultCatchHandler != nil {
		x.bindingDefaultCatchHandler.handle(err)
	}
}

// ------------------------------------------------ CatchHandler -------------------------------------------------------

// CatchHandler 错误匹配
type CatchHandler struct {
	// 每个错误匹配绑定到一个TryHandler
	bindingTryHandler *TryHandler
	err               error
	handler           func(err error)
}

func NewCatchHandler(bindingTryHandler *TryHandler, err error, handler func(err error)) *CatchHandler {
	x := &CatchHandler{
		bindingTryHandler: bindingTryHandler,
		err:               err,
		handler:           handler,
	}
	bindingTryHandler.bindingCatchHandlerSlice = append(bindingTryHandler.bindingCatchHandlerSlice, x)
	return x
}

func (x *CatchHandler) match(err error) bool {
	return errors.Is(err, x.err)
}

func (x *CatchHandler) handle(err error) {
	x.handler(err)
}

func (x *CatchHandler) Catch(err error, handler func(err error)) *CatchHandler {
	return NewCatchHandler(x.bindingTryHandler, err, handler)
}

func (x *CatchHandler) DefaultCatch(handler func(err error)) *DefaultCatchHandler {
	return NewDefaultCatchHandler(x.bindingTryHandler, handler)
}

func (x *CatchHandler) Else(handler func()) *TryCatchElseHandler {
	return NewTryCatchElseHandler(x.bindingTryHandler, handler)
}

func (x *CatchHandler) Finally(handler func()) *FinallyHandler {
	return NewFinallyHandler(x.bindingTryHandler, handler)
}

func (x *CatchHandler) Do() {
	x.bindingTryHandler.Do()
}

// ------------------------------------------------ DefaultCatchHandler ------------------------------------------------

// DefaultCatchHandler 默认的TryCatch分支
type DefaultCatchHandler struct {
	// 每个错误匹配绑定到一个TryHandler
	bindingTryHandler *TryHandler
	//
	handler func(err error)
}

func NewDefaultCatchHandler(bindingTryHandler *TryHandler, handler func(err error)) *DefaultCatchHandler {
	x := &DefaultCatchHandler{
		bindingTryHandler: bindingTryHandler,
		handler:           handler,
	}
	bindingTryHandler.bindingDefaultCatchHandler = x
	return x
}

func (x *DefaultCatchHandler) handle(err error) {
	x.handler(err)
}

func (x *DefaultCatchHandler) Else(handler func()) *TryCatchElseHandler {
	return NewTryCatchElseHandler(x.bindingTryHandler, handler)
}

func (x *DefaultCatchHandler) Finally(handler func()) *FinallyHandler {
	return NewFinallyHandler(x.bindingTryHandler, handler)
}

func (x *DefaultCatchHandler) Do() {
	x.bindingTryHandler.Do()
}

// ------------------------------------------------ TryCatchElseHandler ------------------------------------------------

// TryCatchElseHandler 在try执行块未发生错误时执行
type TryCatchElseHandler struct {

	// 每个try-catch-else绑定到一个TryHandler上
	bindingTryHandler *TryHandler

	// 此else块的代码
	handler func()
}

func NewTryCatchElseHandler(bindingTryHandler *TryHandler, handler func()) *TryCatchElseHandler {
	x := &TryCatchElseHandler{
		bindingTryHandler: bindingTryHandler,
		handler:           handler,
	}
	bindingTryHandler.bindingTryCatchElseHandler = x
	return x
}

func (x *TryCatchElseHandler) handle() {
	x.handler()
}

func (x *TryCatchElseHandler) Finally(handler func()) *FinallyHandler {
	return NewFinallyHandler(x.bindingTryHandler, handler)
}

func (x *TryCatchElseHandler) Do() {
	x.bindingTryHandler.Do()
}

// ------------------------------------------------ FinallyHandler -----------------------------------------------------

// FinallyHandler 始终会执行
type FinallyHandler struct {
	bindingTryHandler *TryHandler
	handler           func()
}

func NewFinallyHandler(bindingTryHandler *TryHandler, handler func()) *FinallyHandler {
	x := &FinallyHandler{
		bindingTryHandler: bindingTryHandler,
		handler:           handler,
	}
	bindingTryHandler.finallyHandler = x
	return x
}

func (x *FinallyHandler) handle() {
	x.handler()
}

func (x *FinallyHandler) Do() {
	x.bindingTryHandler.Do()
}

// ---------------------------------------------------------------------------------------------------------------------
