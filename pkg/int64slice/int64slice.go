package int64slice

import (
	"sort"
)

// Int64Slice is []int64 that can be sorted (ascending).
type Int64Slice []int64

// Len implements sort.Interface.
func (s Int64Slice) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s Int64Slice) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

// Swap implements sort.Interface.
func (s Int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Search returns the index of element x in s, -1 otherwise.
func (s Int64Slice) Search(x int64) int {
	index := sort.Search(len(s), func(i int) bool {
		return s[i] >= x
	})
	if index >= len(s) || s[index] != x {
		return -1
	}
	return index
}

// Insert inserts element x in s preserving order.
func (s *Int64Slice) Insert(x int64) int {
	index := sort.Search(len(*s), func(i int) bool {
		return (*s)[i] >= x
	})
	*s = append(*s, 0)
	copy((*s)[index+1:], (*s)[index:])
	(*s)[index] = x

	return index
}

// Remove removes element x of s, if present.
func (s *Int64Slice) Remove(x int64) bool {
	i := s.Search(x)
	if i == -1 {
		return false
	}
	*s = append((*s)[:i], (*s)[i+1:]...) // Remove preserving order.
	return true
}

// Copy returns a copy of slice s.
func (s Int64Slice) Copy() Int64Slice {
	cp := make(Int64Slice, len(s))
	copy(cp, s)
	return cp
}
