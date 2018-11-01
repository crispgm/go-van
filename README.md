# go-van

Go implementation of [Caravan](https://github.com/crispgm/caravan).

## Installation

```shell
go get -u github.com/crispgm/go-van
```

## Usage

Compare to Ruby version of Caravan, `go-van` highly depends on `caravan.yml`. It assumes that `caravan.yml` is already setup.

```shell
$ van
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
$ van -conf another_caravan.yml
# And spec name
$ van -conf another_caravan.yml -spec my_spec
```

Generate an empty `caravan.yml`:

```shell
$ van -init
Creating `caravan.yml`...
Make sure to specify `src` and `dst` to watch and deploy to right place.
```
