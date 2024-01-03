package GoSamplePackage

import "fmt"

func the_foo() {
	fmt.Println("go foo")

}

func echo() {
	the_foo()
	fmt.Println("echo")
}
