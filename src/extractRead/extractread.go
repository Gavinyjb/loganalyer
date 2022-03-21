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

func main() {
	srcpath := flag.String("srcpath", "/home/gavin", "json文件所在目录")
	outpath := flag.String("outpath", "/home/gavin/extractread.json", "提取读块文件流目录")
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
	}
}
