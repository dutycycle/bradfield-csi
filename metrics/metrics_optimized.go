package metrics

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"math"
)

type UserArrays struct {
	ages []uint8
	payments []uint64
}

func AverageAgeOptimized(users UserArrays) float64 {
	sum := uint64(0)
	for _, age := range users.ages {
		sum += uint64(age)
	}
	return float64(sum) / float64(len(users.ages))
}

func AveragePaymentAmountOptimized(users UserArrays) float64 {
	sum := uint64(0)
	for _, payment := range users.payments {
		sum += payment;
	}
	return float64(sum) / 100.0 / float64(len(users.payments))
}

// Sum of squares method, taken from https://play.golang.org/p/xQXiHFzmxxN
func StdDevPaymentAmountOptimized(users UserArrays) float64 {
	sumSquares, sum := float64(0), float64(0)
	numPayments := float64(len(users.payments))

	for _, payment := range users.payments {
		x := float64(payment) / 100.0
		sumSquares += x * x
		sum += x
	}
	avgSquares := sumSquares / numPayments
	avg := sum / numPayments
	return math.Sqrt(avgSquares - avg*avg)
}

func LoadDataOptimized() UserArrays {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	ages := make([]uint8, len(userLines))
	for i, line := range userLines {
		age, _ := strconv.Atoi(line[2])
		ages[i] = uint8(age)
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	payments := make([]uint64, len(paymentLines))
	for i, line := range paymentLines {
		payment, _ := strconv.ParseUint(line[0], 10, 64)
		payments[i] = uint64(payment)
	}

	return UserArrays{ages, payments}
}

func CalculateMetricsOptimized() Metrics {
	userArrays := LoadDataOptimized()

	return Metrics{AverageAgeOptimized(userArrays), AveragePaymentAmountOptimized(userArrays), StdDevPaymentAmountOptimized(userArrays) }
}