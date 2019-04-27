package main

import "fmt"
import "sort"
import "strings"

const (
	x = iota
	y = 5
	z = iota
)

func test() (int, bool, string) {
	a, b, c := 1, true, "chenlei"
	return a, b, c
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func maxInNum(num []int, size int) int {
	var maxnum int

	for _, x := range num {
		maxnum = max(x, maxnum)
	}

	return maxnum
}

func arraylen(a []int) int {
	return len(a)
}

func twoSum(nums []int, target int) []int {
	ra := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ra[0] = i
				ra[1] = j
				return ra
			}
		}
	}

	return nil
}

func arraycopy(array []int) []int {

	for i, x := range array {
		array[i] = x
	}

	return array
}

func main() {
	name := "chenlei"

	fmt.Print("Hello!", name)

	a, b, c := 1, false, "123"
	_, b, c = test()

	fmt.Print(a, b, c)

	fmt.Print(x, y, z)

	var p *int
	var qq = 100

	p = &qq

	fmt.Printf("qq = %d, p = %d", qq, *p)

	var number = [10]int{1, 2, 3, 4, 5, 60, 7, 8, 9, 0}
	var i, x int

	for i, x = range number {
		fmt.Printf("nmber[%d]=%d\n", i, x)
	}

	var n1, n2 int = 10, 20
	var num = []int{1, 2, 3, 4, 5, 60, 7, 8, 90, 10}

	fmt.Printf("max is %d\n", max(n1, n2))
	fmt.Printf("max is %d\n", maxInNum(num, 10))

	aa := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len = %d, cap = %d\n", len(aa), cap(aa))
	for _, x := range aa {
		fmt.Printf("%d", x)
	}
	fmt.Println()

	s1 := aa[3:6]
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	for _, x := range s1 {
		fmt.Printf("%d", x)
	}
	fmt.Println()

	s2 := make([]int, 10, 20)
	s2[4] = 5
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	for _, x := range s2 {
		fmt.Printf("%d", x)
	}
	fmt.Println()

	var ss = []int{1, 3, 2, 4, 8, 5, 6}
	var sp = make([]int, 10, 20)
	sort.Ints(ss)
	for _, x := range ss {
		fmt.Printf("%d ", x)
	}
	fmt.Println()

	fmt.Printf("arraylen = %d\n", arraylen(ss))

	fmt.Printf("len = %d, cap = %d\n", len(sp), cap(sp))

	sp = arraycopy(ss)
	for _, x := range sp {
		fmt.Printf("%d ", x)
	}
	fmt.Println()

	var man = Person{"ChenLei", 30}
	//var pman = &man

	man.SayHello()
	man.ChangeName("aa")
	man.ChangeAge(50)
	man.SayHello()

	fmt.Println(lengthOfLongestSubstring4("abcded1234"))

	//LeetCode 13
	romannum := romanToInt("IX")
	fmt.Println(romannum)

	//LeetCode 14
	var aStr = []string{"abc", "abcd", "abcde"}
	fmt.Println(longestCommonPrefix2(aStr))

	fmt.Println("----- LeetCode 004 -----")
	var nums1 = []int{1, 2}
	var nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, max1, max2 int
	var mid float64
	var len1 = len(nums1)
	var len2 = len(nums2)
	var sumlen = len1 + len2
	var mixlen int
	var mix = []int{}

	if sumlen == 0 {
		return 0
	}

	for {

		if len(mix) == sumlen/2+1 {
			break
		}

		if i < len1 {
			max1 = nums1[i]
		} else {
			max1 = 0x7fffffff
		}

		if j < len2 {
			max2 = nums2[j]
		} else {
			max2 = 0x7fffffff
		}

		if max1 <= max2 {
			mix = append(mix, max1)
			i++
		} else {
			mix = append(mix, max2)
			j++
		}
	}

	fmt.Println(mix)

	mixlen = len(mix)

	if mixlen%2 == 1 {
		mid = (float64(mix[mixlen-1]) + float64(mix[mixlen-2])) / 2
	} else {
		mid = float64(mix[mixlen-1])
	}

	return mid
}

func longestCommonPrefix2(strs []string) string {
	var ret, prefix string
	var exit bool
	var length int

	var min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(strs) == 0 {
		return ""
	}

	length = len(strs[0])

	for i := 0; i <= length; i++ {
		prefix = strs[0][0:i]

		for j := 0; j < len(strs); j++ {
			length = min(length, len(strs[j]))

			if false == strings.HasPrefix(strs[j], prefix) {
				exit = true
			}
		}

		if exit != true {
			ret = prefix
		}
	}

	return ret
}

func longestCommonPrefix(strs []string) string {
	var ret string
	var exit bool

	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]

		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= i || c != strs[j][i] {
				exit = true
				break
			}
		}

		if exit == false {
			ret += string(c)
		} else {
			break
		}
	}

	return ret
}

func romanToInt(s string) int {
	var ret, num, last int
	var convert = make(map[byte]int)

	convert['I'] = 1
	convert['V'] = 5
	convert['X'] = 10
	convert['L'] = 50
	convert['C'] = 100
	convert['D'] = 500
	convert['M'] = 1000

	for i := 0; i < len(s); i++ {
		num = convert[s[i]]

		if num > last {
			ret += num - last - last
		} else {
			ret += num
		}

		last = num
	}

	return ret
}

func lengthOfLongestSubstring(s string) int {
	var lenMax int

	for i := 0; i < len(s); i++ {
		dict := make(map[byte]int)
		dict[s[i]] = i
		length := 1

		for j := i + 1; j < len(s); j++ {
			v, ok := dict[s[j]]

			if !ok {
				dict[s[j]] = v
				length++
			} else {
				break
			}
		}

		if length > lenMax {
			lenMax = length
		}
	}

	return lenMax
}

func lengthOfLongestSubstring2(s string) int {
	var i, j, max int
	var dict = make(map[byte]int)

	for i < len(s) && j < len(s) {
		x, ok := dict[s[j]]

		if ok {
			delete(dict, s[i])
			i++

		} else {
			dict[s[j]] = x
			j++

			if max < j-i {
				max = j - i
			}
		}

	}

	return max
}
func lengthOfLongestSubstring3(s string) int {
	var i, j, m, max int
	var dict = make(map[byte]int)
	var length = len(s)

	for j = 0; j < length; j++ {
		_, ok := dict[s[j]]

		if ok {
			for i = m; i <= dict[s[j]]; i++ {
				delete(dict, s[i])
			}
			m = i
		}

		dict[s[j]] = j

		if max < j-i+1 {
			max = j - i + 1
		}
	}

	return max
}

func lengthOfLongestSubstring4(s string) int {
	var left, right, start, max int

	for right = 0; right < len(s); right++ {
		for left = start; left < right; left++ {
			if s[left] == s[right] {
				break
			}
		}

		if left == right {
			if max <= right-start {
				max = right - start + 1
			}
		} else {
			start = left + 1
		}
	}

	return max
}

type Person struct {
	name string
	age  int
}

func (p *Person) ChangeName(name string) {
	p.name = name
}

func (p Person) ChangeAge(age int) {
	p.age = age
}

func (p Person) SayHello() {
	fmt.Printf("Hello, My name is %s, I am %d years old!\n", p.name, p.age)
}
