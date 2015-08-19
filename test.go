// test
package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
)

func main() {
	var x float32;
	fmt.Println("Hello World!")
	x=3.4;
	y:=33;
	fmt.Println(x,y)
	t := time.Now()
	fmt.Print(t)
	
	r := bufio.NewReader(os.Stdin)
	r.ReadLine()
}
