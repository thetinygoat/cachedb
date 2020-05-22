![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/wjr7dpv2xro705z7xz88.png)

  

[![Build Status](https://travis-ci.com/thetinygoat/cachedb.svg?token=QMSyQuzztbU3qV9Nxgsf&branch=master)](https://travis-ci.com/thetinygoat/cachedb)

## What is Cachedb
Cachedb is a lightweight and fast, in-memory data store, which can be used as database cache.
## Installation
Make sure `$GOPATH/bin` is added to your path, then run the following commands.
-  `$ go get -u github.com/thetinygoat/cachedb`
- `$ cachedb`

The server will be started on `port 9898`.

## Benchmarks
![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/8staw4cudsyeepg8cj55.png)

This benchmark was done using apache benchmark, with 10000 requests and 1000 concurrent requests.
## Todo
Cachedb is currently under development, if you want to contribute feel free to open an issue or work on the following problems, if you are working on some issue please reach out, i would love to help.

- [ ] Work on client libraries
- [ ] Work on multilevel caching
- [ ] Documentation
- [x] Benchmarks
- [x] Work on router 

If you would like to help with is project, feel free to email me at iamsainisachin@gmail.com.