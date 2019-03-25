package utils

import "math"

type rect interface {
	getRect() (float64, float64, float64, float64)
	getRoundedRectSlice() [4]int
}

type ptRect [4]float64
type termRect [4]float64

func (r *termRect) getRect() (float64, float64, float64, float64) {
	return r[0], r[1], r[2], r[3]
}

func (r *termRect) getRoundedRectSlice() [4]int {
	return roundRectSlice([4]float64(*r))
}

func (r *ptRect) getRect() (float64, float64, float64, float64) {
	return r[0], r[1], r[2], r[3]
}

func (r *ptRect) getRoundedRectSlice() [4]int {
	return roundRectSlice([4]float64(*r))
}

func roundRectSlice(input [4]float64) [4]int {
	var result [4]int
	for i, f := range input {
		result[i] = int(math.Round(f))
	}
	return result
}
