package Math

func Sum(a, b float64) float64 {
	return a + b
}

func Dif(a, b float64) float64 {
	return a - b
}

func Mult(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) float64 {
	return float64(a/b)
}

func Mod(a, b int) int {
	return a % b
}

func Abs(a float64) float64 {
	if (a > 0) {
		return a
	}else {
		return -1 * a
	}
}

func Max(a, b float64) float64 {
	if (a > b) {
		return a
	}else {
		return b
	}
}

func Min(a, b float64) float64 {
	if (a < b) {
		return a
	}else {
		return b
	}
}

func Acos(a float64) float64 {
	return (-0.69813170079773212 * a * a - 0.87266462599716477) * a + 1.5707963267948966
}

func Square(a float64) float64 {
	return a * a
}

func SquareRoot(a float64) float64 {
	var  sr float64 = a / 2
  	var temp float64;
  	for{
    	temp = sr
		sr = (temp + (a / temp)) / 2
    	if(temp - sr) == 0{
      		break;
    	}
  	}
  	return sr
}

func Pow(a float64, b int) float64 {
	res := 1.0
	for i := 0; i < b; i++ {
		res = res * a
	}
	return res
}

func Ceil(a float64) int {
	return int(a) + 1
}

func Floor(a float64) int {
	return int(a)
}