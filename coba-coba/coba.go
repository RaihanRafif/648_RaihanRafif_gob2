package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println(dt.Format("15:04:05"))
}
