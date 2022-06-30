package initState

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func InitState(srcpath, outFile string) {
	var files []string
	/*
		会遍历全部文件
		err := filepath.Walk(srcpath, func(path string, info os.FileInfo, err error) error {
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
	*/
	f, err := os.Open(srcpath)
	if err != nil {
		log.Fatal(err)
	}
	fileInfos, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			files = append(files, filepath.Join(srcpath, fileInfo.Name()))
		}
	}
	infoList := make([]wrInfo, 0)
	wrfd, createrr := os.OpenFile(outFile+"初始状态.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
	//记录节点数量
	nodes := make(map[string]int, 0)
	for _, file := range files {
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
				//fmt.Println(string(buf))
			}
			//判断有几个节点
			nodes[info.SrcIP]++
			nodes[info.DestIp]++
			infoList = append(infoList, info)

		}
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
	for s, _ := range nodes {
		blkInit := make(map[string]int, 0)
		for _, info := range infoList {
			if info.DestIp == s && strings.Contains(info.Wr, "读") && blkInit[info.BlockId] == 0 {
				//  info.Wr == "\\u003c----读" &&
				blkInit[info.BlockId] = 1
			}
			if info.DestIp == s && strings.Contains(info.Wr, "写") && blkInit[info.BlockId] == 0 {
				// info.Wr == "----\\u003e写" &&
				blkInit[info.BlockId] = -1
			}
			//if info.DestIp == s && strings.Contains(info.Wr, "读") && blkInit[info.BlockId] == 1 {
			//	// info.Wr == "\\u003c----读" &&
			//	blkInit[info.BlockId] = 1
			//}
			//if info.DestIp == s && strings.Contains(info.Wr, "写") && blkInit[info.BlockId] == 1 {
			//	//  info.Wr == "----\\u003e写" &&
			//	//fmt.Println("之前不是有了吗？写入覆盖？")
			//}
			//if info.DestIp == s && strings.Contains(info.Wr, "读") && blkInit[info.BlockId] == -1 {
			//	//  info.Wr == "\\u003c----读" &&
			//	blkInit[info.BlockId] = -1
			//}
			//if info.DestIp == s && strings.Contains(info.Wr, "写") && blkInit[info.BlockId] == -1 {
			//	//  info.Wr == "----\\u003e写" &&
			//	//fmt.Println("再次写入覆盖？")
			//}
		}
		w.WriteString(s)
		w.WriteByte('\n')
		for s2, i := range blkInit {
			if i > 0 {
				//fmt.Println(s2)
				w.WriteString(s2)
				w.WriteByte('\n')
			}
		}
		w.Flush()
	}
	//for _, info := range infoList {
	//
	//	js, _ := json.Marshal(info)
	//	w.WriteString(string(js))
	//	w.WriteByte('\n')
	//}
	//w.Flush()

}
