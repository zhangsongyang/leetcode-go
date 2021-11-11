package main

import (
	"fmt"
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	//fmt.Println(cmp.Equal("he", "he"))
	//a := 100
	//fmt.Println(&a)
	//var pr *int = &a
	//fmt.Println(pr)
	//fmt.Println(*pr)

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = -34444, 2
	fmt.Println(b, c)

	println(Unknown, Female, Male)

	a11, a22 := 1, 2
	if a11 < a22 {
		fmt.Printf("@@@@@@@@ %d\n", a11+a22)
	} else {
		fmt.Println("-------")
	}

	//sum := 0
	//for i := 0; i < 100; i++ {
	//	sum += 1
	//}
	//fmt.Println(sum)

	//sum := 1
	//for ; sum <= 10; {
	//	sum += sum
	//}
	//fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	str := []string{"zhang", "meng"}
	for _, s := range str {
		fmt.Println(s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	println(max(1, 2))
	println(swap("meng", "zhang"))

	// 数组
	//var balance [10] float32

	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(balance))
	balanceTwo := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(balanceTwo))

	var balanceA = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(balanceA))

	balanceB := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(balanceB))

	//  将索引为 1 和 3 的元素初始化
	balanceC := [5]float32{1: 2.0, 3: 7.0}
	for i, s := range balanceC {
		fmt.Printf("第 %d 位 s 的值 = %f\n", i, s)
	}

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var ii, jj, k int
	// 声明数组的同时快速初始化数组
	balance100 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for ii = 0; ii < 5; ii++ {
		fmt.Printf("balance[%d] = %f\n", ii, balance100[ii])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for jj = 0; jj < 5; jj++ {
		fmt.Printf("balance2[%d] = %f\n", jj, balance2[jj])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	var aa int = 10

	fmt.Printf("变量的地址: %x\n", &aa)

}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
