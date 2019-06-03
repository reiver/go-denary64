package denary64_test

import (
	"github.com/reiver/go-denary64"

	"fmt"
)

func ExampleResult_Return_nothing() {

	var result denary64.Result = denary64.Nothing()

	fmt.Printf("result = %#v\n", result)

	value, err := result.Return()
	if nil != err {
		fmt.Printf("err: %q", err)
		return
	}

	fmt.Printf("value = %s", value)

	// Output:
	// result = denary64.Nothing()
	// value = ⟨undefined⟩
}

func ExampleResult_Return_error() {

	var result denary64.Result = denary64.Error("Uh oh!")

	fmt.Printf("result = %#v\n", result)

	value, err := result.Return()
	if nil != err {
		fmt.Printf("err: %q", err)
		return
	}

	fmt.Printf("value = %s", value)

	// Output:
	// result = denary64.Error("Uh oh!")
	// err: "Uh oh!"
}

func ExampleResult_Return_value() {

	var result denary64.Result = denary64.Something( denary64.Uint64(123) )

	fmt.Printf("result = %#v\n", result)

	value, err := result.Return()
	if nil != err {
		fmt.Printf("err: %q", err)
		return
	}

	fmt.Printf("value = %s", value)

	// Output:
	// result = denary64.Something(denary64.Uint64(123))
	// value = 123
}
