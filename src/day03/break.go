package main

import "fmt"

func main() {

	//求 1 + 2 + 3 + ... + 100
	var ans = 0

	/*
		1.switch分支中,每个case分支后用break结束当前分支,但是在Golang中break可以省略不写
		2.break可以结束正在执行的循环
	*/
	for i := 1; i <= 100; i++ {
		ans += i

		if ans > 1000 {
			break //在此次循环打断
		}
	}

	fmt.Println(ans)
}
