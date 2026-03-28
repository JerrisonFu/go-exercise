package testing

import (
	"fmt"
	"strings"
)

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}
	return count
}

func FindDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)

	for _, n := range nums {
		if seen[n] {
			duplicates[n] = true
		}
		seen[n] = true
	}

	result := make([]int, 0, len(duplicates))
	for k := range duplicates {
		result = append(result, k)
	}
	return result
}

func GroupAnagrams(strs []string) map[string][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}
	return groups
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		complement := target - n
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}

func RunAll() {
	fmt.Println("=== Testing Examples ===")

	fmt.Println("\n--- Fibonacci ---")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fib(%d) = %d (iterative: %d)\n", i, Fib(i), FibIterative(i))
	}

	fmt.Println("\n--- Prime Check ---")
	primes := []int{2, 3, 5, 7, 11, 13, 15, 17, 19, 23}
	for _, p := range primes {
		fmt.Printf("%d is prime: %v\n", p, IsPrime(p))
	}

	fmt.Println("\n--- String Reverse ---")
	testStr := "Hello, World!"
	fmt.Printf("Reverse('%s') = '%s'\n", testStr, ReverseString(testStr))

	fmt.Println("\n--- Word Count ---")
	text := "hello world hello go world"
	counts := WordCount(text)
	fmt.Printf("Word count: %v\n", counts)

	fmt.Println("\n--- Find Duplicates ---")
	nums := []int{1, 2, 3, 2, 4, 3, 5}
	fmt.Printf("Duplicates in %v: %v\n", nums, FindDuplicates(nums))

	fmt.Println("\n--- Two Sum ---")
	targetNums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(targetNums, target)
	fmt.Printf("TwoSum(%v, %d) = %v\n", targetNums, target, result)

	fmt.Println("\n--- Group Anagrams ---")
	anagrams := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := GroupAnagrams(anagrams)
	fmt.Printf("Grouped: %v\n", grouped)
}
