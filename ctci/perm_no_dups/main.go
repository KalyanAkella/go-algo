package main

import (
	"container/list"
	"fmt"
)

func perm(str string) *list.List {
	if len(str) == 0 {
		return list.New()
	}

	ch := str[0]
	subPerms := perm(str[1:])
	ans := list.New()
	if subPerms.Len() == 0 {
		ans.PushBack(string(ch))
	} else {
		for e := subPerms.Front(); e != nil; e = e.Next() {
			subPerm := e.Value.(string)
			for i := 0; i <= len(subPerm); i++ {
				perm := subPerm[0:i] + string(ch) + subPerm[i:]
				ans.PushBack(perm)
			}
		}
	}
	return ans
}

// T(N) = T(N-1) + (N) * (N-1)! = T(N-1) + N!
// T(N-1) = T(N-2) + (N-1)!
// ...
// T(1) = T(0) + 1!
// T(0) = C
// T(N) = C + 1! + 2! + 3! + ... + N!

func printList(perms *list.List) {
	for e := perms.Front(); e != nil; e = e.Next() {
		perm := e.Value.(string)
		fmt.Printf("%s ", perm)
	}
}

func main() {
	perms := perm("abcdef")
	printList(perms)
	fmt.Println()
}
