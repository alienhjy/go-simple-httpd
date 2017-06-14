package main

import (
	"fmt"
)

type StringFlags []string

func (opt *StringFlags) String() string {
	return fmt.Sprint(*opt)
}

func (opt *StringFlags) Set(value string) error {
	*opt = append(*opt, value)
	return nil
}

