package main

func main() {

	chanData := make(chan int, 2)
	chanData <- 2
	chanData <- 3
	println(chanData)

}
