package mathutils

import "math"

func CalcEntropy(data []DataPoint) float32 {
	if len(data) == 0 {
		return 0.0
	}

	var entropy float32
	var p float64

	classCounts := make(map[int]int)
	entropy = 0.0

	for _, point := range data {
		classCounts[point.Class]++
	}

	for _, count := range classCounts {
		p = float64(count) / float64(len(data))
		if p > 0 {
			entropy -= float32(p * math.Log2(p))
		}
	}

	return entropy
}
