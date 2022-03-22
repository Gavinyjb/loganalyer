package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"src/policy/lru"
)

type kv struct {
	BlockId string `json:"块ID"`
	DestIp  string `json:"目的IP"`
}

func main() {
	srcfile := flag.String("srcfile", "/home/gavin/桌面/test/srcfileKV.json", "K-V文件")
	//maxCaches := flag.Int("maxCaches", "1,2,4,8,16", "最大缓存容量")
	maxCache := flag.Int("maxCache", 8, "最大缓存容量")
	//policies := flag.String("policies", "lru,lfu,fifo,mru,clock", "页面置换算法")
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
	//maxCacheList := strings.Split(*maxCaches, ",")
	//for _, maxCache := range maxCacheList {
	//
	//}
	cache_hit_ratio(m, *maxCache)

}

func cache_hit_ratio(m []kv, maxCache int) {
	hit, unhit := 0, 0
	cache := lru.NewCache[string, string](lru.WithCapacity(maxCache))
	for _, kv := range m {
		got, ok := cache.Get(kv.BlockId)
		if ok {
			fmt.Println("命中-----" + got)
			hit++
		} else {
			fmt.Println("未命中----" + got)
			unhit++
			cache.Set(kv.BlockId, kv.DestIp)
		}
		cache.Show()
	}
	var res float64
	res = float64(hit) / float64(hit+unhit)
	res *= 100
	fmt.Printf("命中率：%v%%\n", res)
}
