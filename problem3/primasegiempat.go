package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func PrimaSegiEmpat(wide, high, start int) string {
	var primes []int
	num := start
	
	for len(primes) < wide*high {
		if isPrime(num) {
			if num > start{
				primes = append(primes, num)
			}
			
		}
		num++
	}

	result := ""
	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			index := i*wide + j
			if index < len(primes) {
				result += fmt.Sprintf("%d\t", primes[index])
			}
		}
		result += "\n"
	}
	sum := 0
    for _,prime := range primes{
		sum +=prime

	}
 
	result += fmt.Sprintf("%d", sum)
	return result
}

func main() {
	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// Output:
	// 17	19
	// 23	29
	// 31	37
	// 156

	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// Output:
	// 2	3	5	7	11
	// 13	17	19	23	29
	// 129
}
