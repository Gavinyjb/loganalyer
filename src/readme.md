```shell
root@gavin:/home/gavin/trace/tools# ./01-log -h
Usage of ./01-log:
  -input string
    	log日志路径 (default "/Users/gavin/Documents/facetime trace/slave1")
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave1/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/hadoop-root-datanode-slave1.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave1/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave2/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.11
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/hadoop-root-datanode-slave2.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave2/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave3/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/hadoop-root-datanode-slave3.log.5
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave3/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave4/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.11
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/hadoop-root-datanode-slave4.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave4/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave5/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/hadoop-root-datanode-slave5.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave5/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave6/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/hadoop-root-datanode-slave6.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave6/JsonOut/
root@gavin:/home/gavin/trace/tools# ./01-log -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave8/
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.1
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.10
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.11
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.2
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.3
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.4
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.5
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.6
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.7
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.8
/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/hadoop-root-datanode-slave8.log.9
块读写信息已提取到：/mnt/hgfs/FB_2009_0_DNlog(DEBUG)/slave8/JsonOut/
root@gavin:/home/gavin/trace/tools#
```

```shell
root@gavin:/home/gavin/trace/tools# ./02-extract -h
Usage of ./02-extract:
  -input string
    	json文件所在目录 (default "/Users/gavin/Documents/facetime trace/slave1/JsonOut")
  -output string
    	提取读块文件流目录 (default "/Users/gavin/Documents/facetime trace/slave1/JsonOut/02-extractread.json")
  -tag string
    	标志位，标志提取r,w,rw,wr (default "r")
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave1/JsonOut/
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave1r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave2/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave2r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave3/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave3r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave4/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave4r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave5/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave5r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave6/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave6r.json
root@gavin:/home/gavin/trace/tools# ./02-extract -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/slave8/JsonOut/ -output /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave8r.json
```

```shell
root@gavin:/home/gavin/trace/tools# ./03-processJSON -h
Usage of ./03-processJSON:
  -input string
    	json文件所在目录 (default "./jsonDir")
root@gavin:/home/gavin/trace/tools# ./03-processJSON -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/
```

```shell
root@gavin:/home/gavin/trace/tools# ./05-analyer -input /mnt/hgfs/FB_2009_0_DNlog\(DEBUG\)/jsonDir/slave4rKV.json
当前容量：2;当前置换算法的:lru;命中率：25.34717715768981%
当前容量：2;当前置换算法的:lfu;命中率：21.11615833874108%
当前容量：2;当前置换算法的:fifo;命中率：25.321219987021415%
当前容量：2;当前置换算法的:mru;命中率：25.34717715768981%
当前容量：2;当前置换算法的:clock;命中率：25.321219987021415%
当前容量：2;当前置换算法的:arc;命中率：24.698247890979882%
当前容量：2;当前置换算法的:2Q;命中率：23.6469824789098%
<---------------->
当前容量：2;当前最佳置换算法的:lru;命中率：25.34717715768981%
<---------------->
当前容量：4;当前置换算法的:lru;命中率：30.28552887735237%
当前容量：4;当前置换算法的:lfu;命中率：21.310837118754055%
当前容量：4;当前置换算法的:fifo;命中率：30.402336145360152%
当前容量：4;当前置换算法的:mru;命中率：30.28552887735237%
当前容量：4;当前置换算法的:clock;命中率：30.38286826735886%
当前容量：4;当前置换算法的:arc;命中率：29.1174561972745%
当前容量：4;当前置换算法的:2Q;命中率：24.600908500973397%
<---------------->
当前容量：4;当前最佳置换算法的:fifo;命中率：30.402336145360152%
<---------------->
当前容量：8;当前置换算法的:lru;命中率：32.31018818948734%
当前容量：8;当前置换算法的:lfu;命中率：21.447112264763142%
当前容量：8;当前置换算法的:fifo;命中率：32.323166774821544%
当前容量：8;当前置换算法的:mru;命中率：32.31018818948734%
当前容量：8;当前置换算法的:clock;命中率：32.32965606748864%
当前容量：8;当前置换算法的:arc;命中率：32.19987021414666%
当前容量：8;当前置换算法的:2Q;命中率：27.61842959117456%
<---------------->
当前容量：8;当前最佳置换算法的:clock;命中率：32.32965606748864%
<---------------->
当前容量：16;当前置换算法的:lru;命中率：32.58922777417261%
当前容量：16;当前置换算法的:lfu;命中率：21.70019467878001%
当前容量：16;当前置换算法的:fifo;命中率：32.537313432835816%
当前容量：16;当前置换算法的:mru;命中率：32.58922777417261%
当前容量：16;当前置换算法的:clock;命中率：32.57624918883842%
当前容量：16;当前置换算法的:arc;命中率：32.55678131083712%
当前容量：16;当前置换算法的:2Q;命中率：31.57040882543803%
<---------------->
当前容量：16;当前最佳置换算法的:lru;命中率：32.58922777417261%
<---------------->
当前容量：32;当前置换算法的:lru;命中率：32.66060999351071%
当前容量：32;当前置换算法的:lfu;命中率：22.037637897469175%
当前容量：32;当前置换算法的:fifo;命中率：32.64114211550941%
当前容量：32;当前置换算法的:mru;命中率：32.66060999351071%
当前容量：32;当前置换算法的:clock;命中率：32.64763140817651%
当前容量：32;当前置换算法的:arc;命中率：32.59571706683971%
当前容量：32;当前置换算法的:2Q;命中率：32.7384815055159%
<---------------->
当前容量：32;当前最佳置换算法的:2Q;命中率：32.7384815055159%
<---------------->
```

