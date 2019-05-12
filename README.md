# 🗑 git-tidy &middot; [![GitHub](https://img.shields.io/github/license/drewwyatt/git-tidy.svg)](https://github.com/drewwyatt/git-tidy/blob/master/LICENSE) [![GitHub release](https://img.shields.io/github/release/drewwyatt/git-tidy.svg)](https://github.com/drewwyatt/git-tidy/releases) [![codecov](https://codecov.io/gh/drewwyatt/git-tidy/branch/master/graph/badge.svg)](https://codecov.io/gh/drewwyatt/git-tidy)

Delete local git branches that have a remote tracking branch that is `: gone`.

[![asciicast](https://asciinema.org/a/1HgPIGWL1U24nR998vJEvPOHN.svg)](https://asciinema.org/a/1HgPIGWL1U24nR998vJEvPOHN)

## Installation

### Homebrew

```bash
$ brew tap drewwyatt/tap
$ brew install git-tidy
```

### Snapcraft &middot; [![Snap Status](https://build.snapcraft.io/badge/drewwyatt/git-tidy.svg)](https://build.snapcraft.io/user/drewwyatt/git-tidy)

```bash
$ sudo snap install git-tidy
```

### Go

```#!bash
$ go get -u github.com/drewwyatt/git-tidy
```

## Usage

```bash
$ git tidy # executes "git branch -d" on ": gone" branches
```

### With force delete

```bash
$ git tidy -f # same as above, but with "-D" instead of "-d"
# or
$ git tidy --force
```

### Interactive

Present all ": gone" branches in a checkbox list, allowing user to opt-in to deletions.

```bash
$ git tidy -i
# or
$ git tidy --interactive
# with force
$ git tidy -if
# or
$ git tidy --interactive --force
```

### Path

By default, `tidy` will execute all commands in the **current directory** (`.`), however, you can pass a path to another git repository after any/all other flags.

```bash
$ git tidy ../some/other/repo
# with flags
$ git tidy -if ../some/other/repo
```
