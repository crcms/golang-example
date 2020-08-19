package main

import (
	"fmt"
	"time"
)

func main()  {
	// tick()

	// ticker()

	timer()
}

func timer()  {
	timer := time.NewTimer(time.Second * 3)
	//go func() {
	//	time.Sleep(time.Second * 12)
	//	// 关闭当时定时器
	//	// 如果之前定时器已被关闭，或过期，则返回false
	//	// 执行stop时，定时器的 chan Time并不会被关闭
	//	// 主要是为了防止多个goroutine竟争
	//	// Stop不会停止定时器。但是会停止Timer，停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误。
	//	timer.Stop()
	//}()
	// 关闭当时定时器
	// 如果之前定时器已被关闭，或过期，则返回false
	// 执行stop时，定时器的 chan Time并不会被关闭
	// 主要是为了防止多个goroutine竟争
	// Stop不会停止定时器。但是会停止Timer，停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误。
	defer timer.Stop()

	// 获取当前定时器n秒后的值
	v := <-timer.C
	fmt.Printf("current time: %v\n",v)

	// 必须要每次执行后都要进行重置这个time.Ticker不一样
	// 只能在已经停止或到期的定时器中调用Reset
	// 如果未从 timer.C 中接收到值，则调用无效，必须先停止计时器
	// 如果Stop报告计时器在停止之前已到期，则该通道将被清空
	// 如果当前定时器非活跃状态，或者已关闭或已过期，则会返回false
	timer.Reset(time.Second * 2)

	for {
		v1  := <-timer.C
		fmt.Printf("current time: %v\n", v1)
		timer.Reset(time.Second * 2)
	}
}

func ticker()  {
	// NewTicker会自动创建一个chan Time的通道
	// 然后按指定的 Duration 时间段向通道中发送数据
	// Duration 必须 >0 否则会panic
	ticker := time.NewTicker(time.Second)
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		done <- struct{}{}
		// Stop执行后，chan Time 不会关闭
		// 但不会再向chan中写入数据
		ticker.Stop()
	}()

	for {
		select {
		// 不断的从chan Time中收取当前时间
		case v := <-ticker.C:
			fmt.Printf("current time:%v\n",v)
		case <-done:
			fmt.Println("completed")
			goto COM
		}
	}
	COM:
}

func tick()  {
	// tick是 NewTicker 的简洁包装，返回的其实是 NewTicker.C 的一个<-channel
	// 并且这个channel是无法关闭的
	// 如果Duration给定的时间<0，则会返回nil
	// tick，垃圾回收机制是无法进行回收的
	tick := time.Tick(time.Second)
	for v := range tick {
		fmt.Printf("current time: %v\n",v)
	}
}