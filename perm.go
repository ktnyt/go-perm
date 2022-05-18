package perm

// Permuter represents a slice permuter. Permutations of the slice will be
// computed using the conting quickperm algorithm:
//    https://www.quickperm.org/quickperm.php
type Permuter interface {
	// Next returns true as long as there is an unseen permutation left. The
	// first call to Next will keep the original slice untouched.
	Next() bool
}

type countingPermuter[T any] struct {
	a []T
	p []int
	i int
}

// Permute returns a Permuter initialized with the given slice. The slice is
// permuted in-place. As long as Next returns true, the slice variable passed
// to Permute will contain an unseen permutation of the slice.
func Permute[T any](a []T) Permuter {
	return &countingPermuter[T]{a: a}
}

func (p *countingPermuter[T]) Next() bool {
	if p.p == nil {
		p.p = make([]int, len(p.a))
		p.i = 1
		return true
	}
	for p.i < len(p.a) {
		if p.p[p.i] < p.i {
			j := 0
			if p.i%2 == 1 {
				j = p.p[p.i]
			}
			p.a[p.i], p.a[j] = p.a[j], p.a[p.i]
			p.p[p.i]++
			p.i = 1
			return true
		} else {
			p.p[p.i] = 0
			p.i++
		}
	}
	return false
}
