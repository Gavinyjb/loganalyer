package log

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type wrInfo struct {
	Timestamp string `json:"时间戳"`
	BlockId   string `json:"块ID"`
	SrcIP     string `json:"源IP"`
	Wr        string `json:"读/写"`
	DestIp    string `json:"目的IP"`
}

func creatOutDir(folderPath string) string {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}
func LogToJSON(srcPath string) (outDir string) {
	//删除末尾可能的"/"
	if srcPath[len(srcPath)-1] == '/' {
		srcPath = srcPath[:len(srcPath)-1]
	}
	//创建输出目录
	outDir = creatOutDir(srcPath + "/JsonOut/")
	//获取文件目录文件名列表
	var files []string
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		//不处理文件夹
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file)
		// 1.打开一个文件
		// 注意: 文件不存在不会创建, 会报错
		// 注意: 通过Open打开只能读取, 不能写入
		fpSrc, err := os.Open(file)
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

		// 1.打开文件(没有会创建文件)
		fpDest, err := os.OpenFile(outDir+filepath.Base(file)+".json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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

		// 3.创建写缓冲区
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
					fmt.Println(err)
				}
				err = w.Flush()
				if err != nil {
					fmt.Println("刷新失败")
				}
			}
			if err == io.EOF {
				break
			}
		}
	}
	return outDir
}
