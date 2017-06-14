# go-simple-httpd

## Overview

This is simple http static server, written by golang.  
This server can map different http PATH to different directory.  

## Installation

Install and configure your [Golang](https://golang.org) environment, then:
```
[user@localhost] $ go get -u -v github.com/alienhjy/go-simple-httpd
[user@localhost] $ go install github.com/alienhjy/go-simple-httpd
```

## Usage

```
[user@localhost] $ go-simple-httpd -help
Usage of ./go-simple-httpd:
  -dir value
      Dirs for http file server. (e.g. 'list1:/dir1')
  -http string
      HTTP service address. (default ":8880")
```

