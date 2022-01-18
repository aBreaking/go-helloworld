package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
539. 最小时间差
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

示例 1：
输入：timePoints = ["23:59","00:00"]
输出：1

示例 2：
输入：timePoints = ["00:00","23:59","00:00"]
输出：0
提示：

2 <= timePoints.length <= 2 * 10e4
timePoints[i] 格式为 "HH:MM"
*/
func main() {
	var timePoints []string = []string{"00:00", "23:20", "01:00", "23:58"}
	fmt.Println(findMinDifference(timePoints))

}

func findMinDifference(timePoints []string) int {
	n := 24 * 60
	l := len(timePoints)
	if len(timePoints) > 24*60 {
		return 0
	}
	sort.Strings(timePoints)

	min := 99999
	for i := 0; i < l-1; i++ {
		min = calc(timePoints[i], timePoints[i+1], min)
	}
	//再看下首尾两个的差值
	n = n - toMinutes(timePoints[l-1]) + toMinutes(timePoints[0])
	if n > min {
		return min
	}
	return n
}

func calc(s1, s2 string, min int) int {
	var d = toMinutes(s1) - toMinutes(s2)
	if d < 0 {
		d = -d
	}
	if d < min {
		return d
	}
	return min
}

func toMinutes(s string) int {
	split := strings.Split(s, ":")
	h, _ := strconv.Atoi(split[0])
	m, _ := strconv.Atoi(split[1])
	return h*60 + m
}
