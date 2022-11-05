package main

import (
	"flag"
	"time"
)

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

	intervalPtr := flag.Int("i", 5000, "range of generated random numbers")
	ratePtr := flag.Int("r", 50, "generator rate in numbers per second")
	generatorPtr := flag.Int("g", 10, "number of generators")
	lifetimePtr := flag.Int("l", 30, "maintainer timeout in seconds")

	inputqPtr := flag.Int("q", 100, "length of generator channel")
	pruneqPtr := flag.Int("p", 100, "length of prune channel")
	flag.Parse()

	go reporter(1)

	randoms := make(chan datablock, *inputqPtr)
	go dispatcher(randoms, *lifetimePtr, *pruneqPtr)
	for i := 0; i < *generatorPtr; i++ {
		go generator(randoms, *intervalPtr, *ratePtr)
		time.Sleep(time.Duration(1000*1000*1000 / *ratePtr) * time.Nanosecond)
	}

	<-(chan int)(nil) // wait forever
}
