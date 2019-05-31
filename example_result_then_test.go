package denary64_test

import (
	"github.com/reiver/go-denary64"

	"fmt"
)

func ExampleResult_Then() {

	fn := func(value denary64.Type) denary64.Result {
		return denary64.Some(denary64.Uint64(9876543210))
	}

	var nothing denary64.Result = denary64.Nothing()
	var err  denary64.Result    = denary64.Error("This is an error.")
	var some denary64.Result    = denary64.Some(denary64.Uint64(123))

	fmt.Printf("nothing = %#v\n", nothing)
	fmt.Printf("err     = %#v\n", err)
	fmt.Printf("some    = %#v\n", some)

	fmt.Println("---")

	fmt.Printf("nothing.Then(fn) = %#v\n", nothing.Then(fn))
	fmt.Printf("err.Then(fn)     = %#v\n", err.Then(fn))
	fmt.Printf("some.Then(fn)    = %#v\n", some.Then(fn))

	// Output:
	// nothing = denary64.Nothing()
	// err     = denary64.Error("This is an error.")
	// some    = denary64.Some(denary64.Uint64(123))
	// ---
	// nothing.Then(fn) = denary64.Nothing()
	// err.Then(fn)     = denary64.Error("This is an error.")
	// some.Then(fn)    = denary64.Some(denary64.Uint64(9876543210))
}
