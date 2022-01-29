package frechet

import (
	"errors"
	"math"
	"reflect"
)

// DistanceFunc is the function used to calculate the distance between x and y
type DistanceFunc func(x, y interface{}) float64

// Distance calculates the frechet distance between series s and t.
func Distance(s, t interface{}, df DistanceFunc) (float64, error) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return 0, errors.New("series s is not a slice")
	}
	if reflect.TypeOf(t).Kind() != reflect.Slice {
		return 0, errors.New("series t is not a slice")
	}
	if df == nil {
		return 0, errors.New("invalid distance func")
	}

	ss := reflect.ValueOf(s)
	tt := reflect.ValueOf(t)
	slen, tlen := ss.Len(), tt.Len()
	if slen == 0 {
		return 0, errors.New("s series is empty")
	}
	if tlen == 0 {
		return 0, errors.New("t series is empty")
	}

	maxs := math.Inf(-1)
	for i := 0; i < slen; i++ {
		min := math.Inf(1)
		for j := 0; j < tlen; j++ {
			d := df(ss.Index(i).Interface(), tt.Index(j).Interface())
			min = math.Min(min, d)
		}
		maxs = math.Max(min, maxs)
	}

	maxt := math.Inf(-1)
	for i := 0; i < tlen; i++ {
		min := math.Inf(1)
		for j := 0; j < slen; j++ {
			d := df(tt.Index(i).Interface(), ss.Index(j).Interface())
			min = math.Min(min, d)
		}
		maxt = math.Max(min, maxt)
	}
	return math.Max(maxs, maxt), nil
}
