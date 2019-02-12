package main

import (
	"fmt"
	"time"
)

func main() {
	unix_time := time.Now().Unix()
	fmt.Println(unix_time)
	fmt.Println(unix_time - unix_time%60)
}
