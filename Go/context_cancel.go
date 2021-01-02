package main

import (
	"context" 
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
 
	go watch(ctx, "监控-1")
	go watch(ctx, "监控-2") 

	fmt.Println("现在准备开始等待 8 秒, time=", time.Now().Unix())
	time.Sleep(time.Second * 8)

	fmt.Println("已运行 8 秒，准备调用下一行的 cancel() 函数，发现 2 个子协程已经结束，time=", time.Now().Unix())
	cancel() 
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到结束信号，退出监控，time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中，time=", time.Now().Unix())
			time.Sleep(time.Second * 1)
		}
	}
}

// 现在准备开始等待 8 秒, time= 1605082007 
// 监控-2 goroutine监控中，time= 1605082007 
// 监控-1 goroutine监控中，time= 1605082007
// 监控-1 goroutine监控中，time= 1605082008 
// 监控-2 goroutine监控中，time= 1605082008 
// 监控-2 goroutine监控中，time= 1605082009
// 监控-1 goroutine监控中，time= 1605082009
// 监控-2 收到结束信号，退出监控，time= 1605082010
// 监控-1 收到结束信号，退出监控，time= 1605082010
// 已运行 8 秒，准备调用下一行的 cancel() 函数，发现 2 个子协程已经结束，time= 1605082015

