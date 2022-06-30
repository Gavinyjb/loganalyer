//package main
//
//import (
//	"./01-log"
//	_ "./02-extractReadWrite"
//	_ "./03-processJSON"
//	_ "./04-initState"
//	_ "./05-analyer"
//	"flag"
//	"fmt"
//	_ "path/filepath"
//)
//
//func main() {
//
//	//--------01
//	srcpath01 := flag.String("srcpath", "/Users/gavin/Documents/facetime trace/slave1", "log日志路径")
//	//tag := flag.String("tag", "r", "标志位，标志提取r,w,r&w")
//	/*
//			//--------02
//			srcpath02 := flag.String("srcpath", "", "json文件所在目录")
//			outpath02 := flag.String("outpath", "/home/gavin/extractread.json", "提取读块文件流目录")
//			tag := flag.String("tag", "r", "标志位，标志提取r,w,r&w")
//			////--------03
//			srcpath03 := flag.String("srcpath03", "/Users/gavin/Desktop/loganalyer-master/logs/03-src", "json文件所在目录")
//			////--------04
//			srcpath04 := flag.String("srcpath", "/Users/gavin/Desktop/loganalyer-master/logs/04-src", "json文件所在目录")
//			outFile := flag.String("outFile", "/Users/gavin/Desktop/loganalyer-master/logs/04-src/04-out/", "初始状态各节点拥有的文件块")
//			////--------05
//			srcfile := flag.String("srcfile", "/Users/gavin/Desktop/loganalyer-master/logs/slave1KV.json", "K-V文件")
//
//		    maxCaches := flag.String("maxCaches", "2,4,8,16,32", "最大缓存容量")
//			//2Q替换算法 容量不可为1
//			policies := flag.String("policies", "lru,lfu,fifo,mru,clock,arc,2Q", "页面置换算法")
//	*/
//
//	flag.Parse()
//	//01
//	fmt.Println("开始抽取日志文件中的块读写信息------>")
//	srcpath02 := log.LogToJSON(*srcpath01)
//	fmt.Println("块读写信息已提取到：" + srcpath02)
//
//	/*
//		//02
//			fmt.Printf("开始提取%v文件------>", *tag)
//			srcpath03 := extractReadWrite.ExtractReadWrite(srcpath02, *tag)
//			fmt.Println("")
//			//03
//			processJSON.MergerNodeJson(srcpath03)
//			kvFilePath := processJSON.ProcessJSON(srcpath03)
//			//04
//			outFile04 := filepath.Dir(srcpath02)
//			initState.InitState(srcpath02, outFile04)
//			//05
//			analyer.Analyer(kvFilePath, *maxCaches, *policies)
//
//	*/
//
//}
