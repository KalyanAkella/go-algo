package main

import (
	"fmt"
)

// O(N) space - used has set to filter duplicates
func perm(str string) map[string]struct{} {
	ans := make(map[string]struct{})
	if len(str) == 0 {
		return ans
	}

	ch := str[0]
	subPerms := perm(str[1:])
	if len(subPerms) == 0 {
		ans[string(ch)] = struct{}{}
	} else {
		for subPerm := range subPerms {
			for i := 0; i <= len(subPerm); i++ {
				perm := subPerm[0:i] + string(ch) + subPerm[i:]
				ans[perm] = struct{}{}
			}
		}
	}
	return ans
}

func main() {
	perms := perm("aba")
	fmt.Println(perms)
}
