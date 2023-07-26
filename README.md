# co
coroutine library

A coroutine library to implement deterministic scheduling of concurrent processes.

Mostly inspired by Rob Pike's recursive state function type in his [Lexical Scanning in Go](https://go.dev/talks/2011/lex.slide#1) talk.

I use this pattern a lot in the [sb library](https://github.com/reusee/sb), modelling stream/sink process as recursive functions. 

## Roadmap

* I/O waiting. channel polling
* threaded scheduler
* prioritized scheduling

