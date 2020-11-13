package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 基于信号量的通信

// 需要注意的是signal在接收 chan os.Signal 的通道数据时，并不会因为 chan 满了而阻塞
// 因此signal.Notify函数的调用方法应该确保 signal chan 有足够的空间去接收新的信号值
// 有个更好的办法是创建一个缓冲为1的chan，准备时刻接收signal信道

// SIGKILL和SIGSTOP是不能自行处理的，只能接受操作系统的默认执行

// signal.Stop() 函数会取消之前自定义的信号处理，而使用系统默认处理，可以多次重复调用
func main()  {
	// 创建一个信号chan
	sig := make(chan os.Signal,1)
	sigs := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	// 接收通知信号
	signal.Notify(sig,sigs...)
	// 处理接收的信号 sig是chan 阻塞获取
	for v := range sig {
		fmt.Printf("当前接收的信号：%s\n",v)

		// 关闭当前传输的信号，使用系统默认的
		signal.Stop(sig)
		close(sig) //关闭通道

		fmt.Printf("还原为系统默认的信号：%s\n",v)
	}
}
