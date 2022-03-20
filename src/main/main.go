package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type wrInfo struct {
	Timestamp string `json:"时间戳"`
	BlockId   string `json:"块ID"`
	SrcIP     string `json:"源IP"`
	Wr        string `json:"读/写"`
	DestIp    string `json:"目的IP"`
}

func main() {

	path := flag.String("path", "/home/gavin/", "log日志路径")
	outPath := flag.String("outPath", "/home/gavin/out.txt", "输出文件路径") //eg:/home/gavin/桌面/Kmeans/slave2/res.txt
	//解析命令行参数
	flag.Parse()
	//resolute := make([]wrInfo, 0)
	//nodeList := make(map[string]bool)
	// 1.打开一个文件
	// 注意: 文件不存在不会创建, 会报错
	// 注意: 通过Open打开只能读取, 不能写入
	fpSrc, err := os.Open(*path)
	if err != nil {
		fmt.Println(err)
	}

	// 2.关闭一个文件
	defer func() {
		err = fpSrc.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 1.打开文件
	fpDest, err := os.OpenFile(*outPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// 2.关闭打开的文件
	defer func() {
		err := fpDest.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 3.创建缓冲区
	w := bufio.NewWriter(fpDest)

	// 4.读取文件中所有内容, 直到文件末尾为止
	r := bufio.NewReader(fpSrc)
	for {
		//buf, err := r.ReadBytes('\n')
		buf, err := r.ReadString('\n')
		if strings.Contains(buf, "HDFS_READ") {
			res := strings.Split(buf, " ")
			timestamp := " " + res[0] + " " + res[1]
			blockId := res[len(res)-3]
			srcIp := strings.Split(strings.Replace(res[5], "/", "", -1), ":")[0]
			destIp := strings.Split(strings.Replace(res[7], "/", "", -1), ":")[0]
			temp := wrInfo{
				Timestamp: timestamp,
				BlockId:   blockId,
				SrcIP:     srcIp,
				DestIp:    destIp,
				Wr:        "<----读",
			}
			//resolute = append(resolute, temp)
			info, err2 := json.Marshal(temp)
			if err2 != nil {
				fmt.Println(err2)
			}
			_, err := w.Write(info)
			_, err = w.WriteString("\n")
			if err != nil {
				return
			}
			if err != nil {
				return
			}
			if err != nil {
				fmt.Println(err)
			}
			err = w.Flush()
			if err != nil {
				fmt.Println("刷新失败")
			}
			//nodeList[srcIp] = true
			//nodeList[destIp] = true
			//fmt.Println(timestamp, blockId, srcIp,"<----",destIp)
			//fmt.Println(res)
		}
		if strings.Contains(buf, "HDFS_WRITE") {
			res := strings.Split(buf, " ")
			timestamp := " " + res[0] + " " + res[1]
			blockId := res[len(res)-3]
			srcIp := strings.Split(strings.Replace(res[5], "/", "", -1), ":")[0]
			destIp := strings.Split(strings.Replace(res[7], "/", "", -1), ":")[0]
			temp := wrInfo{
				Timestamp: timestamp,
				BlockId:   blockId,
				SrcIP:     srcIp,
				DestIp:    destIp,
				Wr:        "---->写",
			}
			//resolute = append(resolute, temp)
			info, err2 := json.Marshal(temp)
			if err2 != nil {
				fmt.Println(err2)
			}
			_, err := w.Write(info)
			_, err = w.WriteString("\n")
			if err != nil {
				return
			}
			if err != nil {
				fmt.Println(err)
			}
			err = w.Flush()
			if err != nil {
				fmt.Println("刷新失败")
			}
			//nodeList[srcIp] = true
			//nodeList[destIp] = true
			//fmt.Println(timestamp, blockId, srcIp,"---->",destIp)
			//fmt.Println(res)
		}
		if err == io.EOF {
			break
		}
	}

	//for s, _ := range nodeList {
	//	nodeInfo := make([]wrInfo, 0)
	//	for _, info := range resolute {
	//		if info.srcIP == s || info.destIp == s {
	//			nodeInfo = append(nodeInfo, info)
	//		}
	//	}
	//	w.WriteString(s + "------该节点的读写操作" + `\n`)
	//	for _, info := range nodeInfo {
	//		_, err2 := w.WriteString(info.timestamp + info.blockId + info.srcIP + info.wr + info.destIp + `\n`)
	//		if err2 != nil {
	//			fmt.Println(err2)
	//		}
	//	}
	//	w.Flush()
	//}
}
