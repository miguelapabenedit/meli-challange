package algorithm

import (
	"errors"
	"math"
)

/*Trilateration given the 2d coordinates of 3 know points, it determitates the location of the 4th unknow point
	base on their distances.
 [Required] d1,d2 & d3 must be positive values
*/
func Trilateration(x1, y1, d1, x2, y2, d2, x3, y3, d3 float64) (float64, float64, error) {
	if d1 < 0 || d2 < 0 || d3 < 0 {
		return 0, 0, errors.New("Distances can't be < 0")
	}

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := math.Pow(d1, 2) - math.Pow(d2, 2) - math.Pow(x1, 2) + math.Pow(x2, 2) - math.Pow(y1, 2) + math.Pow(y2, 2)
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := math.Pow(d2, 2) - math.Pow(d3, 2) - math.Pow(x1, 2) + math.Pow(x3, 2) - math.Pow(y1, 2) + math.Pow(x3, 2)
	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)

	return x, y, nil
}
