package hls

import (
	"math/big"
	"testing"
)

func compareBySignficantDigits(t *testing.T, f1, f2 float64, precision int, signficance int) bool {
	t.Helper()

	F1 := big.NewFloat(f1).SetPrec(uint(precision))
	F2 := big.NewFloat(f2).SetPrec(uint(precision))

	if F1.Sign() != F2.Sign() {
		return false
	}

	s1 := F1.Text('f', precision)
	s2 := F2.Text('f', precision)

	if s1[0] == '-' {
		s1 = s1[1:]
	}
	if s2[0] == '-' {
		s2 = s2[1:]
	}

	for (s1[0] == '.' && s2[0] == '.') || (s1[0] == '0' && s2[0] == '0') {
		s1 = s1[1:]
		s2 = s2[1:]
	}

	if len(s1) != len(s2) && (len(s1) < signficance || len(s2) < signficance) {
		return false
	}

	for i := 0; i < signficance; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Compare with the output of `tensorflow.signal.dct`
func TestDCT32_Tensorflow(t *testing.T) {
	got := DCT32([]float64{
		0.00000000e+00, 1.95090322e-01, 3.82683432e-01, 5.55570233e-01,
		7.07106781e-01, 8.31469612e-01, 9.23879533e-01, 9.80785280e-01,
		1.00000000e+00, 9.80785280e-01, 9.23879533e-01, 8.31469612e-01,
		7.07106781e-01, 5.55570233e-01, 3.82683432e-01, 1.95090322e-01,
		1.22464680e-16, -1.95090322e-01, -3.82683432e-01, -5.55570233e-01,
		-7.07106781e-01, -8.31469612e-01, -9.23879533e-01, -9.80785280e-01,
		-1.00000000e+00, -9.80785280e-01, -9.23879533e-01, -8.31469612e-01,
		-7.07106781e-01, -5.55570233e-01, -3.82683432e-01, -1.95090322e-01,
	})
	want := []float64{
		2.4492937e-16, 2.7064280e+01, -3.1365492e+00, -1.6186134e+01,
		2.4022312e-16, -3.8283632e+00, 1.4526381e-06, -1.7681264e+00,
		2.2628522e-16, -1.0182655e+00, -4.2100743e-07, -6.5700215e-01,
		2.0365132e-16, -4.5386967e-01, 6.1112178e-06, -3.2749316e-01,
		1.7319121e-16, -2.4288861e-01, -7.6461049e-07, -1.8286194e-01,
		1.3607546e-16, -1.3813220e-01, 3.6906343e-07, -1.0328477e-01,
		9.3730409e-17, -7.4950568e-02, -4.1193351e-07, -5.0897490e-02,
		4.7783329e-17, -2.9547997e-02, 3.8822355e-07, -9.6844761e-03,
	}

	for i := range got {
		if !compareBySignficantDigits(t, got[i], want[i], 50, 6) {
			t.Errorf(
				"DCT32 bin is incorrect.\ngot[%[1]d]:\n%.50[2]f,\nwant[%[1]d]:\n%.50[3]f\n",
				i, got[i], want[i],
			)
		}
	}
}

// Compare with the output of `scipy.fft.dct`
func TestDCT32_Scipy(t *testing.T) {
	got := DCT32([]float64{
		0.00000000e+00, 1.95090322e-01, 3.82683432e-01, 5.55570233e-01,
		7.07106781e-01, 8.31469612e-01, 9.23879533e-01, 9.80785280e-01,
		1.00000000e+00, 9.80785280e-01, 9.23879533e-01, 8.31469612e-01,
		7.07106781e-01, 5.55570233e-01, 3.82683432e-01, 1.95090322e-01,
		1.22464680e-16, -1.95090322e-01, -3.82683432e-01, -5.55570233e-01,
		-7.07106781e-01, -8.31469612e-01, -9.23879533e-01, -9.80785280e-01,
		-1.00000000e+00, -9.80785280e-01, -9.23879533e-01, -8.31469612e-01,
		-7.07106781e-01, -5.55570233e-01, -3.82683432e-01, -1.95090322e-01,
	})
	want := []float64{
		2.22044605e-16, 2.70642806e+01, -3.13654849e+00, -1.61861364e+01,
		2.17778080e-16, -3.82836379e+00, 9.47875655e-10, -1.76812694e+00,
		2.05142466e-16, -1.01826580e+00, 1.27065078e-09, -6.57001624e-01,
		1.84623342e-16, -4.53867859e-01, -1.80050469e-09, -3.27497986e-01,
		1.57009246e-16, -2.42889061e-01, 3.29392614e-09, -1.82861953e-01,
		1.23361373e-16, -1.38132090e-01, -3.16000953e-09, -1.03285003e-01,
		8.49727916e-17, -7.49505175e-02, -6.92284860e-09, -5.08969950e-02,
		4.33187535e-17, -2.95456654e-02, 6.69533184e-09, -9.68904199e-03,
	}

	for i := range got {
		if !compareBySignficantDigits(t, got[i], want[i], 50, 6) {
			t.Errorf(
				"DCT32 bin is incorrect.\ngot[%[1]d]:\n%.50[2]f,\nwant[%[1]d]:\n%.50[3]f\n",
				i, got[i], want[i],
			)
		}
	}
}

func TestDCT32ByFFTW_Scipy(t *testing.T) {
	got := DCT32ByFFTW([]float64{
		0.00000000e+00, 1.95090322e-01, 3.82683432e-01, 5.55570233e-01,
		7.07106781e-01, 8.31469612e-01, 9.23879533e-01, 9.80785280e-01,
		1.00000000e+00, 9.80785280e-01, 9.23879533e-01, 8.31469612e-01,
		7.07106781e-01, 5.55570233e-01, 3.82683432e-01, 1.95090322e-01,
		1.22464680e-16, -1.95090322e-01, -3.82683432e-01, -5.55570233e-01,
		-7.07106781e-01, -8.31469612e-01, -9.23879533e-01, -9.80785280e-01,
		-1.00000000e+00, -9.80785280e-01, -9.23879533e-01, -8.31469612e-01,
		-7.07106781e-01, -5.55570233e-01, -3.82683432e-01, -1.95090322e-01,
	})
	want := []float64{
		2.22044605e-16, 2.70642806e+01, -3.13654849e+00, -1.61861364e+01,
		2.17778080e-16, -3.82836379e+00, 9.47875655e-10, -1.76812694e+00,
		2.05142466e-16, -1.01826580e+00, 1.27065078e-09, -6.57001624e-01,
		1.84623342e-16, -4.53867859e-01, -1.80050469e-09, -3.27497986e-01,
		1.57009246e-16, -2.42889061e-01, 3.29392614e-09, -1.82861953e-01,
		1.23361373e-16, -1.38132090e-01, -3.16000953e-09, -1.03285003e-01,
		8.49727916e-17, -7.49505175e-02, -6.92284860e-09, -5.08969950e-02,
		4.33187535e-17, -2.95456654e-02, 6.69533184e-09, -9.68904199e-03,
	}

	for i := range got {
		if !compareBySignficantDigits(t, got[i], want[i], 50, 6) {
			t.Errorf(
				"DCT32 bin is incorrect.\ngot[%[1]d]:\n%.50[2]f,\nwant[%[1]d]:\n%.50[3]f\n",
				i, got[i], want[i],
			)
		}
	}
}
