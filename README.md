# 🗑 git-tidy

[![crates.io](https://img.shields.io/crates/v/git-tidy?style=flat-square)](https://crates.io/crates/git-tidy)

Tidy up stale git branches.

## Installation

### Cargo

```bash
$ cargo install git-tidy
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
