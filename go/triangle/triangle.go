// Package triangle provides the Kind type and related functions
// to represent different types of triangles.
package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}
	sort.Float64s(sides[:])

	for _, val := range sides {
		if val <= 0 || math.IsNaN(val) || math.IsInf(val, 0) {
			return NaT
		}
	}

	if !verifyTriangleInequality(sides) {
		return NaT
	}

	switch {
	case isEquilateral(sides):
		return Equ
	case isIsoceles(sides):
		return Iso
	}
	return Sca
}

// verifyTrangleInequality assumes that the slice is sorted in increasing order
func verifyTriangleInequality(sides [3]float64) bool {
	return sides[2] <= sides[1]+sides[0]
}

func isEquilateral(sides [3]float64) bool {
	return sides[0] == sides[1] && sides[1] == sides[2]
}

func isIsoceles(sides [3]float64) bool {
	return sides[0] == sides[1] ||
		sides[1] == sides[2] ||
		sides[2] == sides[0]
}
