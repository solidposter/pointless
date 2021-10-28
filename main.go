package main

func main() {

	const interval int = 50000 // interval for random numbers
	const rate int = 50        // rate per generator
	const generators int = 10  // number of generators
	const lifetime int = 30    // data lifetime in seconds

	const inputQsize int = 100 // buffer size generators to dispatcher
	const pruneQsize int = 100 // buffer size maintainers to dispatcher

	go reporter(1)

	randoms := make(chan datablock, inputQsize)
	for i := 0; i < generators; i++ {
		go generator(randoms, interval, rate)
	}
	go dispatcher(randoms, lifetime, pruneQsize)

	<-(chan int)(nil) // wait forever
}
