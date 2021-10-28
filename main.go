package main

func main() {

	const interval int = 100000 // interval for random numbers
	const rate int = 50         // rate per generator
	const generators int = 100  // number of generators
	const lifetime int = 30     // data lifetime in seconds

	go reporter(1)

	randoms := make(chan datablock, 100)
	for i := 0; i < generators; i++ {
		go generator(randoms, interval, rate)
	}
	go dispatcher(randoms, lifetime)

	<-(chan int)(nil) // wait forever
}
