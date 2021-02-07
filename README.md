# 🗑 git-tidy

[![crates.io](https://img.shields.io/crates/v/git-tidy?style=flat-square)](https://crates.io/crates/git-tidy)

Tidy up stale git branches.

[![asciicast](https://asciinema.org/a/389715.svg)](https://asciinema.org/a/389715)

## Installation

### Cargo

```bash
$ cargo install git-tidy
```

### Previous versions

Newer versions of `git-tidy` are (for now) only available from [crates.io](https://crates.io/crates/git-tidy), but you can still get `1.0.0` from the following places:

#### Homebrew

```bash
$ brew tap drewwyatt/tap
$ brew install git-tidy
```

#### Snapcraft

```bash
$ sudo snap install git-tidy
```

#### Go

```bash
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

Present all stale (": gone") branches in a checkbox list, allowing user to opt-in to deletions.

```bash
$ git tidy -i
# or
$ git tidy --interactive
# with force
$ git tidy -if
# or
$ git tidy --interactive --force
```
