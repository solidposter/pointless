package main

func main() {

	randoms := make(chan datablock, 100000)
	go generator(randoms, 1000, 50)
	go dispatcher(randoms, 30)

	<-(chan int)(nil) // wait forever
}
