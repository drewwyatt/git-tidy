# 🗑 git-tidy

[![crates.io](https://img.shields.io/crates/v/git-tidy?style=flat-square)](https://crates.io/crates/git-tidy)

Tidy up stale git branches.

[![asciicast](https://asciinema.org/a/389715.svg)](https://asciinema.org/a/389715)

## Installation

### Homebrew

```bash
$ brew tap drewwyatt/tap
$ brew install git-tidy
```

### Cargo

```bash
$ cargo install git-tidy
```

#### ⚠️ You may need to update `cargo` for this ⚠️

If you are seeing an error like the one in [this issue](https://github.com/drewwyatt/git-tidy/issues/45):

```
▪ cargo install git-tidy
    Updating crates.io index
  Installing git-tidy v2.0.1
error: failed to compile `git-tidy v2.0.1`, intermediate artifacts can be found at `/tmp/cargo-installgtcftB`

Caused by:
  failed to select a version for the requirement `zeroize = "^0.9.3"`
  candidate versions found which didn't match: 1.3.0, 1.2.0, 1.1.1, ...
  location searched: crates.io index
  required by package `dialoguer v0.7.1`
      ... which is depended on by `git-tidy v2.0.1`
```

You can probably fix this by updating cargo with:

```sh
rustup update
```


### Previous versions

Newer versions of `git-tidy` are (for now) only available from Homebrew and [crates.io](https://crates.io/crates/git-tidy), but you can still get `1.0.0` from the following places:

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
