package funky

import "golang.org/x/exp/constraints"

func Clamp[T constraints.Ordered](n, min, max: T) T {
	return Max(min, Min(max, v))
}