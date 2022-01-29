package frechet

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type point struct {
	x float64
	y float64
}

func TestFrechet(t *testing.T) {
	testCases := []struct {
		desc    string
		s       interface{}
		t       interface{}
		df      DistanceFunc
		dist    float64
		noError bool
	}{
		{
			desc: "int series",
			s:    []int{1, 2, 3},
			t:    []int{2, 3, 4},
			df: func(x, y interface{}) float64 {
				xx := x.(int)
				yy := y.(int)
				return math.Abs(float64(xx - yy))
			},
			dist:    1,
			noError: true,
		},
		{
			desc: "float series",
			s:    []float64{1.0, 2.0, 3.0},
			t:    []float64{2.0, 3.0, 4.0},
			df: func(x, y interface{}) float64 {
				xx := x.(float64)
				yy := y.(float64)
				return math.Abs(xx - yy)
			},
			dist:    1,
			noError: true,
		},
		{
			//  y
			//  ^
			//  |         t
			//  |         /
			//  |      /     s
			//  |   /      /
			//  |/      /
			//  |    /
			//  | /
			//  +---------------------> x
			desc: "2D series",
			s:    []point{{0, 0}, {1, 1}, {2, 2}},
			t:    []point{{0, 1}, {1, 2}, {2, 3}},
			df: func(x, y interface{}) float64 {
				p1 := x.(point)
				p2 := y.(point)
				return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
			},
			dist:    1,
			noError: true,
		},
		{
			desc: "same series",
			s:    []int{1, 2, 3},
			t:    []int{1, 2, 3},
			df: func(x, y interface{}) float64 {
				xx := x.(int)
				yy := y.(int)
				return math.Abs(float64(xx - yy))
			},
			dist:    0,
			noError: true,
		},
		{
			desc: "series with different length",
			s:    []int{1, 2, 3},
			t:    []int{1, 2, 3, 4},
			df: func(x, y interface{}) float64 {
				xx := x.(int)
				yy := y.(int)
				return math.Abs(float64(xx - yy))
			},
			dist:    1,
			noError: true,
		},
		{
			desc: "empty series",
			s:    []int{},
			t:    []int{},
			df: func(x, y interface{}) float64 {
				return 0
			},
			dist:    0,
			noError: false,
		},
		{
			desc:    "invalid distance func",
			s:       []int{1, 2},
			t:       []int{1, 3},
			df:      nil,
			dist:    0,
			noError: false,
		},
	}

	for _, test := range testCases {
		d, err := Distance(test.s, test.t, test.df)
		if test.noError {
			assert.NoError(t, err)
			assert.InDelta(t, test.dist, d, 0.00001)
			continue
		}
		assert.Error(t, err)
	}
}
