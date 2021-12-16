package day16

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
	bits := convert(strings.TrimSpace(input))
	pkt, _ := newParser(bits).Parse()
	if part == 1 {
		return fmt.Sprint(sumVersions(pkt)), nil
	}
	return fmt.Sprint(evaluate(pkt)), nil
}

func convert(input string) string {
	bs := make([]byte, 4*len(input))
	for i, r := range input {
		var nibble uint8
		if r >= '0' && r <= '9' {
			nibble = uint8(r - '0')
		} else {
			// r >= 'A' && r <= 'F'
			nibble = uint8(r - 'A' + 10)
		}
		for b := 3; b >= 0; b-- {
			idx := 4*i + (3 - b)
			if nibble&(1<<b) > 0 {
				bs[idx] = '1'
			} else {
				bs[idx] = '0'
			}
		}
	}
	return string(bs)
}

const (
	typeIDSum         = 0
	typeIDProduct     = 1
	typeIDMinimum     = 2
	typeIDMaximum     = 3
	typeIDLiteral     = 4
	typeIDGreaterThan = 5
	typeIDLessThan    = 6
	typeIDEqualTo     = 7
)

type packet struct {
	// Standard header.
	Version int
	TypeID  int

	// Only set if TypeID == typeIDLiteral.
	Literal int

	// Only set if TypeID != typeIDLiteral.
	Operator   int
	Subpackets []packet
}

type parser struct {
	bits string
}

func newParser(bits string) *parser {
	return &parser{
		bits: bits,
	}
}

func (p *parser) Parse() (packet, string) {
	var pkt packet
	pkt.Version = p.read(3)
	pkt.TypeID = p.read(3)
	if pkt.TypeID == typeIDLiteral {
		keepGoing := true
		for keepGoing {
			pkt.Literal <<= 4
			keepGoing = p.read(1) == 1
			pkt.Literal += p.read(4)
		}
		return pkt, p.bits
	}
	lengthTypeID := p.read(1)
	if lengthTypeID == 0 {
		subpktLength := p.read(15)
		subBits := p.skip(subpktLength)
		for subBits != "" {
			var subpkt packet
			subpkt, subBits = newParser(subBits).Parse()
			pkt.Subpackets = append(pkt.Subpackets, subpkt)
		}
	} else {
		subpktCount := p.read(11)
		pkt.Subpackets = make([]packet, subpktCount)
		for i := 0; i < subpktCount; i++ {
			subpkt, bits := newParser(p.bits).Parse()
			pkt.Subpackets[i] = subpkt
			p.bits = bits
		}
	}
	return pkt, p.bits
}

func (p *parser) read(n int) int {
	return btoi(p.skip(n))
}

func (p *parser) skip(n int) string {
	s := p.bits[:n]
	p.bits = p.bits[n:]
	return s
}

func btoi(bits string) int {
	i := 0
	for _, r := range bits {
		i <<= 1
		if r == '1' {
			i |= 1
		}
	}
	return i
}

func sumVersions(pkt packet) int {
	sum := pkt.Version
	for _, subpkt := range pkt.Subpackets {
		sum += sumVersions(subpkt)
	}
	return sum
}

func evaluate(pkt packet) int {
	switch pkt.TypeID {
	case typeIDSum:
		sum := 0
		for _, subpkt := range pkt.Subpackets {
			sum += evaluate(subpkt)
		}
		return sum

	case typeIDProduct:
		product := 1
		for _, subpkt := range pkt.Subpackets {
			product *= evaluate(subpkt)
		}
		return product

	case typeIDMinimum:
		vals := make([]int, len(pkt.Subpackets))
		for i, subpkt := range pkt.Subpackets {
			vals[i] = evaluate(subpkt)
		}
		return intmath.Min(vals[0], vals[1:]...)

	case typeIDMaximum:
		vals := make([]int, len(pkt.Subpackets))
		for i, subpkt := range pkt.Subpackets {
			vals[i] = evaluate(subpkt)
		}
		return intmath.Max(vals[0], vals[1:]...)

	case typeIDLiteral:
		return pkt.Literal

	case typeIDGreaterThan:
		if evaluate(pkt.Subpackets[0]) > evaluate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}

	case typeIDLessThan:
		if evaluate(pkt.Subpackets[0]) < evaluate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}

	case typeIDEqualTo:
		if evaluate(pkt.Subpackets[0]) == evaluate(pkt.Subpackets[1]) {
			return 1
		} else {
			return 0
		}

	default:
		panic(fmt.Errorf("invalid type ID: %d", pkt.TypeID))
	}
}
