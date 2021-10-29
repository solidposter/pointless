package main

//
// Copyright (c) 2021 Tony Sarendal <tony@polarcap.org>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

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
