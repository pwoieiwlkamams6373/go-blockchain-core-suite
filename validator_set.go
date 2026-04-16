package main

import "fmt"

type ValidatorSet struct {
	list []*Validator
}

func NewValidatorSet() *ValidatorSet {
	return &ValidatorSet{}
}

func (s *ValidatorSet) Add(v *Validator) {
	s.list = append(s.list, v)
}

func (s *ValidatorSet) Remove(addr string) bool {
	for i, v := range s.list {
		if v.Address == addr {
			s.list = append(s.list[:i], s.list[i+1:]...)
			return true
		}
	}
	return false
}

func (s *ValidatorSet) Count() int {
	return len(s.list)
}

type Validator struct {
	Address string
}

func main() {
	set := NewValidatorSet()
	fmt.Println("验证节点集合管理器就绪")
}
