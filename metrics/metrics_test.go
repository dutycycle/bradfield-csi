package metrics

import (
	"math"
	"testing"
)

func BenchmarkMetrics(b *testing.B) {
	b.Run("Metrics Baseline", func(b *testing.B) {
		var metrics Metrics

		for n := 0; n < b.N; n++ {
			metrics = CalculateMetrics()
		}

		expectedAverageAge := 59.62
		if math.IsNaN(metrics.averageAge) || math.Abs(metrics.averageAge - expectedAverageAge) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expectedAverageAge, metrics.averageAge)
		}

		expectedAveragePaymentAmount := 499850.559
		if math.IsNaN(metrics.averagePaymentAmount) || math.Abs(metrics.averagePaymentAmount-expectedAveragePaymentAmount) > 0.01 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expectedAveragePaymentAmount, metrics.averagePaymentAmount)
		}

		expectedStdDevAmount := 288684.850
		if math.IsNaN(metrics.stdDevPaymentAmount) || math.Abs(metrics.stdDevPaymentAmount-expectedStdDevAmount) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expectedStdDevAmount, metrics.stdDevPaymentAmount)
		}
	})

	b.Run("Metrics Array", func(b *testing.B) {
		var metrics Metrics

		for n := 0; n < b.N; n++ {
			metrics = CalculateMetricsArray()
		}

		expectedAverageAge := 59.62
		if math.IsNaN(metrics.averageAge) || math.Abs(metrics.averageAge - expectedAverageAge) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expectedAverageAge, metrics.averageAge)
		}

		expectedAveragePaymentAmount := 499850.559
		if math.IsNaN(metrics.averagePaymentAmount) || math.Abs(metrics.averagePaymentAmount-expectedAveragePaymentAmount) > 0.01 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expectedAveragePaymentAmount, metrics.averagePaymentAmount)
		}

		expectedStdDevAmount := 288684.850
		if math.IsNaN(metrics.stdDevPaymentAmount) || math.Abs(metrics.stdDevPaymentAmount-expectedStdDevAmount) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expectedStdDevAmount, metrics.stdDevPaymentAmount)
		}
	})


	b.Run("Metrics Streaming", func(b *testing.B) {
		var metrics Metrics

		for n := 0; n < b.N; n++ {
			metrics = CalculateMetricStreaming()
		}

		expectedAverageAge := 59.62
		if math.IsNaN(metrics.averageAge) || math.Abs(metrics.averageAge - expectedAverageAge) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expectedAverageAge, metrics.averageAge)
		}

		expectedAveragePaymentAmount := 499850.559
		if math.IsNaN(metrics.averagePaymentAmount) || math.Abs(metrics.averagePaymentAmount-expectedAveragePaymentAmount) > 0.01 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expectedAveragePaymentAmount, metrics.averagePaymentAmount)
		}

		expectedStdDevAmount := 288684.850
		if math.IsNaN(metrics.stdDevPaymentAmount) || math.Abs(metrics.stdDevPaymentAmount-expectedStdDevAmount) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expectedStdDevAmount, metrics.stdDevPaymentAmount)
		}
	})
}
