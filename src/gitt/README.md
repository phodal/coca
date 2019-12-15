
GitLogs

```
git log --all --numstat --date=short --pretty=format:'--%h--%ad--%aN' --no-renames
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

