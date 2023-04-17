package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	fmt.Println("Please input your guess")
	var guess int
	// 输入我们猜的数字
	for {
		_, err := fmt.Scanf("%d", &guess)
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		if guess == secretNumber {
			break
		} else if guess > secretNumber {
			fmt.Println("bigger than secretNumber")
		} else {
			fmt.Println("smaller than secretNumber")
		}
		var ch byte
		fmt.Scanf("%c", &ch) //处理回车
	}
	fmt.Println("success")
}
