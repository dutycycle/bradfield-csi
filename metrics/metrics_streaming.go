package metrics

import (
	"os"
	"log"
	"strconv"
	"math"
	"strings"
	"bufio"
)

func CalculateMetricStreaming() Metrics {
	averageAge := 0.0
	averagePaymentAmount := 0.0
	stdDevPaymentAmount := 0.0

	f, err := os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}

	numPayments, delta, M2 := 0.0, 0.0, 0.0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		paymentCents, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
		payment := float64(paymentCents) / 100.0

		numPayments++
		
		// Calculate StdDev using Welford's Online Algorithm
		// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance#Welford's_online_algorithm
		delta = payment - averagePaymentAmount
		averagePaymentAmount += delta / numPayments
		M2 += delta * (payment - averagePaymentAmount)
	}

	averagePaymentAmount = averagePaymentAmount
	stdDevPaymentAmount = math.Sqrt(M2 / numPayments)

	sumAges, numUsers := 0, 0
	f, err = os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		age, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[2])
		numUsers++
		sumAges += age
	}

	averageAge = float64(sumAges) / float64(numUsers);

	return Metrics{averageAge, averagePaymentAmount, stdDevPaymentAmount}
}

