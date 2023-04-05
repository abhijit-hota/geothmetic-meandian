package main // import "abhijithota.me/gmdn"

import (
	"log"
	"math"
)

func arithmeticMean(nums []float64) float64 {
	var mean float64
	for _, num := range nums {
		mean += num
	}
	mean /= float64(len(nums))

	return mean
}

func geometricMean(nums []float64) float64 {
	var mean float64 = 1

	for _, num := range nums {
		mean *= num
	}
	mean = math.Pow(mean, 1/float64(len(nums)))

	return mean
}

func median(nums []float64) float64 {
	var median float64

	n := len(nums)
	if n%2 == 0 {
		median = (nums[n/2-1] + nums[n/2]) / 2.0
	} else {
		median = nums[n/2]
	}

	return median
}

func fmdn(nums []float64) [3]float64 {
	var fmdn [3]float64

	fmdn[0] = arithmeticMean(nums)
	fmdn[1] = geometricMean(nums)
	fmdn[2] = median(nums)

	return fmdn
}

const epsilon = 0.00000001

func checkConvergence(nums [3]float64) bool {

	if math.Abs(nums[0]-nums[1]) < epsilon &&
		math.Abs(nums[1]-nums[2]) < epsilon {
		return true
	}

	return false
}

const maxCount = 1000000

func GeothmeticMeandian(nums []float64) float64 {
	isConverged := false

	arr := fmdn(nums)
	if *debug {
		log.Printf("[debug]: (#%d) %v", 0, arr)
	}

	for i := 0; (i < maxCount) && !isConverged; i++ {
		arr = fmdn(arr[:])
		isConverged = checkConvergence(arr)
		if *debug {
			log.Printf("[debug]: (#%d) %v", i+1, arr)
		}
	}

	return arr[0]
}
