package stdlib

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func StringOperations() string {
	s := "hello world"

	upper := strings.ToUpper(s)
	lower := strings.ToLower("GO")
	contains := strings.Contains(s, "world")
	hasPrefix := strings.HasPrefix(s, "hello")
	hasSuffix := strings.HasSuffix(s, "world")

	replace := strings.Replace(s, "world", "golang", 1)
	split := strings.Split("a,b,c", ",")
	join := strings.Join([]string{"go", "java", "python"}, "-")
	trim := strings.TrimSpace("  hello  ")
	repeat := strings.Repeat("a", 3)

	return fmt.Sprintf("%s|%s|%v|%v|%v|%s|%v|%s|%s|%s",
		upper, lower, contains, hasPrefix, hasSuffix, replace, split, join, trim, repeat)
}

func SortOperations() []int {
	nums := []int{5, 2, 8, 1, 9}
	sort.Ints(nums)

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)

	exists := sort.SearchInts(nums, 5)
	_ = exists

	return nums
}

func SliceOperations() bool {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	same := slices.Equal(a, b)

	slices.Sort(a)
	slices.Reverse(a)
	clone := slices.Clone(b)
	slices.Reverse(clone)

	return same && slices.Equal(a, clone)
}
