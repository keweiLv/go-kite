package sliceUtils

import "reflect"

func RemoveAtIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func ShrinkSlice[T any](s []T) []T {
	// We use reflect package to shrink the underlying array
	reflectValue := reflect.ValueOf(s)
	s = reflect.MakeSlice(reflectValue.Type(), len(s), len(s)).Interface().([]T)
	copy(s, reflectValue.Interface().([]T))
	return s
}
