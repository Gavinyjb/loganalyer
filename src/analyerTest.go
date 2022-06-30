package main

import (
	"./05-analyer"
	"fmt"
)

func main() {
	scrPsth1 := "/Users/gavin/Desktop/loganalyer-master/logs/01-src/slave2/JsonOut/extract/hadoop-root-datanode-slave2_rKV.json"
	scrPsth2 := "/Users/gavin/Desktop/loganalyer-master/logs/01-src/slave1/JsonOut/extract/hadoop-root-datanode-slave1_rKV.json"
	maxCache := "1,2,4,8,16"
	policies := "lru,lfu,fifo,mru,clock"
	analyer.Analyer(scrPsth1, maxCache, policies)
	fmt.Println("<----------------------->")
	analyer.Analyer(scrPsth2, maxCache, policies)
}
