package main

import (
	"fmt"
	"testing"
	"time"
)

/**
参考资料：https://www.kancloud.cn/kancloud/the-way-to-go/165091
*/

func main() {
	//chan_sendGet()
}

/**
select 做的就是：选择处理列出的多个通信情况中的一个。

如果都阻塞了，会等待直到其中一个可以处理
如果多个可以处理，随机选择一个
如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
*/
func Test_Select(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "a"
		ch1 <- "b"
	}()
	go func() {
		ch2 <- "c"
		ch2 <- "d"
	}()
	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println("from ch1:", v)
			case v := <-ch2:
				fmt.Println("from ch2:", v)
			default:
				fmt.Println("nothing")
			}
		}
	}()

	time.Sleep(1e9)

}

func Test_close(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "a"
		ch <- "b"
		defer close(ch)
	}()
	fmt.Println(<-ch)
}

/**
for 循环的 range 语句可以用在通道 ch 上，便可以从通道中获取值，像这样
*/
func Test_for(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
	time.Sleep(1e9)
}

/**
默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的：

1）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

尽管这看上去是非常严格的约束，实际在大部分情况下工作的很不错。
下面这个demo : 一个协程在无限循环中给通道发送整数数据。有几个接收者，ch里就放多少值
*/
func chan_block() {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("set chan value:", i)
			ch1 <- i
		}
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

/**
interface{} 表示可以接收任意类型
*/
func chan_interface() {
	ch := make(chan interface{})
	go func() {
		ch <- 8
		ch <- "a"
		ch <- 1.5
	}()
	go func() {
		for {
			fmt.Println(<-ch)
		}
		fmt.Println("end")
	}()

	time.Sleep(1e9)
}

/**
一个简单的chanl 接收和发送者
*/
func chan_sendGet() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

/**
默认情况下，通信是同步且无缓冲的。这种特性导致通道的发送/接收操作，在对方准备好之前是阻塞的。

对于同一个通道，发送操作在接收者准备好之前是阻塞的。如果通道中的数据无人接收，就无法再给通道传入其他数据。新的输入无法在通道非空的情况下传入，所以发送操作会等待channel再次变为可用状态，即通道值被接收后。
对于同一个通道，接收操作是阻塞的，直到发送者可用。如果通道中没有数据，接收者就阻塞了。
————————————————
原文链接：https://blog.csdn.net/benben_2015/article/details/84842486
*/
func chan_deadLock1() {
	ch := make(chan string)
	ch <- "a"
	ch <- "b"
	fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}

func chanDeadLock_Resovled1() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
	}()
	fmt.Println(<-ch1)

	//或者

	/*ch1 := make(chan string, 1) //此时的ch1通道可以称为缓冲通道，在缓冲满载(缓冲被全部使用)之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。定义方法如：ch:=make(chan type, value)。
	ch1 <- "hello world"
	fmt.Println(<-ch1)*/
}
