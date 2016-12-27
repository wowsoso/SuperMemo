package SuperMemo

import (
	"math"
	"log"
)

func CalcInterval(try, factor float64) float64 {
    var interval float64

    if try < 1 {
	return 0
    } else if try == 1 {
	interval = 1
    } else if try == 2 {
	interval = 6
    } else {
	interval = CalcInterval(try-1, factor) * factor
    }

    return math.Ceil(interval)
}

func CalcNewFactor(oldFactor, quality float64) float64 {

	if quality > 5 || quality < 0 {
		log.Fatal("QUALITY MUST BE BETWEEN 0 and 5")
	}
	var newFactor float64 = oldFactor+(0.1-(5-quality)*(0.08+(5-quality)*0.02))
	if newFactor > 1.3 {
		if newFactor < 2.5 {
			return newFactor
		} else {
			return 2.5
		}		
	} else {
		return 1.3
	}
}