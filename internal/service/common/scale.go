package common

import "fmt"

type Scale struct {
	Scale          int
	EffectiveDigit int
}

func (s Scale) String() string {
	return fmt.Sprintf("scale = %d\neffective_digit = %d\n", s.Scale, s.EffectiveDigit)
}
