package main

import (
	"fmt"
	"sort"
)

//map按照key顺序读取
func sortMapByKey(maps map[string]string) {
	keys := []string{}
	for k := range maps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("key:", k, "value:", maps[k])
	}
}

func sortMapByValue(maps map[string]string) {
	values := []string{}
	newMaps := make(map[string]string)
	for k, v := range maps {
		values = append(values, v)
		newMaps[v] = k
	}
	sort.Strings(values)
	for _, v := range values {
		fmt.Println("key:", newMaps[v], "value:", v)
	}
}

//map的key是无序的
func main() {
	dist := make(map[string]string)
	dist["1"] = "666"
	dist["2"] = "777"
	dist["4"] = "555"
	fmt.Println(dist)
	sortMapByKey(dist)
	sortMapByValue(dist)

}
