package expenses

// Allowe calculte the average
func Average(expns ...float32) float32 {
	return Sum(expns...) / float32(len(expns))
}

// Allowe sum
func Sum(expns ...float32) float32 {
	var sum float32

	for _, expn := range expns {
		sum += expn
	}

	return sum
}

// Allowe calculte the max
func Max(expns ...float32) float32 {
	var max float32

	for _, expn := range expns {
		if expn > max {
			max = expn
		}
	}

	return max
}

// Allowe calculte the min
func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}

	//var min float32 = expns[0]
	min := expns[0]

	for _, expn := range expns {
		if expn < min {
			min = expn
		}
	}
	return min
}
