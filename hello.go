package main

import "fmt"
import "sort"

const (
	x = iota
	y = 5
	z = iota
)

func test()(int, bool, string) {
	a,b,c := 1,true,"chenlei"
	return a,b,c
}

func max(num1 int, num2 int) int {
	if (num1 > num2) {
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

func arraylen(a [] int) int {
	return len(a)
}

func twoSum(nums []int, target int) []int {
	k := 0
    length := len(nums)
    
    sort.Ints(nums)
    
    
    for i := 0; i < length; i++ {
        for j := length - 1; j > 0; j-- {
            if nums[i] + nums[j] == target {
				p[k] = i
				k++
				p[k] = j
				k++
                
            }
        }
        
    }
    
    return p
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

	 a,b,c := 1,false,"123"
	_,b,c = test()

	fmt.Print(a,b,c)

	fmt.Print(x,y,z)

	var p *int
	var qq int = 100

	p = &qq

	fmt.Printf("qq = %d, p = %d", qq, *p)

	var number = [10] int {1,2,3,4,5,60,7,8,9,0}
	var i, x int

	for i , x = range number {
		fmt.Printf("nmber[%d]=%d\n", i, x)
	}

	var n1,n2 int = 10,20
	var num = [] int {1,2,3,4,5,60,7,8,90,10}

	fmt.Printf("max is %d\n", max(n1,n2))
	fmt.Printf("max is %d\n", maxInNum(num, 10))


	aa := [] int {1,2,3,4,5,6,7}
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

	s2 := make([] int, 10, 20)
	s2[4] = 5
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	for _, x := range s2 {
		fmt.Printf("%d", x)
	}
	fmt.Println()

	var ss = [] int {1,3,2,4,8,5,6}
	var sp = make([] int, 10, 20)
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
}

