# 🧽 Gitclean

Delete local git branches (in the current directory) that have a remote tracking branch that is `: gone`.

## Usage

```
$ gitclean # executes "git branch -d" on ": gone" branches
```

**With force delete:**

```
$ gitclean -f # same as above, but with "-D" instead of "-d"
```
