Pointless program that plays with goroutines and channels.
I generate random numbers, and use a goroutine to manage the random number.

generator(s) generate random numbers, encapsulates in a struct and inserts them into a channel.

dispatcher reads that channel, and for each random number starts a maintainer go-routine to manage it.

the manager keeps track of updates for that number, and if a time-out expires it will message the dispatcher, and upon acknowledgement kill the go-routine.


Change the generator values to change the rate and the range of random numbers. For OS's where ticker can't run at high rate, add start more generators.

go generator(randoms, 1000, 50)

1000 is the random number range, 50 is the rate.

