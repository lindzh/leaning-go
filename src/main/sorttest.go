package main

import (
	"fmt"
	"sort"
)

/*

需要实现以下接口才能调用sort
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
*/
func main() {

	data := make(map[string]int)
	data["cc"] = 21
	data["vc"] = 34
	data["test"] = 54
	data["aa"] = 66
	data["43"] = 55

	keys := make([]string, 0)

	for k, _ := range data {
		keys = append(keys, k)
	}

	//接口转换

	skeys := sort.StringSlice(keys)
	fmt.Println(keys)

	sort.Sort(skeys)

	fmt.Println(skeys)

	for _, v := range skeys {
		fmt.Println("key:", v, " value:", data[v])
	}
}
