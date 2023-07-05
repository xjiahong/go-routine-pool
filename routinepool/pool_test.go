package routinepool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	routinePool := NewRoutinePool(10, 10000)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), ":start")
	count := 1000
	for i := 0; i < count; i++ {
		i := i
		routinePool.Append(func() {
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), ":执行>", i)
		})
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), ":ok")

	fmt.Println("后面代码会报错死锁，无视即可，主要是测试上面的协程池")
	select {}
}
