// package analyer
package main

import (
	"../policy/clock"
	"../policy/fifo"
	"../policy/lfu"
	"../policy/lru"
	"../policy/mru"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	lru2 "github.com/hashicorp/golang-lru"
	"os"
	"strconv"
	"strings"
)

type kv struct {
	BlockId string `json:"块ID"`
	DestIp  string `json:"目的IP"`
}

func Analyer(srcfile, maxCaches, policies string) {
	fd, open := os.Open(srcfile)
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
	maxCacheList := strings.Split(maxCaches, ",")
	policylist := strings.Split(policies, ",")
	for _, v := range maxCacheList {
		maxRatio, bestPolicy := 0.0, ""
		for _, policy := range policylist {
			maxCache, _ := strconv.ParseInt(v, 10, 64)
			ret := cache_hit_ratio(m, maxCache, policy)
			if ret > maxRatio {
				maxRatio = ret
				bestPolicy = policy
			}
		}
		fmt.Println("<---------------->")
		fmt.Printf("当前容量：%v;当前最佳置换算法的:%v;命中率：%v%%\n", v, bestPolicy, maxRatio)
		fmt.Println("<---------------->")
	}
}

func cache_hit_ratio(m []kv, maxCache int64, policy string) float64 {
	//fmt.Printf("KV长度：%v\n",len(m))
	var res float64
	switch policy {
	case "lru":
		cache := lru.NewCache[string, string](lru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		//cache.Show()
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "lfu":
		cache := lfu.NewCache[string, string](lfu.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "mru":
		cache := mru.NewCache[string, string](mru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "fifo":
		cache := fifo.NewCache[string, string](fifo.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "clock":
		cache := clock.NewCache[string, string](clock.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "arc":
		cache, _ := lru2.NewARC(int(maxCache))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Add(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		//cache.Show()
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "2Q":
		cache, _ := lru2.New2Q(int(maxCache))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Add(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		//cache.Show()
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	default:
		//默认LRU
		cache := lru.NewCache[string, string](lru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	}
	return res
}

func main() {
	maxCaches := flag.String("maxCaches", "2,4,8,16,32", "最大缓存容量")
	//2Q替换算法 容量不可为1
	policies := flag.String("policies", "lru,lfu,fifo,mru,clock,arc,2Q", "页面置换算法")

	kvFilePath := flag.String("input", "./slave1KV.json", "K-V文件")

	flag.Parse()
	Analyer(*kvFilePath, *maxCaches, *policies)
}
