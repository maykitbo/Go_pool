package anomaly

import (
	"fmt"
	"math"
)

type AnomalyDetector struct {
	sD           float64
	mean         float64
	len          int64
	k            float64
	anomalyCount int
	Do           func(float64) bool
}

func Init(kinit float64) *AnomalyDetector {
	var ad AnomalyDetector
	ad.Do = ad.firstValue
	ad.k = kinit
	return &ad
}

func (ad *AnomalyDetector) firstValue(value float64) bool {
	ad.len = 1
	ad.mean = float64(value)
	ad.sD = 0
	ad.Do = ad.notEnoughValues
	return false
}

func (ad *AnomalyDetector) notEnoughValues(value float64) bool {
	ad.statisticCount(value)
	if ad.len > 100 {
		ad.Do = ad.enoughValues
	}
	return false
}

func (ad *AnomalyDetector) enoughValues(value float64) bool {
	sigma := ad.sD * ad.k
	if math.Abs(value-ad.mean) > sigma {
		fmt.Println("!!! ANOMALY !!! ", ad.anomalyCount, value, ad.len, ad.mean, ad.sD)
		ad.anomalyCount++
		return true
	}
	ad.statisticCount(value)
	return false
}

func (ad *AnomalyDetector) statisticCount(value float64) {
	ad.len++
	ad.sD = math.Sqrt((float64(ad.len-1) / float64(ad.len)) * (ad.sD*ad.sD + (math.Pow((ad.mean-value), 2) / float64(ad.len))))
	ad.mean = (ad.mean * float64(ad.len-1) / float64(ad.len)) + value/float64(ad.len)
}
