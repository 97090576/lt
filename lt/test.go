package main

import "fmt"

func test1(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		tmpSum := 0
		for j := i; j < len(nums); j++ {
			tmpSum += nums[j]
			if tmpSum == k {
				cnt++
			}
		}
	}
	return cnt
}

// 1,4,-4,2,3,-3,-2,1   0,1,5,1,3,6,3,1,2    2
func test2(nums []int, k int) int {
	// m 保存前缀和对应的前缀数组的尾端元素在原数组中的次序
	m := make(map[int][]int)
	tmp := 0
	for i, v := range nums {
		tmp += v
		m[tmp] = append(m[tmp], i)
	}
	tmp = 0
	cnt := 0
	for i, v := range nums {
		tmp += v
		if tmp == k {
			cnt++
		}
		if idxs, ok := m[tmp+k]; ok {
			for _, idx := range idxs {
				if idx > i {
					cnt++
				}
			}
		}
	}
	return cnt
}

func test3(nums []int, k int) int {
	preSum := 0
	m := make(map[int]int)
	for _, v := range nums {
		preSum += v
		m[preSum]++
	}
	res := 0
	preSum = 0
	for _, v := range nums {
		preSum += v
		if preSum == k {
			res++
		}
		if cnt, ok := m[preSum+k]; ok {
			res += cnt
		}
		m[preSum]--
	}
	return res
}

// hmDistance 求 x,y 的汉明距离
func hmDistance(x, y int) int {
	xor := x ^ y
	res := 0
	for xor != 0 {
		fmt.Println(xor)
		res++
		xor &= xor - 1 // 0110 & 0101 = 0100
	}
	return res
}
