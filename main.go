package main

func main() {

	const interval int = 1000 // interval for random numbers
	const rate int = 50       // rate per generator
	const generators int = 1  // number of generators
	const lifetime int = 30   // data lifetime in seconds

	randoms := make(chan datablock, 100000)
	for i := 0; i < generators; i++ {
		go generator(randoms, interval, rate)
	}
	go dispatcher(randoms, lifetime)

	<-(chan int)(nil) // wait forever
}
