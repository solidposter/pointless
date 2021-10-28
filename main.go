package main

func main() {

	const interval int = 20000 // interval for random numbers
	const rate int = 50        // rate per generator
	const generators int = 20  // number of generators
	const lifetime int = 30    // data lifetime in seconds

	go reporter(1)

	randoms := make(chan datablock, 10)
	for i := 0; i < generators; i++ {
		go generator(randoms, interval, rate)
	}
	go dispatcher(randoms, lifetime)

	<-(chan int)(nil) // wait forever
}
