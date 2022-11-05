Pointless program that plays with goroutines and channels.
It generates random numbers, and then manages them by creating a goroutine per number.

The generator(s) generate random numbers, encapsulates them in a struct and inserts them into a channel.
The dispatcher reads that channel, and for each random number starts a maintainer goroutine to manage it.
The maintainer waits for an update or a timeout. If the maintainer gets an update, which happens when the number that is managed is generated again, it will reset the timeout and again wait for update or timeout. If the timeout expires the maintainer will signal the dispatcher that it will shut down, and upon acknowledgement exit the goroutine.

Pointless is configured with command line options.

nacho$ pointless -h  
Usage of pointless:  
  -g int  
        number of generators (default 10)  
  -i int  
        range of generated random numbers (default 5000)  
  -l int  
        maintainer timeout in seconds (default 30)  
  -p int  
        length of prune channel (default 100)  
  -q int  
        length of generator channel (default 100)  
  -r int  
        generator rate in numbers per second (default 50)  
nacho$  

Example run on my old quad-core server, running OpenBSD7.2-current with go1.19.2:

nacho$ pointless -i 10000000 -q 100000 -p 100000 -g 3400  
... 30 minutes or so to stabilize, let GC hurt a little, and so on.  
Dispatcher values received: 168983  
Dispatcher prunes received: 102453  
Dispatcher map size: 3960029  
randomQblocks: 0  
pruneQblocks:  0  

Around 4M goroutines on an old server running OpenBSD, not bad.

Note that when running OpenBSD the generators cannot exceed 50 values per second,
this is not an issue with Linux or FreeBSD, with Linux being the king of timers.

