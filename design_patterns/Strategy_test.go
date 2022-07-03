package design_patterns

import (
	"fmt"
	"testing"
)

func TestStrategy1_Do(t *testing.T) {

	context := Context{Strategy1{}}
	context.Invoke()

	fmt.Println(" ================= ")

	context = Context{Strategy2{}}
	context.Invoke()
}
