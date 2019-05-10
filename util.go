package main

func InsertStringSliceCopy(slice, insertion []*Point, index int) []*Point {
	result := make([]*Point, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
