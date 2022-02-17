//provide golang-based tests
//Tommy Jiang in 2022.2.16
package modtest

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

//FFT
func fft(data []complex128) {
	if len(data) == 1 {
		return
	}
	hl := len(data) / 2
	even := make([]complex128, hl)
	odd := make([]complex128, hl)
	for i := 0; i < hl; i++ {
		even[i] = data[i] + data[i+hl]
		odd[i] = (data[i] - data[i+hl]) *
			cmplx.Exp(complex(0, 2*float64(i)*math.Pi/float64(len(data))))
	}
	Fft(even)
	Fft(odd)
	for i := 0; i < hl; i++ {
		data[2*i] = even[i]
		data[2*i+1] = odd[i]
	}
}

func Fft(data []complex128) {
	fft(data)
}

func testfft(data []complex128) {
	tmp := make([]complex128, len(data))
	for k := 0; k < len(data); k++ {
		tmp[k] = 0
		for i := 0; i < len(data); i++ {
			tmp[k] += data[i] *
				cmplx.Exp(complex(0, 2*float64(i*k)*math.Pi/float64(len(data))))
		}
	}
	for i, v := range tmp {
		data[i] = v
	}
}

func InvFft() {
	ND := 6192
	fmt.Printf("====start fft====\n")
	t1 := time.Now().UnixNano()
	data1 := make([]complex128, ND)
	data2 := make([]complex128, ND)
	for i := range data1 {
		r := complex(rand.Float64()*2-1, rand.Float64()*2-1)
		data1[i] = r
		data2[i] = r
	}
	Fft(data1)
	testfft(data2)
	for i := range data1 {
		diff := cmplx.Abs(data2[i] - data1[i])
		if diff > 1e10 {
			fmt.Printf("Resulted value differs from normal fourir transform expected=%v, got=%v", data2[i], data1[i])
		}
	}
	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t1)
	fmt.Printf("==== fft end ====\n")
}

//Monte_carlo
func integrate(cycles int) float64 {
	var Seed int64 = 113
	rand.Seed(Seed)
	j := 0
	for i := 0; i < cycles; i++ {
		x := rand.Float64()
		y := rand.Float64()
		//fmt.Println(x*x + y*y)
		if x*x+y*y <= 1.0 {
			j++
		}
	}
	//fmt.Println(j)
	return float64(j) / float64(cycles) * 4.0
}

func Monte_carlo() {
	fmt.Printf("====start monte_carlo====")
	cycles := 16777216
	t1 := time.Now().UnixNano()
	x := integrate(cycles)
	t2 := time.Now().UnixNano()
	fmt.Println(x, t2-t1)
	fmt.Printf("==== monte_carlo end ====")
}

//Lu
func Lu() {

}

func Sor() {

}

func Sparse() {

}
