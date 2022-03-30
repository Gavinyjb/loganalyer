> 用于提取某一个目录下所有json文件中的读取块数据流

```Golang
eg:go run extractread.go -tag="w&r" -srcpath=/home/gavin/project01/Kmeans/slave1 -outpath=/home/gavin/project01/Kmeans/slave1.json
```
```
tag:"w"  "r"  "w&r"  "r&w"
```
```json
{"时间戳":" 2020-05-14 19:04:13,147","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744250_3461,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:13,426","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744328_3539,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:13,791","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744325_3536,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:14,418","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744255_3466,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:19,109","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744328_3539,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:19,328","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744136_3347,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:19,484","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744325_3536,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
{"时间戳":" 2020-05-14 19:04:19,492","块ID":"BP-1008346802-172.19.2.170-1588909441424:blk_1073744142_3353,","源IP":"172.19.2.172","读/写":"\u003c----读","目的IP":"172.19.2.172"}
```
