package day18

/*
Some ideas on how this could maybe be made faster, based on what I've found when
profiling.

1. Use some kind of queue (maybe a priority queue) for pairs that can explode.
   Most of the time in Explode is spent looking for a pair to explode (if any).
	- In Add, clear and reinitialize the queue after joining the two trees.
	  Depths need to be recalculated so the contents of the queue will probably
	  change.
	- In Explode, pop the first item from the queue (if any), otherwise return
	  early.
	- In Split, check the depth of the just created pair, and add it to the
	  queue. The difficult part is figuring out where in the queue it should be
	  added.

2. Use a priority queue for regular numbers that can be split. All >=10 numbers
   are prioritized over all <10 numbers, and among >=10 numbers the
   left-to-right order decides priority. The difficult part is how to determine
   this order. One idea could be to assign them left-to-right indices in init()
   and then use that in the priority queue implementation.
*/

import (
	"fmt"
	"strconv"
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
	numbers := parse(input)
	if part == 1 {
		result := sumAll(numbers)
		return fmt.Sprint(result.Magnitude()), nil
	}
	max := 0
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 == n2 {
				continue
			}
			result := n1.Clone()
			result.Add(n2.Clone())
			result.Reduce()
			max = intmath.Max(max, result.Magnitude())
		}
	}
	return fmt.Sprint(max), nil
}

func sumAll(numbers []*number) *number {
	result := numbers[0]
	for _, other := range numbers[1:] {
		result.Add(other)
		result.Reduce()
	}
	return result
}

func parse(input string) []*number {
	var numbers []*number
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		numbers = append(numbers, parseNumber(line))
	}
	return numbers
}

func parseNumber(line string) *number {
	return makeNumber(parseNode(line))
}

func parseNode(s string) *node {
	n, rem := parseNodeAux(s)
	if rem != "" {
		panic(fmt.Errorf("parsing %q returned non-empty remainder %q", s, rem))
	}
	return n
}

func parseNodeAux(s string) (*node, string) {
	if s[0] == '[' {
		left, rem := parseNodeAux(s[1:])    // skip the '['
		right, rem := parseNodeAux(rem[1:]) // skip the ','
		rem = rem[1:]                       // skip the ']'
		return &node{
			Left:  left,
			Right: right,
		}, rem
	}
	var (
		v   int
		rem string
		err error
	)
	end := strings.IndexAny(s, "[],")
	if end == -1 {
		v, err = strconv.Atoi(s)
		rem = ""
	} else {
		v, err = strconv.Atoi(s[:end])
		rem = s[end:]
	}
	if err != nil {
		panic(err)
	}
	return &node{
		Value: v,
	}, rem
}

func makeNumber(root *node) *number {
	n := &number{root: root}
	n.init()
	return n
}

func assignElements(n *number, nd *node) {
	if nd.regular {
		n.pushBack(nd)
		return
	}
	assignElements(n, nd.Left)
	assignElements(n, nd.Right)
}

type number struct {
	root *node

	front, back *node
}

func (n *number) init() {
	n.front = nil
	n.back = nil
	n.root.init(0)
	n.assignLinks(n.root)
}

func (n *number) assignLinks(nd *node) {
	if nd.regular {
		n.pushBack(nd)
		return
	}
	n.assignLinks(nd.Left)
	n.assignLinks(nd.Right)
}

func (n *number) Clone() *number {
	n2 := &number{root: n.root.clone()}
	n2.init()
	return n2
}

func (n *number) String() string {
	return n.root.String() + "\n" + n.ListString()
}

func (n *number) ListString() string {
	var rs []string
	for r := n.front; r != nil; r = r.next {
		rs = append(rs, fmt.Sprint(r.Value))
	}
	return "[" + strings.Join(rs, ",") + "]"
}

func (n *number) Magnitude() int {
	return n.root.magnitude()
}

func (n *number) Add(other *number) {
	// Create a new root by setting n to the left and other to the right
	// elements.
	root := &node{
		Parent: nil,
		Left:   n.root,
		Right:  other.root,
	}
	root.Left.Parent = root
	root.Right.Parent = root
	root.init(0)
	n.root = root
	// Join n's linked list of regular numbers with other's list.
	n.back.next = other.front
	other.front.prev = n.back
	n.back = other.back
	// Other shouldn't be used after calling Add, but if it is used (due to a
	// bug), fail loudly by panicking due to nil derefs.
	other.root = nil
	other.front = nil
	other.back = nil
}

func (n *number) Reduce() {
	for {
		if changed := n.reduceStep(); !changed {
			break
		}
	}
}

func (n *number) reduceStep() bool {
	if n.Explode() {
		return true
	}
	if n.Split() {
		return true
	}
	return false
}

func (n *number) Explode() bool {
	var exploding *node
	for r := n.front; r != nil; r = r.next {
		if parent := r.Parent; parent.Left.regular && parent.Right.regular && parent.depth == 4 {
			exploding = parent
			break
		}
	}
	if exploding == nil {
		return false
	}
	// Now that we have found the exploding pair, increase the values of the
	// previous and next regular numbers.
	left := exploding.Left
	right := exploding.Right
	if prev := left.prev; prev != nil {
		prev.Value += left.Value
	}
	if next := right.next; next != nil {
		next.Value += right.Value
	}
	// After increasing the surrounding values, we replace the exploding pair
	// with a regular 0.
	exploding.Value = 0
	exploding.Left = nil
	exploding.Right = nil
	exploding.regular = true
	// Now we can remove the left and right regular numbers from the linked
	// list, and insert the 0 in their place. We do this by placing the 0
	// between them and then removing them. Why between? To make the code
	// shorter by not dealing with edge cases at the front or back of the list.
	n.insertAfter(exploding, left)
	n.remove(left)
	n.remove(right)
	return true
}

func (n *number) Split() bool {
	var split *node
	for r := n.front; r != nil; r = r.next {
		if r.Value >= 10 {
			split = r
			break
		}
	}
	if split == nil {
		return false
	}
	// Replace the split node with a pair.
	split.Left = &node{
		Parent:  split,
		Value:   split.Value / 2,
		regular: true,
		depth:   split.depth + 1,
	}
	split.Right = &node{
		Parent:  split,
		Value:   split.Value/2 + split.Value%2,
		regular: true,
		depth:   split.depth + 1,
	}
	split.Value = 0
	split.regular = false
	// Remove the split node from the linked list and insert the two new
	// elements in its place. Do this by placing the new elements around the
	// split node and then remove the split node. Why around? To make the code
	// shorter by not having to deal with edge cases where split is at the front
	// or back of the list.
	n.insertBefore(split.Left, split)
	n.insertAfter(split.Right, split)
	n.remove(split)
	return true
}

func (n *number) pushBack(r *node) {
	if n.back == nil {
		r.prev = nil
		r.next = nil
		n.front = r
		n.back = r
		return
	}
	n.back.next = r
	r.prev = n.back
	r.next = nil
	n.back = r
}

func (n *number) insertBefore(r, mark *node) {
	if mark == n.front {
		r.prev = nil
		r.next = mark
		mark.prev = r
		n.front = r
		return
	}
	r.prev = mark.prev
	r.next = mark
	mark.prev.next = r
	mark.prev = r
}

func (n *number) insertAfter(r, mark *node) {
	if mark == n.back {
		r.prev = mark
		r.next = nil
		mark.next = r
		n.back = r
		return
	}
	r.prev = mark
	r.next = mark.next
	mark.next.prev = r
	mark.next = r
}

func (n *number) remove(mark *node) {
	if mark == n.front && mark == n.back {
		n.front = nil
		n.back = nil
		mark.prev = nil
		mark.next = nil
		return
	}
	if mark == n.front {
		n.front = mark.next
		mark.next.prev = nil
		mark.prev = nil
		mark.next = nil
		return
	}
	if mark == n.back {
		n.back = mark.prev
		mark.prev.next = nil
		mark.prev = nil
		mark.next = nil
		return
	}
	mark.prev.next = mark.next
	mark.next.prev = mark.prev
	mark.prev = nil
	mark.next = nil
}

type node struct {
	Parent *node

	// These are either both nil or both non-nil. If both are nil, then this is
	// a regular number. If both are non-nil, then this is a pair.
	Left, Right *node

	// These fields are only valid if this is a regular number.
	Value int
	// The previous and next nodes in the linked list of regular numbers. These
	// fields can be nil, and are always updated by the operations on number
	// (Add, pushBack, insertBefore, insertAfter, remove).
	prev, next *node

	regular bool
	depth   int
}

func (n *node) pair() bool {
	return n.Left != nil && n.Right != nil
}

func (n *node) magnitude() int {
	if n.regular {
		return n.Value
	}
	return 3*n.Left.magnitude() + 2*n.Right.magnitude()
}

func (n *node) init(depth int) {
	n.depth = depth
	if n.Left == nil && n.Right == nil {
		n.regular = true
		return
	}
	n.Left.Parent = n
	n.Right.Parent = n
	n.Left.init(depth + 1)
	n.Right.init(depth + 1)
}

func (n *node) clone() *node {
	if n.regular {
		return &node{Value: n.Value}
	}
	return &node{
		Left:  n.Left.clone(),
		Right: n.Right.clone(),
	}
}

func (n *node) String() string {
	if n.regular {
		return fmt.Sprint(n.Value)
	}
	return "[" + n.Left.String() + "," + n.Right.String() + "]"
}
