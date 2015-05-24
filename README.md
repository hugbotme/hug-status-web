# hug-status-web

[![Build Status](https://travis-ci.org/hugbotme/hug-status-web.svg?branch=master)](https://travis-ci.org/hugbotme/hug-status-web)

Outputs statistics about the work of hugbot in JSON format.
This is awsome. You know? Heavy reqirements.

## Statistics

```json
{
  "in_progress": 0,
  "merged": 0,
  "closed": 0,
  "queue": 3
}
```

* **in_progress**: Current amount of open pull requests.
* **merged**: Amount of merge requests which were merged.
* **closed**: Amount of merge requests which were closed and not merged.
* **queue**: Current amount of job entries in the queue.

A *job entry* is a message which contains a repository to correct.

## Reguirements

* [Redis](http://redis.io/)

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
  "status-web": {
    "url": "127.0.0.1:8080"
  },
  "redis": {
    "url": ":6379",
    "auth": ""
  }
}
```

* **status-web.url**: IP and Port (in IP:Port format) where the status will be available and callable.
* **redis.url**: IP and Port (in IP:Port format) where Redis is running.
* **redis.auth**: Password to authenticate to the Redis server. If this is an empty string, no [AUTH command](http://redis.io/commands/auth) will be executed.

## Motivation

This project was created by [Andy Grunwald](https://github.com/andygrunwald), [Jan-Erik Rediger](https://github.com/badboy), [Christoph Reinartz](https://github.com/creinartz), [Jan van Thoor](https://github.com/janvt), [Madeleine Neumann](https://github.com/madeleine-neumann) and [Matthias Endler](https://github.com/mre) at the [Hamburg Hackathon (May 23rd & 24th 2015)](http://hamburg-hackathon.de/hackathon/).

Why? Because no one likes to correct typos :smile:

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
