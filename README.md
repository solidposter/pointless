Pointless program that plays with goroutines and channels.
I generate random numbers, and use a goroutine to manage the random number.

generator(s) generate random numbers, encapsulates in a struct and inserts them into a channel.

dispatcher reads that channel, and for each random number starts a maintainer goroutine to manage it.

the manager keeps track of updates for that number, and if a timeout expires it will message the dispatcher, and upon acknowledgement kill the goroutine.

Update the constants in main.go to play around. Note that some OS (OpenBSD, Windows, more?) can't handle a ticker rate faster than 50 per second.

----

On my OpenBSD test server:

        const interval int = 1000000 // interval for random numbers
        const rate int = 50          // rate per generator
        const generators int = 200   // number of generators
        const lifetime int = 30      // data lifetime in seconds

Gives me around 245k active goroutines once it stabilizes.

