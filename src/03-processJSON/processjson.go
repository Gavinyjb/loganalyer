// package processJSON
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
type kv struct {
	BlockId string `json:"块ID"`
	DestIp  string `json:"目的IP"`
}

// MergerNodeJson 合并同一个节点的日志文件
func MergerNodeJson(srcpath string) {
	var filesNew []string
	filepath.Walk(srcpath, func(path string, info os.FileInfo, err error) error {
		//不处理文件夹
		if info.IsDir() {
			return nil
		}
		//非json文件不处理
		if filepath.Ext(path) != ".json" {
			return nil
		}
		filesNew = append(filesNew, path)
		return nil
	})

	for _, file := range filesNew {
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
		FileName := strings.Split(file, ".")
		wrfd, _ := os.OpenFile(FileName[0]+".json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
				fmt.Println(err3)
			}
			w.Write(buf)
		}
		w.Flush()
		os.Remove(file)
	}
}

// ProcessJSON 处理文件
func ProcessJSON(srcpath string) (kvFilePath string) {
	var filesNew []string
	filepath.Walk(srcpath, func(path string, info os.FileInfo, err error) error {
		//不处理文件夹
		if info.IsDir() {
			return nil
		}
		//非json文件不处理
		if filepath.Ext(path) != ".json" {
			return nil
		}
		filesNew = append(filesNew, path)
		return nil
	})
	for _, file := range filesNew {
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
		//wrfd, createrr := os.OpenFile(strings.TrimRight(file, ".json")+"已排序.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
				fmt.Println(err3)
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
		kvFilePath = strings.TrimRight(file, ".json") + "KV.json"
		mapfd, maperr := os.OpenFile(strings.TrimRight(file, ".json")+"KV.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if maperr != nil {
			fmt.Println(maperr)
		}
		defer func() {
			mapcloseerr := mapfd.Close()
			if mapcloseerr != nil {
				fmt.Println(mapcloseerr)
			}
		}()
		k_v := make([]kv, 0)
		wmap := bufio.NewWriter(mapfd)
		for _, info := range infoList {
			k_v = append(k_v, kv{info.BlockId, info.DestIp})
		}
		mapjson, err := json.Marshal(k_v)
		if err != nil {
			fmt.Println(err)
		}
		_, wrerr := wmap.WriteString(string(mapjson))
		if wrerr != nil {
			fmt.Println(wrerr)
		}
		wmap.WriteByte('\n')
		wmap.Flush()

	}
	return kvFilePath
}

func main() {
	srcpath := flag.String("input", "./jsonDir", "json文件所在目录")
	flag.Parse()

	//MergerNodeJson(*srcpath)
	ProcessJSON(*srcpath)
}
