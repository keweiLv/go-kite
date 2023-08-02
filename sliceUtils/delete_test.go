package sliceUtils

import (
	"fmt"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	newArr := RemoveAtIndex(arr, 2)
	fmt.Printf("删除后结果为:%v \n", newArr)
	if len(newArr) != 4 {
		t.Errorf("Expected length 4, but got %v", len(newArr))
	}
	if newArr[2] != 4 {
		t.Errorf("Expected 4 at index 2, but got %v", newArr[2])
	}
}

func TestShrinkSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr = arr[:2]
	newArr := ShrinkSlice(arr)
	fmt.Printf("缩容后后结果为:%v，容量为:%v \n", newArr, cap(newArr))
	if cap(newArr) != 2 {
		t.Errorf("Expected capacity 2, but got %v", cap(newArr))
	}
}
