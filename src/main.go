package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"src/lru"
)

type kv struct {
	BlockId string `json:"块ID"`
	DestIp  String `json:"目的IP"`
}

type String string

func (d String) Len() int {
	return 1
}
func main() {

	hit, unhit := 0, 0
	srcfile := flag.String("srcfile", "/home/gavin/桌面/test/srcfileKV.json", "K-V文件")
	maxCache := flag.Int("maxCache", 8, "最大缓存容量:每8代表1Gb，因为blk大小为128M，此处用了1")
	flag.Parse()

	fd, open := os.Open(*srcfile)
	if open != nil {
		fmt.Println(open)
	}
	defer func() {
		fd.Close()
	}()
	r := bufio.NewReader(fd)
	var m []kv
	readString, _ := r.ReadString('\n')
	json.Unmarshal([]byte(readString), &m)

	cache := lru.New(int64(*maxCache), nil)
	for _, kv := range m {
		_, ok := cache.Get(kv.BlockId)
		if ok {
			fmt.Println("命中")
			hit++
		} else {
			fmt.Println("未命中")
			unhit++
			cache.Add(kv.BlockId, kv.DestIp)
		}
	}
	var res float64
	res = float64(hit) / float64(hit+unhit)
	fmt.Println(res)

}
