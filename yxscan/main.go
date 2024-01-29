package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	parse()
	fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
