package main

import (
	"fmt"
	"math"
)

func main() {
	i := MaxMatrixArea([]int{2, 1, 5, 6, 2, 3})
	fmt.Println(i)
	i2 := Max1MatrixArea([][]int{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}})
	fmt.Println(i2)
}

// CanSplit 判断数组是否能分成两个和相等的子数组
func CanSplit(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := num; j <= target; j++ {
			dp[j] = dp[j] || dp[j-num] // 一但 dp[j] 为 true，就不会再变成 false
		}
	}
	return dp[target]
}

// FindString 找到字符串 p 中所有字符串 c 的字母异位词的起始索引
func FindString(p, c string) []int {
	if len(c) > len(p) {
		return nil
	}
	sum := 0
	for _, ch := range c {
		sum += int((ch - 'a' + 1))
	}
	l, r := 0, len(c)-1
	var res []int
	tmp := 0
	for i := l; i <= r; i++ {
		tmp += int((p[i] - 'a' + 1))
	}
	for {
		if tmp == sum {
			res = append(res, l)
		}
		r++
		if r >= len(p) {
			break
		}
		tmp += int((p[r] - 'a' + 1))
		tmp -= int((p[l] - 'a' + 1))
		l++
	}
	return res
}

// 字符串解码
func DecodeString(s string) string {
	ss := IntStack{} // 保存出现的左括号索引
	is := IntStack{} // 保存出现的数字
	// f 将 byte 转换为 string 或 int，bool = true 表示 int, false 表示 string
	f := func(r byte) (string, int, bool) {
		if r >= '0' && r <= '9' {
			return "", int(r - '0'), true
		}
		return string(r), 0, false
	}
	// repeat 将字符串 str 重复 n 次
	repeat := func(str string, n int) string {
		var res string
		for i := 0; i < n; i++ {
			res += str
		}
		return res
	}
	for i := 0; i < len(s); i++ {
		c, n, ok := f(s[i])
		if ok {
			is.Push(n)
			continue
		}
		if c == "[" {
			ss.Push(i)
		} else if c == "]" {
			// 取出最近的左括号索引
			l := ss.Pop()
			// 取出最近的数字
			n := is.Pop()
			// 取出中间的字符串
			str := s[l+1 : i]
			// 重复字符串，注意左括号前的数字也需要丢弃
			s = s[:l-1] + repeat(str, n) + s[i+1:]
			// 修正索引
			i = l + len(str)*n - 2
		}
	}
	return s
}

// 最长回文子序列
func MaxLengthSubStr(s string) {
	// dp[i][j] 表示 i->j 的最长回文子序列的长度，结果: dp[0][len(s)-1]
	// if s[i]==s[j] dp[i+1][j-1]+2
	// if s[i] != s[j] max(dp[i][j-1], dp[i+1][j])
}

// 最长回文子串
func MaxLengthSubStr2(s string) {
	// dp[i][j] 表示i->j的字符串是否是回文字符串，假如是，那么字符串长度是 j-i+1，所以可以拿到所有回文串的长度，最大值即为最长回文子串的长度
	// dp[i][j] = s[i]==s[j] && ((j-i)<=2 ||dp[i+1][j-1]) j-i <= 2 表示 i->j 之间最多只有一个字符，肯定不会影响回文串的判定，所以肯定可以剪枝
	// 初始化： dp[i][i] = true
}

// 正则表达式匹配，p 可以包含普通字符和 .*，. 能匹配任意字符，*能匹配0到多个字符，注意，*应该和其前面的字符视为一体去匹配0或多个字符
func IsMatch(s, p string) bool {
	/*
		dp[i][j] 表示 s[:i-1] 与 p[:j-1] 是否匹配，即 p 到第 j 个字符能否匹配上 s 到第 i 个字符，第 0 个字符表示空串
		当第 j 个字符不是 * 时，则第 j 个字符必须匹配第 i 个字符，并且把第 j 个字符和第 i 个字符去掉之后前面的字符也能匹配上，才能匹配上
		当第 j 个字符是 * 时，有俩个匹配逻辑，.* 匹配 0 个，或者 .* 匹配 1 个以上与第 j-1 个字符逻辑相等的字符，只要有任意一种情况能匹配上，则能匹配上
		equal(sb, pb) == true 表示 sb 和 pb 逻辑相等
		递推公式：
		dp[i][j] =
		if p[j-1] != * then dp[i-1][j-1] && equal(s[i-1], p[j-1])
		else
			if dp[i][j-2] then true
		    else for ti:=i; i>0; i-- {
				// 无法匹配了还没有匹配成功，说明已经不可能成功了
				if !equal(s[ti-1], p[j-2]) then false
				if dp[ti-1][j-2] then true
			}
		初始化：判断空串能匹配上的模式, .*.*.*.*，即第 2k 个出现的字符是 * 且去除第 2k 和 2k-1 个字符的前面的子串能匹配上空串，所以一旦出现无法匹配的情况，就可以剪枝，没必要继续检测了
		dp[0][0] = true
		for j:=2; j<=len(p); j++ {
			if j%2 == 0 {
				// 第一个非*偶数位置之后的子串都不能匹配上空串了，可以直接结束初始化
				if p[j-1] != '*' {
					break
				}
				dp[0][j] = true
			}
		}
	*/
	equal := func(sb, pb byte) bool {
		return pb == '.' || sb == pb
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 2; j <= len(p); j++ {
		if j%2 == 0 {
			if p[j] != '*' {
				break
			}
			dp[0][j] = true
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				}
				for ti := i; ti > 0; i-- {
					if !equal(s[ti-1], p[j-2]) {
						break
					}
					if dp[ti-1][j-2] {
						break
					}
				}
			} else {
				dp[i][j] = dp[i-1][j-1] && equal(s[i-1], p[j-1])
			}
		}
	}
	return dp[len(s)][len(p)]
}

// 合并升序链表
func MergeAscLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	emptyCnt := 0
	lastNode := dummyHead
	for emptyCnt < len(lists) {
		minValIdx := 0
		minVal := math.MaxInt
		// 这里是每次都遍历一次每个链表头找到最小的链表头然后将其放在合并后的链表最后面，明显是可以优化的，因为除了被选出来的链表头，其他的链表头都没有发生改变，不需要再重新比较一次大小，所以应该将其替换为最小堆或者是优先队列
		for i, node := range lists {
			if node != nil && node.Val.(int) < minVal {
				minVal = node.Val.(int)
				minValIdx = i
			}
		}
		lastNode.Next = lists[minValIdx]
		lastNode = lastNode.Next
		lists[minValIdx] = lists[minValIdx].Next
		if lists[minValIdx] == nil {
			emptyCnt++
		}
	}
	return dummyHead.Next
}

// s = "a", t = "b" 错误
func MinWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}
	f2 := func(s, t map[byte]int) bool {
		for k, v := range t {
			if s[k] < v {
				return false
			}
		}
		return true
	}
	tByte := make(map[byte]int)
	for i := range t {
		tByte[t[i]]++
	}
	m := make(map[byte]int, len(tByte))
	minStr := s
	l := 0
	r := 0
	for {
		if f2(m, tByte) {
			if r-l+1 < len(minStr) {
				minStr = s[l:r]
			}
			m[s[l]]--
			l++
			continue
		}
		// 数组类问题，一定要注意下标越界的问题，同时要区分第k个元素和第k个元素的下标
		if r >= len(s) {
			break
		}
		m[s[r]]++
		r++
	}
	return minStr
}

// 纯暴力的解法会超出时间限制，需要优化
func MaxMatrixArea(heights []int) int {
	maxArea := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(heights); i++ {
		l, r := i, i
		for l >= 0 && heights[l] >= heights[i] {
			l--
		}
		for r < len(heights) && heights[r] >= heights[i] {
			r++
		}
		maxArea = max(maxArea, heights[i]*(r-l-1))
	}
	return maxArea
}


/*
[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
*/

func Max1MatrixArea(matrix [][]int) int {
	maxArea := 0
	rows := len(matrix)
	if rows <= 0 {
		return 0
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	cols := len(matrix[0])
	// 纯暴力解法，遍历每个全1子矩阵，获取最大的面积
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != 1 {
				continue
			}
			// 以 i,j 为左上角的全1矩阵的最大面积和当前矩阵的最小长度，两者之间没有关系
			curMaxArea := 0
			curMinLength := cols
			// 遍历数组的时候不要使用三段式 for 循环，改用 while 循环，这样可以避免数组越界的问题，在 go 语言里就是使用 for 循环的时候不要使用循环条件而是使用 break 跳出循环
			k, l := i, j
			for {
				for {
					if l >= cols || matrix[k][l] != 1 {
						curMinLength = min(curMinLength, l-j)
						curMaxArea = max(curMaxArea, curMinLength*(k-i+1))
						break
					}
					l++
				}
				k++
				if k >= rows || matrix[k][j] != 1 {
					break
				}
			}
			maxArea = max(maxArea, curMaxArea)
		}
	}
	return maxArea
}
