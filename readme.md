# Description

Small, personal program to help me organize my media by using symlinks to create a tags.

# Usage

```bash
ztag -file testfile -type pic tag1 tag2
```

## Arguments

-file: the name of the file to create tag(s) for

-type: the type of file being tagged

All other arguments are tags to be applied to the file

### Types

Accepted types are:

- doujin
- pic
- vid
- story

### Tags

Tags are automatically converted to lowercase to help prevent duplication.

## Requirements

The environment variable `ZDIR` is required to be set. This is where the tags will be created.

# Backlog

- Refactor to better facilitate unit testing
- If the given file is a symlink, apply the tag to the actual file and not to the symlink
- Decide on some way to set accepted types without hard-coding
