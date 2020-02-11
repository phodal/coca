
GitLogs

```
git log --all --date=short --pretty="format:[%h] %aN %ad %s" --numstat --reverse --summary
```

Delete

```
git log --pretty=%H --name-status
```

```
git log --pretty=format:"[%h] %aN %ad %s" --reverse --summary --numstat --encoding=UTF-8 --no-renames
```



Related Projects: [https://github.com/bast/gitink](https://github.com/bast/gitink)


```
$ cat example.txt

                     [feature]
                      |
                      v
               x1-----x2
              /
c1----c2----m1----c3----c4
  \        /            ^
   b1----b2----b3       |
   ^           ^       [master,HEAD]
   |           |
  [_branch]   [branch]

$ gitink --time-direction=90 --in-file=example.txt | display
```

