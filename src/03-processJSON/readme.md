> 用于将某一个目录下所有json文件按照时间戳从小到大顺序排序并在目录下生成文件。

```bash
eg:go run processjson.go -srcpath=/home/gavin/project01/Kmeans/blk_read
```

```bash
-rw-rw-r--  1 gavin gavin 229172 3月  23 00:46 slave1已排序.json
-rw-rw-r--  1 gavin gavin 229172 3月  23 00:43 slave1.json
-rw-rw-r--  1 gavin gavin 121902 3月  23 00:46 slave1KV.json
-rw-rw-r--  1 gavin gavin 211124 3月  23 00:46 slave2已排序.json
-rw-rw-r--  1 gavin gavin 211124 3月  23 00:44 slave2.json
-rw-rw-r--  1 gavin gavin 112302 3月  23 00:46 slave2KV.json
-rw-rw-r--  1 gavin gavin 235000 3月  23 00:46 slave3已排序.json
-rw-rw-r--  1 gavin gavin 235000 3月  23 00:44 slave3.json
-rw-rw-r--  1 gavin gavin 125002 3月  23 00:46 slave3KV.json
-rw-rw-r--  1 gavin gavin 238008 3月  23 00:46 slave4已排序.json
-rw-rw-r--  1 gavin gavin 238008 3月  23 00:44 slave4.json
-rw-rw-r--  1 gavin gavin 126602 3月  23 00:46 slave4KV.json
```

