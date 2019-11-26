// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字
package main

import (
	"fmt"
	"os"
	"strings"
)

// os.Args[0]即为程序所在路径，[1:end]为参数
func main() {
	fmt.Println(strings.Join(os.Args[0:2], "add"))
}
