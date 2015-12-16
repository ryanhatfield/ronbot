# Go Ron, the RonBot

Command line based SLackbot written in Go

## Install

```shell
$ go get -u github.com/ryanhatfield/ronbot
$ go install github.com/ryanhatfield/ronbot
```

## Usage

```shell
# set token and send quote
$ ronbot -token "xxxx-xxxxxxxxx-xxxx"

# get token from 'SLACK_TOKEN' environment variable and send quote to #general
$ ronbot -channel "#general"
```
