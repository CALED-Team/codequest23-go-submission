package main

// ParsePosition parses the position array and returns a 2D slice of float64 values.
func ParsePosition(positionArray []interface{}) [][]float64 {
	position := make([][]float64, len(positionArray))
	for i, singlePosition := range positionArray {
		floatPosition := make([]float64, 2)
		floatPosition[0] = singlePosition.([]interface{})[0].(float64)
		floatPosition[1] = singlePosition.([]interface{})[1].(float64)
		position[i] = floatPosition
	}
	return position
}
