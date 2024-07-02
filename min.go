package funky

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](min, max: T) T {
	if min < max {
		return min
	}
	return max
}