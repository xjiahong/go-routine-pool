package routinepool

import (
	"fmt"
	"runtime/debug"
	"time"
)

// 链表消费模型,通过append添加数据列，
type RoutinePool[T any] struct {
	ch  chan any //控制并发执行的
	que chan T   //此字段为队列，可采用队列实现
}

// 新建池
// runCount 限制并发数
// maxQue 队列大小，超过的会等待
func NewRoutinePool[T func()](runCount int, maxQue int) *RoutinePool[T] {
	l := RoutinePool[T]{
		ch:  make(chan any, runCount),
		que: make(chan T, maxQue),
	}
	go func() { //循环处理
		for {
			l.ch <- time.Now()
			que := <-l.que
			if que != nil {
				go func() {
					//释放队列
					defer func() {
						<-l.ch
						if r := recover(); r != nil {
							fmt.Printf("异步执行出错:%s，%s", r, debug.Stack())
						}
					}()
					que() // 运行函数
				}()
			}
		}
	}()
	return &l
}

// 还剩余多少
func (l *RoutinePool[T]) Size() int {
	return len(l.que)
}

// 添加一个
func (l *RoutinePool[T]) Append(fun T) {
	l.que <- fun
}
