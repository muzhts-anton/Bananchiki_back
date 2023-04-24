package points

import (
	"banana/pkg/domain"
)

const (
	interval0 = 0
	interval1 = 0.5
	interval2 = 0.75
	interval3 = 1
)

var pricepool = map[float64]float64{
	interval0: 0,
	interval1: 0.5,
	interval2: 1.5,
}

const err = 0.001

func CalculateQuizPoints(mode bool, price uint64, timeTotal uint64, timeSpent float64) (uint64, error) {
	if !mode {
		return price, nil
	}

	t := timeSpent / float64(timeTotal)
	if interval0-err < t && t < interval1 {
		return price + uint64(float64(price)*pricepool[interval0]), nil
	}
	if interval1 < t && t < interval2 {
		return price + uint64(float64(price)*pricepool[interval1]), nil
	}
	if interval2 < t && t < interval3+err {
		return price + uint64(float64(price)*pricepool[interval2]), nil
	}

	return 0, domain.ErrUnexpectedTimeValue
}
