package stack

import "github.com/reconquest/pkg/log"

type Stack interface {
	Push(i int)
	GetTop() int
	GetBottom() int
	Size() int
	DeleteOne(i int)
	Clear()
}

func NewSlice(slice []int) *Int {
	return &Int{
		Slice: slice,
	}
}

type Int struct {
	Slice []int
}

func (s *Int) Push(i int) {
	s.Slice = append(s.Slice, i)
}

func (s *Int) GetTop() int {
	return s.Slice[0]
}

func (s *Int) GetBottom() int {
	return s.Slice[len(s.Slice)-1]
}

func (s *Int) DeleteOne(i int) {
	index, result := findElement(s.Slice, i)
	if result {
		s.Slice = append(s.Slice[:index], s.Slice[index+1:]...)
		return
	}

	log.Warning("wrong number")
}

func (s *Int) Size() int {
	return len(s.Slice)
}

func (s *Int) Clear() {
	s.Slice = []int{}
}

func findElement(slice []int, element int) (int, bool) {
	for index, item := range slice {
		if element == item {
			return index, true
		}
	}

	return 0, false
}
