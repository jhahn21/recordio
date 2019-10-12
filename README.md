# recordio

Package recordio implements a file format for a sequence of variable-length data.

[![Build Status](https://travis-ci.org/jhahn21/recordio.svg?branch=master)](https://travis-ci.org/jhahn21/recordio)
[![GoDoc](https://godoc.org/github.com/jhahn21/recordio?status.svg)](https://godoc.org/github.com/jhahn21/recordio)
[![Go Report Card](https://goreportcard.com/badge/github.com/jhahn21/recordio)](https://goreportcard.com/report/github.com/jhahn21/recordio)

## Benchmarks

On `Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz`

```
goos: darwin
goarch: amd64
pkg: github.com/jhahn21/recordio
BenchmarkRead/nocompression_100           11585869          99.5 ns/op        32 B/op        1 allocs/op
BenchmarkRead/nocompression_1000          10453348         111 ns/op        32 B/op        1 allocs/op
BenchmarkRead/nocompression_10000          6116316         195 ns/op        32 B/op        1 allocs/op
BenchmarkRead/snappy_100                   4472151         267 ns/op        32 B/op        1 allocs/op
BenchmarkRead/snappy_1000                  3407258         357 ns/op        32 B/op        1 allocs/op
BenchmarkRead/snappy_10000                 1000000        1019 ns/op        32 B/op        1 allocs/op
BenchmarkRead/zlib_100                      834781        1426 ns/op      4228 B/op        4 allocs/op
BenchmarkRead/zlib_1000                     571678        1857 ns/op      4228 B/op        4 allocs/op
BenchmarkRead/zlib_10000                    193426        5881 ns/op      4228 B/op        4 allocs/op
BenchmarkWrite/nocompression_100          73785631          14.2 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/nocompression_1000         83530110          14.1 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/nocompression_10000        82618550          14.1 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/snappy_100                  6451332         183 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/snappy_1000                 2747001         438 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/snappy_10000                 465871        2446 ns/op         0 B/op        0 allocs/op
BenchmarkWrite/zlib_100                      37065       32426 ns/op        21 B/op        0 allocs/op
BenchmarkWrite/zlib_1000                     14058       82467 ns/op        57 B/op        0 allocs/op
BenchmarkWrite/zlib_10000                     5492      193872 ns/op       150 B/op        0 allocs/op
PASS
ok    github.com/jhahn21/recordio 24.716s
```

## How to contribute

We use [Gerrit](https://review.gerrithub.io) for reviewing proposed changes.

```
$ git clone https://github.com/jhahn21/recordio.git
$ cd recordio
$ curl -Lo .git/hooks/commit-msg https://review.gerrithub.io/tools/hooks/commit-msg
$ chmod +x .git/hooks/commit-msg
$ git remote add gerrit https://review.gerrithub.io/jhahn21/recordio
...

// Push for code review
...
$ git commit
$ git push gerrit HEAD:refs/for/master
...

// Push a patch Set
...
$ git commit --amend
$ git push gerrit HEAD:refs/for/master
...
```

> You can find more details at [Gerrit Code Review Workflow](https://review.gerrithub.io/Documentation/intro-user.html#code-review).
