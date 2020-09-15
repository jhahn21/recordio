# recordio

Package recordio implements a file format for a sequence of variable-length data.

[![Build Status](https://travis-ci.com/jhahn21/recordio.svg?branch=master)](https://travis-ci.com/jhahn21/recordio)
[![GoDoc](https://godoc.org/github.com/jhahn21/recordio?status.svg)](https://godoc.org/github.com/jhahn21/recordio)
[![Go Report Card](https://goreportcard.com/badge/github.com/jhahn21/recordio)](https://goreportcard.com/report/github.com/jhahn21/recordio)

## Benchmarks

On `Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz`

```
goos: darwin
goarch: amd64
pkg: github.com/jhahn21/recordio
BenchmarkRead/nocompression_100-8         	11914750	        98.7 ns/op	1013.57 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/nocompression_1000-8        	10257574	       112 ns/op	8930.78 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/nocompression_10000-8       	 4645444	       257 ns/op	38851.13 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/snappy_100-8                	 4525521	       265 ns/op	 377.92 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/snappy_1000-8               	 3387344	       352 ns/op	2841.98 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/snappy_10000-8              	 1000000	      1067 ns/op	9373.59 MB/s	      32 B/op	       1 allocs/op
BenchmarkRead/zlib_100-8                  	  893626	      1204 ns/op	  83.07 MB/s	    4228 B/op	       4 allocs/op
BenchmarkRead/zlib_1000-8                 	  617709	      1649 ns/op	 606.25 MB/s	    4228 B/op	       4 allocs/op
BenchmarkRead/zlib_10000-8                	  177680	      5841 ns/op	1712.05 MB/s	    4228 B/op	       4 allocs/op
BenchmarkWrite/nocompression_100-8        	71513655	        16.2 ns/op	6184.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/nocompression_1000-8       	68341375	        16.2 ns/op	61711.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/nocompression_10000-8      	69492022	        16.2 ns/op	617672.71 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/snappy_100-8               	 6543421	       183 ns/op	 545.83 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/snappy_1000-8              	 2818880	       424 ns/op	2358.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/snappy_10000-8             	  498268	      2255 ns/op	4433.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrite/zlib_100-8                 	   31576	     37292 ns/op	   2.68 MB/s	      25 B/op	       0 allocs/op
BenchmarkWrite/zlib_1000-8                	   13008	     92430 ns/op	  10.82 MB/s	      62 B/op	       0 allocs/op
BenchmarkWrite/zlib_10000-8               	    5731	    196485 ns/op	  50.89 MB/s	     143 B/op	       0 allocs/op
PASS
ok  	github.com/jhahn21/recordio	23.971s
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
