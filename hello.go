package main

import "fmt"
import "sort"

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
