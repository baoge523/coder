package error_handle

import (
	"errors"
	"fmt"
	"testing"
)

var (
	// New创建的是errorString类型的错误
	oneErr   = errors.New("one error")
	twoErr   = errors.New("two error")
	threeErr = errors.New("two error")

	myError = &MyError{msg: "myError"}
)

// 自定义错误
type MyError struct {
	msg string
}

// 实现 error interface 接口: 表示错误类型
func (m *MyError) Error() string {
	return m.msg
}

// 通过fmt.Errorf("%w")包装错误, %w 占位符是一个实现了Error function struct pointer
// is 表示 当前error 是否是 target error tree 中的父节点（直接或者间接）
func TestIsError(t *testing.T) {

	// error tree   oneErr --> warpErr --> warp2Err --> warp4Err
	//                                 --> warp3Err
	warpErr := fmt.Errorf("warp oneError %w", oneErr)
	warp2Err := fmt.Errorf("warp warpError %w", warpErr)
	warp3Err := fmt.Errorf("warp warpError %w", warpErr)
	warp4Err := fmt.Errorf("warp warpError %w", warp2Err)

	if errors.Is(warpErr, oneErr) {
		fmt.Printf(" warpErr is oneErr\n")
	} else {
		fmt.Printf(" warpErr is not oneErr\n")
	}

	if errors.Is(warp2Err, oneErr) {
		fmt.Printf(" warp2Err is oneErr\n")
	} else {
		fmt.Printf(" warp2Err is not oneErr\n")
	}

	if errors.Is(warp2Err, warpErr) {
		fmt.Printf(" warp2Err is warpErr\n")
	} else {
		fmt.Printf(" warp2Err is not warpErr\n")
	}

	if errors.Is(warp3Err, warpErr) {
		fmt.Printf(" warp3Err is warpErr\n")
	} else {
		fmt.Printf(" warp3Err is not warpErr\n")
	}

	if errors.Is(warp3Err, warp2Err) {
		fmt.Printf(" warp3Err is warp2Err\n")
	} else {
		fmt.Printf(" warp3Err is not warp2Err\n")
	}

	if errors.Is(warp4Err, warp3Err) {
		fmt.Printf(" warp4Err is warp3Err\n")
	} else {
		fmt.Printf(" warp4Err is not warp3Err\n")
	}
}

// 判断一个error是否属于某一个错误类型
// errors.As(one,two), two需要是一个实现了Error()string 方法的指针对象的地址（指针的指针）
func TestAsError(t *testing.T) {
	var me *MyError // nil pointer

	// fmt.Printf("%v \n", me)  // nil
	// fmt.Printf("%v \n", &me) // 0x14000104048

	// 当传入的是指针，而不是指针地址时报错： second argument to errors.As must be a non-nil pointer to either a type that implements error, or to any interface type
	if errors.As(myError, &me) {
		fmt.Printf(" as ok\n")
	}
	warp1Err := fmt.Errorf("warp error: %w", myError)
	if errors.As(warp1Err, &me) {
		fmt.Printf(" as ok1\n")
	}
}

func TestUnwrap(t *testing.T) {
	warpErr := fmt.Errorf("warp oneError %w", oneErr)
	warp2Err := fmt.Errorf("warp warpError %w", warpErr)
	fmt.Printf("%s \n %v \n", warpErr.Error(), warpErr)
	fmt.Printf("%s \n %v \n", warp2Err.Error(), warp2Err)
	unwrap := errors.Unwrap(warpErr)
	un2wrap := errors.Unwrap(warp2Err)
	fmt.Printf("%s \n %v \n", unwrap.Error(), unwrap)
	fmt.Printf("%s \n %v \n", un2wrap.Error(), un2wrap)

	if un2wrap == warpErr {
		fmt.Printf(" === ")
	}
}

// join: 包装多个err，返回的err是包装的err的子节点，可以通过errors.Is判断
// join 是多个父一个子的包装，fmt.Error()是一父一子？？？ 可以多父一子，但是errors.Unwrap，然后nil
// join包装的不能使用errors.Unwrap，然后nil
func TestJoin(t *testing.T) {
	one := errors.New("one")
	two := errors.New("two")

	three := errors.Join(one)

	if errors.Is(three, one) {
		fmt.Printf("three one\n")
	}
	if errors.Is(three, two) {
		fmt.Printf("three two\n")
	}

	// In particular Unwrap does not unwrap errors returned by Join and returns nil
	unwrap := errors.Unwrap(three)

	fmt.Printf("%v\n", unwrap) // <nil>

}

func TestFmtError(t *testing.T) {
	one := errors.New("one error")
	two := errors.New("two error")

	three := fmt.Errorf("first: %w \n", one)
	fmt.Printf("%v", three)

	// 当使用fmt.Errorf包装多个错误是，不同通过errors.Unwrap，然后nil
	fore := fmt.Errorf("first: %w, second: %w \n", one, two)
	fmt.Printf("%v", fore)

	unwrap := errors.Unwrap(fore)

	fmt.Printf("%v", unwrap)

}
