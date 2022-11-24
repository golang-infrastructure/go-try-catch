package example

import (
	"errors"
	try_catch "github.com/golang-infrastructure/go-try-catch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryCatch(t *testing.T) {

	var errFoo = errors.New("foo")

	// 正常执行
	err := try_catch.TryCatch(func() {
		t.Log("ok")
	})
	assert.Nil(t, err)

	// 执行时发生panic
	err = try_catch.TryCatch(func() {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}


func TestTryCatchReturn(t *testing.T) {

	var errFoo = errors.New("foo")

	// 正常执行
	v, err := try_catch.TryCatchReturn(func() int {
		return 10086
	})
	assert.Nil(t, err)
	assert.Equal(t, 10086, v)

	// 执行时发生panic
	v, err = try_catch.TryCatchReturn(func() int {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}
