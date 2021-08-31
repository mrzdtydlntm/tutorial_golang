package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus            Kubus   = Kubus{4}
	volumeSeharusnya float64 = 64
	luasSeharusnya   float64 = 96
	// kelilingSeharusnya float64 = 48
	kelilingSeharusnya float64 = 49 //uji fail
)

//unit test w/ testing package
/*
func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Salah! Harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("Salah! Harusnya %.2f", kelilingSeharusnya)
	}
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
*/

//unit test using testify
func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}

//unit test
//$ go test 48.\ unittest/main_test.go 48.\ unittest/main.go -v

//benchmark test
//$ go test 48.\ unittest/main.go 48.\ unittest/main_test.go -bench=.
