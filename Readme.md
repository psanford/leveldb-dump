# leveldb-dump

This is a barebones leveldb dumper. I wrote this in reaction to another leveldb dumper project that seemed to be massively over-engineered for its own good.

## Usage

```
usage: ./leveldb-dump <path/to/leveldb/file>
  -output string
        Output mode (line,line_raw,csv,json) (default "line")
```

Example:

```
./leveldb-dump -output csv /tmp/example-leveldb
key-0,value-0
key-1,value-1
key-2,value-2
key-3,value-3
key-4,value-4
key-5,value-5
key-6,value-6
key-7,value-7
key-8,value-8
key-9,value-9
```
