package module1

import (
	"fmt"
	"github.com/rivenhl1/playground/module2"
)


func Hello(){
	fmt.Println("Hello from Module1 public")
	module2.Hello()
	// hello()
}

func hello(){
	fmt.Println("Hello from Module1 internal")
}