package main

import (
"context"
"fmt"
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Work Work内包含一些耗时的处理，可能是密集计算或者外部IO
	Work()
}


// IProducer 请勿修改接口
type IProducer interface {
	// Produce Produce每次调用会返回一个IWorkload实例
	// 当返回nil时表示已经生产完毕
	Produce() IWorkload
}

// 问题2：请编写函数Question2的实现如下功能
// 该函数输入一个IProducer实例，每次调用其Produce()方法会返回一个IWorkload实例。
// 1. 请反复调用该Produce()方法，直到返回nil，表明没有更多IWorkload。
//    此间可能会生产大量IWorkload实例，数目在此未知。
// 2. 对每个生产出的IWorkload实例，请调用一次它的Work()方法。
//    Work()内包含一些耗时的处理，可能是密集计算或者外部IO。
// 3. 请并发调用多个IWorkload的Work()方法，最多允许5个并发的Work()执行。
//    单个并发的实现，或并发数超过5的限制，都不能得分。
//
// 提示：请最小化内存、CPU代价
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

func Question2(producer IProducer) {
	// ====== 在这里书写代码 ====== //
	var (
		in          = make(chan IWorkload, 1)
		finished    = make(chan struct{}, 5)
		ctx, cancel = context.WithCancel(context.Background())
	)
	for i := 0; i < 5; i++ { // 并发调用work
		go run(ctx, in, finished)
		finished <- struct{}{}
	}
	for {
		iWorkload := producer.Produce()
		if iWorkload == nil {
			for i := 0; i < 5; i++ { // 所有任务完成后退出
				<-finished
				fmt.Println("wanchengla", i)
			}

			cancel() // 任务完成后，取消协程上下游工作并返回
			return
		}
		<-finished
		in <- iWorkload
	}
}

func run(ctx context.Context, in chan IWorkload, finished chan struct{}) {
	for {
		select {
		case <-ctx.Done(): // 利用ctx,收到上游协程关闭，就return，释放内存
			return
		case p := <-in:
			p.Work()
			finished <- struct{}{}
		default:
		}
	}
}
