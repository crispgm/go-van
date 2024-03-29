# Caravan (go-van)

[![ci](https://github.com/crispgm/go-van/actions/workflows/ci.yml/badge.svg)](https://github.com/crispgm/go-van/actions/workflows/ci.yml)
![codecov](https://codecov.io/gh/crispgm/go-van/branch/master/graph/badge.svg)

A simple project files watcher and deployer, which syncs dev files to remote machines at ease.

go-van is the Go implementation of [Caravan](https://github.com/crispgm/caravan).

![Caravan in Civ5](/assets/civ-5-caravan.png)

This is the caravan in [Sid Meier's Civilization V](http://www.civilization5.com/), where the project name originally comes from.

## Migrate from Ruby Caravan

Compared to Ruby version of Caravan:

- It highly depends on `caravan.yml`, which assumes that `caravan.yml` is already setup.
- Only `rsync` is supported.
- More powerful features (e.g. `extra_args` and `log_format`).

## Installation

```shell
go get -u github.com/crispgm/go-van
```

## Quick Start

1. Init `caravan.yml`:

   ```shell
   $ go-van -init
   Created caravan.yml in /path/to/project
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
   [20:46:05] EVENT 0x41217e0 /Users/david/path/to/file.py
   ```

## Usage

Generate an empty `caravan.yml`:

```shell
$ go-van -init
Created caravan.yml in /path/to/project
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

Exclusion denotes exclude path for watching, not deploying. Hence, use git/svn in source path instead of destination path or checkout [Extra Arguments](#extra-arguments).

### Extra Arguments

Extra arguments will be passed to deployer (e.g. `rsync`) as arguments.

#### Support Sync with Deletion

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

#### Ignore `.git` Files

```yaml
extra_args:
  - "--exclude=.git"
```

#### Log Format

Format:

- `%t`: Time string, e.g. 16:25:01
- `%T`: Timestamp
- `%e`: Event type
- `%p`: Path
- `%f`: File name

```yaml
log_format: "[%t] EVENT <%e> %p"
```

## Event Hooks (experimental)

Event hooks are designed to handle events of hooks, by which users may inject their scripts.

There are four hooks exposed, which are `OnInit` `OnChange` `OnDeploy` `OnError`.

e.g., You may add `OnInit` in `caravan.yml`:

```yaml
on_init:
  - echo "go-van is initializing"
on_change:
  - make
```

## License

MIT License.

Copyright (c) 2020 David Zhang.
