# 9. git tell file history

Date: 2019-12-25

## Status

2019-12-25 proposed
2020-01-01 done

## Context

Refs: https://stackoverflow.com/questions/9935379/git-show-all-of-the-various-changes-to-a-single-line-in-a-specified-file-over-t

Show Line History:

```
git log -L3,5:README.md
```

Show String History

```
git log -G "## Usage" README.md
```

More:

```
for c in $(git log -G "## Usage" --format=%H -- README.md); do
    git --no-pager grep -e "## Usage" $c -- README.md
done
```

Blame:

```
git blame README.md
```

## Decision

Decision here...

## Consequences

Consequences here...
