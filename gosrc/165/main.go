package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var tags1 = strings.Split(version1, ".")
	var tags2 = strings.Split(version2, ".")
	var l1, l2 = len(tags1), len(tags2)
	for i := 0; i < l1 || i < l2; i++ {
		var v1, v2 = 0, 0
		if i < l1 {
			v1, _ = strconv.Atoi(tags1[i])
		}
		if i < l2 {
			v2, _ = strconv.Atoi(tags2[i])
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
}
