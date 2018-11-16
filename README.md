# go-van

![travis](https://travis-ci.org/crispgm/go-van.svg?branch=master)
![codecov](https://codecov.io/gh/crispgm/go-van/branch/master/graph/badge.svg)

Go implementation of [Caravan](https://github.com/crispgm/caravan).

## Installation

```shell
go get -u github.com/crispgm/go-van
```

## Usage

Compared to Ruby version of Caravan, `go-van` highly depends on `caravan.yml`. It assumes that `caravan.yml` is already setup.

And only `rsync` is supported.

```shell
$ go-van
Reading configuration...
=> debug: false
=> once: false
=> src: /path/to/src
=> dst: david@remote:/path/to/dst
=> deploy_mode: rsync
=> incremental: true
=> exclude: [.git .svn]
```

And you may specify config file name and spec name.

```shell
# Specify config file name
$ go-van -conf another_caravan.yml
# And spec name
$ go-van -conf another_caravan.yml -spec my_spec
```

Generate an empty `caravan.yml`:

```shell
$ go-van -init
Creating `caravan.yml`...
Make sure to specify `src` and `dst` to watch and deploy to right place.
```

Deploy once:

```shell
$ go-van -once
```

## License

MIT License.

Copyright (c) 2018 David Zhang.
