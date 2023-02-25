package main

import (
	"fmt"
)

func main() {
	input := "amanaplanacanalpanama"
	fmt.Println("1.reverseString: ", reverseString([]byte(input)))

	fmt.Println("2.isPalindrome: ", isPalindrome(input))

	fmt.Println("3.calculateFactorial: ", calculateFactorial(10))

	fmt.Println("4.findGCD: ", findGCD([]int{2,5,6,9,10}))

	fmt.Println("5.findLCM: ", findLCM(24, 15))

	fmt.Println("6.isPrime: ", isPrime(3))

	fmt.Println("7.createFibonacciseries: ", createFibonacciseries(3))

	fmt.Println("8.bubbleSort: ", bubbleSort([]int{5,3,8,4,6}))

	fmt.Println("9.sumArray: ", sumArray([]int{5,3,8,4,6}))

	fmt.Println("10.findMaxElement: ", findMaxElement([]int{5,3,8,4,6}))
	
}

func reverseString(s []byte) string{
    var tm byte
    for i := 0 ; i < len(s)/2; i++ {
        tm = s[i]
        s[i] = s[len(s) - (i+1)]
        s[len(s) - (i+1)] = tm
    }
	return string(s)
}

func isPalindrome(s string) bool{

	i := 0
	j := len(s) - 1
	b := []byte(s)

	for i <= j {
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func calculateFactorial(n int) (output int){
	output = 1
	for i := n; i >= 1; i-- {
		output *= i
	}
	return
}

func findGCD(nums []int) int {
    min := nums[0]
    max := nums[0]
    for i := 0; i < len(nums); i++ {
        if nums[i] < min {
            min = nums[i]
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    for i := max; i >= 1; i-- {
        if max % i == 0 && min % i == 0{
            return i
        }
    }
    return 1
}

func findLCM(num1 int, num2 int) int {
	GCD := findGCD([]int{num1, num2})
	result := (num1*num2)/GCD

	return result
}

func isPrime(num int) bool{
	for i := 2; i < num/2; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func createFibonacciseries(num int) ([]int){
	fibonacci := make([]int, num)
	fibonacci[0] = 0
	fibonacci[1] = 1
	i := 2
	for i < num {

		fibonacci[i] = fibonacci[i - 1] + fibonacci[i - 2]
		i++
	}
	return fibonacci
}

func bubbleSort(nums []int) ([]int) {
	//5,3,8,4,6
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}

func sumArray(nums []int) int {
	sum := 0
	for _,item := range  nums {
		sum += item
	}
	return sum
}

func findMaxElement(nums []int) int {
	nums = bubbleSort(nums)
	return nums[len(nums) - 1]
}