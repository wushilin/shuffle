# shuffle
Binary can't be sent? Shuffle your file to a format that nobody knows what it is. Shuffle again to return to original

# Build
$ go install github.com/wushilin/shuffle@v1.0.0

# Shuffling
```
$ ./shuffle -s seedstr file1 file2 ...
```
`file1` will be shuffled to `file1.rev`

`file2` will be shuffled to `file2.rev`

Reversing a shuffle
```
$ ./shuffle -s seedstr -r file1.rev file2.rev ...
```

`file1.rev` will be restored as `file1.rev.res`

`file2.rev` will be restored as `file2.rev.res`

# Help
```
Usage of ./shuffle:
  -r    Do reverse convert (undo shuffle)
  -s string
        Seed for shuffling (default "ABCD1234")
```
