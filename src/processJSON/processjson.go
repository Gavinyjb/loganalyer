package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
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

	srcpath := flag.String("srcpath", "/home/gavin/example.json", "json文件所在目录")
	flag.Parse()
	var files []string
	err := filepath.Walk(*srcpath, func(path string, info os.FileInfo, err error) error {
		//不处理文件夹
		if info.IsDir() {
			return nil
		}
		//非json文件不处理
		if filepath.Ext(path) != ".json" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {

		infoList := make([]wrInfo, 0)
		fd, err2 := os.Open(file)
		if err2 != nil {
			fmt.Println("打开文件报错")
			fmt.Println(err2)
		}
		defer func() {
			err := fd.Close()
			if err != nil {
				fmt.Println(err)
			}
		}()
		r := bufio.NewReader(fd)
		wrfd, createrr := os.OpenFile(strings.TrimRight(file, ".json")+"已排序.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if createrr != nil {
			fmt.Println(createrr)
		}
		defer func() {
			closeerr := wrfd.Close()
			if closeerr != nil {
				fmt.Println(closeerr)
			}
		}()
		w := bufio.NewWriter(wrfd)
		for {
			buf, err3 := r.ReadBytes('\n')
			if err3 == io.EOF {
				break
			}
			if err3 != nil {
				fmt.Println(err)
			}
			var info wrInfo
			err := json.Unmarshal(buf, &info)
			if err != nil {
				fmt.Println("反序列化出错")
			}
			infoList = append(infoList, info)

		}

		sort.Slice(infoList, func(i, j int) bool {
			/*
				时间戳从大到小
			*/
			//if strings.Compare(infoList[i].Timestamp, infoList[j].Timestamp) < 0 {
			//	return false
			//} else {
			//	return true
			//}
			/*
				时间戳从小到大
			*/
			if strings.Compare(infoList[i].Timestamp, infoList[j].Timestamp) > 0 {
				return false
			} else {
				return true
			}
		})
		for _, info := range infoList {
			js, _ := json.Marshal(info)
			w.WriteString(string(js))
			w.WriteByte('\n')
		}
		w.Flush()
	}
}
