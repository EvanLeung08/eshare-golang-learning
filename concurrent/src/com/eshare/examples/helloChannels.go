package main

func main() {
	data := []int{1, 2, 3, 4, 5}
	chanData := make(chan int)

	go sum(data[:len(data)/2], chanData)
	go sum(data[len(data)/2:], chanData)

	x, y := <-chanData, <-chanData
	println(x, y, x+y)

}

func sum(a []int, c chan int) {
	total := 0

	for _, v := range a {
		total += v
	}
	c <- total
}
