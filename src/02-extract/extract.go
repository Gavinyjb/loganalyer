package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func creatOutDir(folderPath string) string {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}
func main() {
	srcpath := flag.String("input", "/Users/gavin/Documents/facetime trace/slave1/JsonOut", "json文件所在目录")
	outpath := flag.String("output", "/Users/gavin/Documents/facetime trace/slave1/JsonOut/02-extractread.json", "提取读块文件流目录")
	tag := flag.String("tag", "r", "标志位，标志提取r,w,rw,wr")
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
	fdout, _ := os.OpenFile(*outpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer func() {
		fdout.Close()
	}()
	w := bufio.NewWriter(fdout)
	var write func(tag string)
	write = func(tag string) {
		for i := range files {
			//file := files[len(files)-i-1] //文件名字逆序读取
			file := files[i]
			fd, err2 := os.Open(file)
			if err2 != nil {
				fmt.Println("打开文件报错")
				fmt.Println(err2)
			}
			defer func() {
				fd.Close()
			}()
			r := bufio.NewReader(fd)
			switch tag {
			case "r":
				for {
					buf, err3 := r.ReadString('\n')

					if strings.Contains(buf, "----读") {
						w.WriteString(buf)
					}
					if err3 == io.EOF {
						break
					}
				}
				w.Flush()
			case "w":
				for {
					buf, err3 := r.ReadString('\n')

					if strings.Contains(buf, "----\\u003e写") {
						w.WriteString(buf)
					}
					if err3 == io.EOF {
						break
					}
				}
				w.Flush()
			case "rw", "wr":
				for {
					buf, err3 := r.ReadString('\n')

					if strings.Contains(buf, "----读") || strings.Contains(buf, "----\\u003e写") {
						w.WriteString(buf)
					}
					if err3 == io.EOF {
						break
					}
				}
				w.Flush()
			}
		}
	}
	write(*tag)
}
