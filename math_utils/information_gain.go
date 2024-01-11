package mathutils

func CalcInformationGain(data []DataPoint, attributeIndex int) float32 {
	var sameClass bool
	// Guard clause: Check if data is empty
	if len(data) == 0 {
		return 0.0
	}

	// Guard clause: Check if attributeIndex is valid
	if attributeIndex < 0 || attributeIndex >= len(data[0].Attributes) {
		return 0.0
	}

	// Guard clause: Check if all elements are of the same class
	sameClass = true
	for i := 1; i < len(data); i++ {
		if data[i].Class != data[0].Class {
			sameClass = false
			break
		}
	}

	if sameClass {
		return 0.0
	}

	var totalEntropy, weightedEntropy float32
	var p float64

	splitData := make(map[float32][]DataPoint)

	totalEntropy = CalcEntropy(data)

	for _, point := range data {
		splitData[point.Attributes[attributeIndex]] = append(splitData[point.Attributes[attributeIndex]], point)
	}

	for _, split := range splitData {
		p = float64(len(split)) / float64(len(data))
		weightedEntropy += float32(p) * CalcEntropy(split)
	}

	informationGain := totalEntropy - weightedEntropy

	if informationGain < 0 {
		return 0.0
	}

	return informationGain
}
