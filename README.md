# go-van

![travis](https://travis-ci.org/crispgm/go-van.svg?branch=master)
![codecov](https://codecov.io/gh/crispgm/go-van/branch/master/graph/badge.svg)

A simple project files watcher and deployer, which syncs dev files to remote machines at ease.

go-van is a Go implementation of [Caravan](https://github.com/crispgm/caravan). Compared to Ruby version of Caravan:

* It highly depends on `caravan.yml`, which assumes that `caravan.yml` is already setup.
* Only `rsync` is supported.

## Installation

```shell
go get -u github.com/crispgm/go-van
```

## Quick Start

1. Init `caravan.yml`:

    ```shell
    $ go-van -init
    Creating `caravan.yml`...
    Make sure to specify `src` and `dst` to watch and deploy to right place.
    ```

2. Edit `caravan.yml`:

    ```shell
    # Open with your favorite editor, `vim` for example
    $ vim caravan.yml
    ```

    Specify `src`, `dst`, and other configuration in `master` scope:

    ```yaml
    ---
    master:
      src: .
      dst: user@target:/path/to/project
      debug: false
      deploy_mode: rsync
      incremental: true
      extra_args:
      - "--delete"
      - "--exclude=.git"
      exclude:
      - ".git"
    ```

3. Start to watch:

    ```shell
    $ go-van
    Reading configuration...
    => debug: false
    => once: false
    => source: .
    => destination: .
    => deploy_mode: rsync
    => incremental: true
    => extra_args: [--delete]
    => exclude: [.git .svn /node_modules]
    Starting to watch...
    ```

4. When a file is changed, it syncs:

    ```shell
    [20:46:05] Event 0x41217e0 /Users/david/path/to/file.py
    ```

## Usage

Generate an empty `caravan.yml`:

```shell
$ go-van -init
Creating `caravan.yml`...
Make sure to specify `src` and `dst` to watch and deploy to right place.
```

Run with default:

```shell
# Default run, with `caravan.yml` and `master` spec
$ go-van
```

And you may specify config file name and spec name:

```shell
# Special spec name
$ go-van -spec my_spec
# Specify config file name
$ go-van -conf another_caravan.yml
# And both
$ go-van -conf another_caravan.yml -spec my_spec
```

Deploy once:

```shell
$ go-van -once
Reading configuration...
=> debug: false
...
Deploying at once and for once...
```

## Configuration

### Debug

Show debug outputs.

### Deploy Mode

Only support rsync in `go-van`, compared to `caravan`.

### Exclusion

Exclusion denotes exclude path for watching, not deploying. Hence, use git/svn in source path instead of destination path or checkout [Extra Arguments](#extra-arguments)

### Extra Arguments

Extra arguments will be passed to deployer (e.g. `rsync`) as arguments.

Support sync deletion:

```yaml
master:
  src: .
  dst: /path/to/project
  debug: false
  deploy_mode: rsync
  incremental: true
  extra_args:
  - "--delete"
  exclude:
  - ".git"
  - ".svn"
```

Ignore `.git` files, use `extra_args`:

```yaml
extra_args:
- "--exclude=.git"
```

## License

MIT License.

Copyright (c) 2019 David Zhang.
