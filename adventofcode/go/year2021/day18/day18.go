package day18

import (
	"fmt"
	"strings"

	"github.com/Saser/pdp/adventofcode/go/intmath"
)

func Part1(input string) (string, error) {
	return solve(input, 1)
}

func Part2(input string) (string, error) {
	return solve(input, 2)
}

func solve(input string, part int) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if part == 1 {
		depth := 0
		result := newNumber()
		result.Parse(lines[0], depth)
		for _, line := range lines[1:] {
			result.Parse(line, depth)
			result.Reduce(depth - 1)
			depth--
		}
		return fmt.Sprint(result.Magnitude()), nil
	}
	numbers := make([]*number, len(lines))
	for i, line := range lines {
		n := newNumber()
		n.Parse(line, 0)
		numbers[i] = n
	}
	max := 0
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 == n2 {
				continue
			}
			result := n1.Copy()
			result.Add(n2)
			result.Reduce(-1)
			max = intmath.Max(max, result.Magnitude())
		}
	}
	return fmt.Sprint(max), nil
}

type regular struct {
	Value int
	Depth int
}

type number []regular

func newNumber() *number {
	// I made a guess that each line contains at most 32 regular numbers. By
	// allocating once for twice that amount, we can add numbers by parsing them
	// into the same slice and avoid reallocating, saving memory and runtime.
	const maxRegulars = 32
	regulars := make([]regular, 0, 2*maxRegulars)
	return (*number)(&regulars)
}

func (n *number) Parse(s string, depth int) {
	for _, r := range s {
		switch r {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			// Do nothing.
		default:
			*n = append(*n, regular{
				Value: int(r - '0'),
				Depth: depth,
			})
		}
	}
}

func (n *number) Copy() *number {
	n2 := make([]regular, len(*n))
	copy(n2, *n)
	return (*number)(&n2)
}

func (n *number) Add(other *number) {
	*n = append(*n, (*other)...)
}

func (n *number) Reduce(depth int) {
	// First make sure all pairs that can explode have done so. We can do that
	// in one iteration over the list of regular numbers, using the fact that no
	// regular number will have a depth of more than 5 [*]. A pair that explodes
	// is replaced by a regular number which will be part of a pair which will
	// never be nested inside 4 other pairs. So we can just explode one pair and
	// then move on.
	//
	// [*]: The instructions say that all exploding pairs will only contain
	// regular numbers. If there was a regular number of a depth of more than 5,
	// then it would be part of a pair that is nested inside more than 4 other
	// pairs.  But then there would be some pair nested inside exactly 4 other
	// pairs which would have non-regular numbers as children -- which the
	// instructions say cannot happen.
	for i := 0; i < len(*n); i++ {
		r := &(*n)[i]
		if r.Depth-depth != 5 {
			continue
		}
		// r is the left number in an exploding pair. Add its value to the
		// previous number (if any), add the next number (which is the right
		// number in the exploding pair) to its subsequent number (if any),
		// and then replace the left and right numbers with a single 0 at a
		// depth of one less than the exploding numbers.
		left := r
		right := (*n)[i+1]
		if prev := i - 1; prev >= 0 {
			(*n)[prev].Value += left.Value
		}
		if next := i + 2; next < len(*n) {
			(*n)[next].Value += right.Value
		}
		r.Value = 0
		r.Depth--
		*n = append((*n)[:i+1], (*n)[i+2:]...)
	}

	// Now we know that the next action (if any) will be a split. Splitting
	// could result in new pairs that should explode, but these pairs will be
	// created in the same place as the number that is split, so we can do a
	// split and the resulting explosion at once. However, the explosion could
	// mean that a number we skipped over should be split. Example:
	//     Initially: [[[[9, 10], 8], ...]]
	//     Split:     [[[[9, [5,5]], 8], ...]]
	//     Explode:   [[[[14, 0], 13], ...]]
	// So whenever we split-and-explode, we have to take one step back (if
	// possible).
	i := 0
	for i < len(*n) {
		r := &(*n)[i]
		if r.Value < 10 {
			i++
			continue
		}
		lo := r.Value / 2
		hi := r.Value/2 + r.Value%2
		if r.Depth-depth == 4 {
			// The split number is at depth 5, which will create a pair at
			// depth 5, which should then explode. We do the explosion by
			// simply going to the previos and next numbers (if any) and
			// adding lo and hi. Then the split number is replaced by a 0.
			if prev := i - 1; prev >= 0 {
				(*n)[prev].Value += lo
			}
			if next := i + 1; next < len(*n) {
				(*n)[next].Value += hi
			}
			r.Value = 0
			if i > 0 {
				i-- // Take one step back, as explained above.
			} else {
				i++
			}
		} else {
			// The split number results in a pair that doesn't immediately
			// explode. We represent this by replacing the split number with lo,
			// adding a new regular number with hi, and set both of these
			// numbers' depth to the split number depth plus one.
			r.Value = lo
			r.Depth++
			r2 := regular{
				Value: hi,
				Depth: r.Depth,
			}
			*n = append((*n)[:i+1], append([]regular{r2}, (*n)[i+1:]...)...)
			// This check seemed weird to me at first, but it's entirely
			// possible that a series of splits and explosions leads to a bunch
			// of larger-than-10 values accumulate close to each other. When
			// these values are split and exploded, some may end up being 20 or
			// larger, and therefore needs several splits. So we can only move
			// on to the next regular number if the new value of r is less than
			// 10.
			if r.Value < 10 {
				i++
			}
		}
	}
}

func (n *number) Magnitude() int {
	var stack []regular
	for _, r := range *n {
		stack = append(stack, r)
		for len(stack) >= 2 {
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			if left.Depth != right.Depth {
				break
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, regular{
				Value: 3*left.Value + 2*right.Value,
				Depth: left.Depth - 1,
			})
		}
	}
	return stack[0].Value
}

func (n *number) String() string {
	type item struct {
		depth int
		s     string
	}
	var stack []item
	for _, r := range *n {
		stack = append(stack, item{
			depth: r.Depth,
			s:     fmt.Sprint(r.Value),
		})
		for len(stack) >= 2 {
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			if left.depth != right.depth {
				break
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, item{
				depth: left.depth - 1,
				s:     "[" + left.s + "," + right.s + "]",
			})
		}
	}
	return stack[0].s
}
