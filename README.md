# hug-status-web

[![Build Status](https://travis-ci.org/hugbotme/hug-status-web.svg?branch=master)](https://travis-ci.org/hugbotme/hug-status-web)

Outputs statistics about the work of hugbot in JSON format.

## Statistics

```json
{
  "in_progress": 0,
  "merged": 0,
  "closed": 0,
  "received": 3
}
```

## Installation

```
$ cp config.json.dist config.json
$ go get
$ go build
$ ./hug-status-web --config="./config.json"
$ open http://localhost:8080/info.json
```

## Configuration

```json
{
  "redis": {
    "url": ":6379",
    "auth": ""
  }
}
```

## Motivation

This project was created by [Andy Grunwald](https://github.com/andygrunwald), [Jan-Erik Rediger](https://github.com/badboy), [Christoph Reinartz](https://github.com/creinartz), [Jan van Thoor](https://github.com/janvt), [Madeleine Neumann](https://github.com/madeleine-neumann) and [Matthias Endler](https://github.com/mre) at the [Hamburg Hackathon (May 23rd & 24th 2015)](http://hamburg-hackathon.de/hackathon/).

Why? Because no one likes to correct typos :smile:

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
