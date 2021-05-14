// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	pid := os.Getpid()

	fmt.Printf("进程 PID: %d \n", pid)

	gracefullyStop()
}



func simpleStop(){
	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//创建一个bool channel
	done := make(chan bool, 1)

	//注册要接收的信号，
	//1.syscall.SIGINT:接收ctrl+c
	//2.syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	//程序将在此处等待，直到它预期信号（如Goroutine所示）
	//在“done”上发送一个值，然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func gracefullyStop(){
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL:
				fmt.Println("Program Exit...", s)
				exitFunc()
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("Program Start...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func exitFunc() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}