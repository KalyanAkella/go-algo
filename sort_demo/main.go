package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Keys []string

func (s Keys) Len() int { return len(s) }

func (s Keys) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Keys) Less(i, j int) bool {
	first_elements, second_elements := strings.Split(s[i], "_"), strings.Split(s[j], "_")
	first_time, first_err := strconv.Atoi(first_elements[1])
	if first_err != nil {
		panic(first_err)
	}
	second_time, second_err := strconv.Atoi(second_elements[1])
	if second_err != nil {
		panic(second_err)
	}
	return first_time < second_time
}

func split_old_new(db_keys []string) ([]string, []string) {
	var old_keys, new_keys []string
	for _, k := range db_keys {
		if strings.HasPrefix(k, "Old_") {
			old_keys = append(old_keys, k)
		} else if strings.HasPrefix(k, "New_") {
			new_keys = append(new_keys, k)
		}
	}
	return old_keys, new_keys
}

func main() {
	input := []string{
		"Old_1532815260",
		"Old_1532810040",
		"Old_1532786580",
		"New_1532809920",
		"New_1532778480",
		"New_1532797020",
		"Old_1532815920",
		"Old_1532780040",
		"New_1532803080",
		"Old_1532809920",
		"Old_1532804400",
		"New_1532810040",
		"New_1532804580",
	}
	old, new := split_old_new(input)
	sort.Sort(Keys(old))
	sort.Sort(Keys(new))
	fmt.Println(old)
	fmt.Println(new)
	fmt.Println(old[:len(old)-1])
	fmt.Println(new[:len(new)-1])
	m := map[string]struct{}{"": struct{}{}, "abc": struct{}{}}
	fmt.Println(m)
	delete(m, "")
	fmt.Println(m)
}
