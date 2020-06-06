![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/wjr7dpv2xro705z7xz88.png)

# What is Cachedb

Cachedb is a lightweight and fast in-memory data store.

#### Features

- Can be used as a database cache.
- Can be used as message queues.
- Concurrent and highly performant.

# Installation

Make sure `$GOPATH` is added to your `$PATH`.
To install Cachedb run the following command.
`$ go get -u github.com/thetinygoat/cachedb`

# Setup

When you first run Cachedb it will create its config directory in `$HOME/.config/cachedb`.
To edit the defualt username and password make sure to edit `config.json`.
By default Cachedb runs on `http://localhost:9898`, it can be changed in `config.json`.

# Core API

Right now Cachedb supports 2 data types

- Lists
- Sets

## Lists API

### Append

Inserts the data to the end of the list.

##### Usage

- URL: `/lists/append/<your list name>?value=<your value>`.
- Method: `POST`.
- Params: `value`, it is the data you want to append.
- Example: `{{base_url}}/lists/append/mylist?value=mydata`

### Prepend

Inserts the data to the front of the list.

##### Usage

- URL: `/lists/prepend/<your list name>?value=<your value>`.
- Method: `POST`.
- Params: `value`, it is the data you want to append.
- Example: `{{base_url}}/lists/prepend/mylist?value=mydata`

### Remove Last

Removes and returns the data from the end of the list.

##### Usage

- URL: `/lists/removelast/<your list name>`.
- Method: `POST`.
- Params: None
- Example: `{{base_url}}/lists/removelast/mylist`

### Remove First

Removes and returns the data from the front of the list.

##### Usage

- URL: `/lists/removefirst/<your list name>`.
- Method: `POST`.
- Params: None
- Example: `{{base_url}}/lists/removefirst/mylist`

### Values

Returns the entire list separated by `%%`

##### Usage

- URL: `/lists/values/<your list name>`.
- Method: `GET`.
- Params: None
- Example: `{{base_url}}/lists/values/mylist`

## Sets API

### Set

Insert a key-value pair.

##### Usage

- URL: `/sets/set/<your key>?value=<your value>&ttl=<time to live, in seconds>`.
- Method: `POST`.
- Params: `value`, `ttl`, it is the data you want to append and ttl in seconds respectively.
- Example: `{{base_url}}/sets/set/mykey?value=myvalue&ttl=120`

### Get

Get a value corresponding to a particlar key.

##### Usage

- URL: `/sets/get/<your key>`.
- Method: `GET`.
- Params: None
- Example: `{{base_url}}/sets/get/mykey`

### Del

Delete a particular key-value pair

##### Usage

- URL: `/sets/del/<your key>`.
- Method: `POST`.
- Params: None
- Example: `{{base_url}}/sets/del/mykey`
