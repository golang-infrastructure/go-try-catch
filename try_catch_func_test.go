package try_catch

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryCatch(t *testing.T) {

	var errFoo = errors.New("foo")

	// 正常执行
	err := TryCatch(func() {
		t.Log("ok")
	})
	assert.Nil(t, err)

	// 执行时发生panic
	err = TryCatch(func() {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}

func TestTryCatchReturn(t *testing.T) {

	var errFoo = errors.New("foo")

	// 正常执行
	v, err := TryCatchReturn(func() int {
		return 10086
	})
	assert.Nil(t, err)
	assert.Equal(t, 10086, v)

	// 执行时发生panic
	v, err = TryCatchReturn(func() int {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}

func TestTryCatchReturn2(t *testing.T) {
	var errFoo = errors.New("foo")

	// 正常执行
	v1, v2, err := TryCatchReturn2(func() (int, string) {
		return 10086, "10010"
	})
	assert.Nil(t, err)
	assert.Equal(t, 10086, v1)
	assert.Equal(t, "10010", v2)

	// 执行时发生panic
	v1, v2, err = TryCatchReturn2(func() (int, string) {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}

func TestTryCatchReturn3(t *testing.T) {
	var errFoo = errors.New("foo")

	// 正常执行
	v1, v2, v3, err := TryCatchReturn3(func() (int, string, float64) {
		return 10086, "10010", 3.14
	})
	assert.Nil(t, err)
	assert.Equal(t, 10086, v1)
	assert.Equal(t, "10010", v2)
	assert.Equal(t, 3.14, v3)

	// 执行时发生panic
	v1, v2, v3, err = TryCatchReturn3(func() (int, string, float64) {
		panic(errFoo)
	})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errFoo)
}

func TestTryCatchStringPanic(t *testing.T) {
	Try(func() {
		panic("string")
	}).DefaultCatch(func(err error) {
		fmt.Println(err)
	}).Do()
}
