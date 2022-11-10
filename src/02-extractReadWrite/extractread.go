package extractReadWrite

import (
	"bufio"
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
func ExtractReadWrite(srcPath, tag string) (outFile string) {
	//创建输出目录
	outDir := creatOutDir(srcPath + "/extract/")
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
	temp := strings.Split(filepath.Base(files[len(files)-1]), ".")
	outFile = outDir + temp[0] + "_" + tag + ".json"
	fdout, _ := os.OpenFile(outFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
			case "r&w", "w&r":
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
	write(tag)
	return outFile
}
