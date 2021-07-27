package main

import "fmt"

func main() {

	//输出1~100中被6整除的数
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}

		fmt.Printf("%d \t", i)
	}

}
