package main

import "runtime"

func main() {

	//开启线程
	go sayHello("好")
	sayHello("你")

}

func sayHello(str string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		println(str)
	}
}
