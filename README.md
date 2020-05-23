![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/wjr7dpv2xro705z7xz88.png)

  

[![Build Status](https://travis-ci.com/thetinygoat/cachedb.svg?token=QMSyQuzztbU3qV9Nxgsf&branch=master)](https://travis-ci.com/thetinygoat/cachedb)

## What is Cachedb
Cachedb is a lightweight and fast, in-memory data store, which can be used as database cache.

## Installation
Make sure `$GOPATH/bin` is added to your path, then run the following commands.
-  `$ go get -u github.com/thetinygoat/cachedb`
- `$ cachedb`

The server will be started on `http://localhost:9898`.

## API
All the requests are `GET` requests with paramaters passed as query params. Make sure to convert your keys and values to strings.

##### SET
The set API is simple we need to pass **three** things as query paramaters.

- `key` - this is the the identifier for our data.
- `value` - this is the data to be stored in reference to the key.
- `ttl` - this is the time to live in **seconds**, if you want your key to never expire pass ttl as `-1`.

Example:
`http://localhost:9898/set?key=mykey&value=my-data&ttl=120`

##### GET
The get API takes only one argument.

- `key` - this is the the identifier for our data.

Example:
`http://localhost:9898/get?key=mykey`

##### DEL
The del API also takes just one argument. It removes the specified key from the store.

- `key` - this is the the identifier for our data.

Example:
`http://localhost:9898/del?key=mykey`

##### FLUSH
The flush API does not take any argument. It removes all keys from the store.


Example:
`http://localhost:9898/flush`

## Response
The response is returned as `JSON`.

Examples:
- ``
{
  "data": "KEY_DOES_NOT_EXIST_ERROR",
  "error": true
}
``
- ``
{
  "data": "OK",
  "error": false
}
``
- ``
{
  "data": "this is cached data",
  "error": false
}
``

## Response Codes
At present there are **three** types of errors strings.

- `OK` - operations was successful.
- `KEY_DOES_NOT_EXIST_ERROR` - the specifed key does not exists.
- `KEY_EXPIRED_ERROR` - the key is expired and cannot be automatically deleted from the store.

**NOTE:** Keys are only checked for expiration when they are queried, if you are developing make sure to `flush` the store or it might lead to undesireable results.


## Benchmarks
![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/8staw4cudsyeepg8cj55.png)

This benchmark was done using apache benchmark, with 10000 requests and 1000 concurrent requests.
## Todo
Cachedb is currently under development, if you want to contribute feel free to open an issue or work on the following problems, if you are working on some issue please reach out, i would love to help.

- [ ] Work on client libraries
- [ ] Work on multilevel caching
- [x] Documentation
- [x] Benchmarks
- [x] Work on router 

If you would like to help with is project, feel free to email me at iamsainisachin@gmail.com.