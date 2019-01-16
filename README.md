# redis-learn

[![Build Status](https://travis-ci.org/iMega/redis-learn.svg?branch=master)](https://travis-ci.org/iMega/redis-learn)

### Start

```
make test
```


### Stop

```
make clean
```

### Test string

```
--- PASS: TestString (0.00s)
  --- PASS: TestString/Insert_string (0.00s)
  --- PASS: TestString/Getting_string (0.00s)
  --- PASS: TestString/Inspected_the_object_by_key (0.00s)
    string_test.go:39: Value at:0x7f634fc73588 refcount:1 encoding:embstr serializedlength:4 lru:4161867 lru_seconds_idle:0
    string_test.go:51: key_sds_len:36, key_sds_avail:0, key_zmalloc: 40, val_sds_len:3, val_sds_avail:0, val_zmalloc: 8
```
