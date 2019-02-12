package main

import (
	"container/list"
	"fmt"
)

func move(disks int, src, dst, buf *list.List) {
	if disks < 1 {
		return
	}
	move(disks-1, src, buf, dst)
	moveTop(src, dst)
	move(disks-1, buf, dst, src)
}

func moveTop(src, dst *list.List) {
	disk := src.Remove(src.Front())
	dst.PushFront(disk)
}

func main() {
	disks := []int{4, 3, 2, 1}
	src := list.New()
	for _, disk := range disks {
		src.PushFront(disk)
	}
	dst := list.New()
	move(len(disks), src, dst, list.New())
	for e := dst.Front(); e != nil; e = e.Next() {
		disk := e.Value.(int)
		fmt.Printf("%d ", disk)
	}
	fmt.Println()
}
