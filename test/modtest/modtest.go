//provide golang-based tests
//Tommy Jiang in 2022.2.16
package modtest

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
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
	ND := 4096 //size 64KB
	//fmt.Printf("====start fft====\n")
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
	//fmt.Println(t2 - t1)
	fmt.Printf("====start fft====\n %d \n==== fft end ====\n", t2-t1)
}
func Multi_fft() { //the number of threads is the same with cores number
	n, _ := cpu.Counts(true)
	fmt.Printf("we will run %d threads\n", n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			InvFft()
			wg.Done()
		}()
	}
	wg.Wait()
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
	//fmt.Printf("====start monte_carlo====\n")
	cycles := 16777216
	t1 := time.Now().UnixNano()
	x := integrate(cycles)
	t2 := time.Now().UnixNano()
	//fmt.Println(x, t2-t1)
	fmt.Printf("====start monte_carlo====\n %f %d\n==== monte_carlo end ====\n", x, t2-t1)

}
func Mult_mont() { //support multi threads running
	n, _ := cpu.Counts(true)
	var wg sync.WaitGroup
	wg.Add(n)
	fmt.Printf("we will run %d threads\n", n)
	for i := 0; i < n; i++ {
		go func() {
			Monte_carlo()
			wg.Done()
		}()
	}

	wg.Wait()
}

//Lu
var R = rand.New(rand.NewSource(time.Now().Unix()))

func randR() (v float64) {
	return 0 + R.Float64()*1
}
func randommatrix(N int) [][]float64 {
	A := make([][]float64, N)
	for i := range A {
		A[i] = make([]float64, N)
	}
	for k := 0; k < N; k++ {
		for m := 0; m < N; m++ {
			A[k][m] = randR()
		}
	}
	return A
}

func copymatrix(lu, A [][]float64) {
	M := len(A)
	N := len(A)
	remainder := N * 3
	for i := 0; i < M; i++ {
		Bi := lu[i]
		Ai := A[i]
		for j := 0; j < N; j++ {
			Bi[j] = Ai[j]
		}
		for j := remainder; j < N; j += 4 {
			Bi[j] = Ai[j]
			Bi[j+1] = Ai[j+1]
			Bi[j+2] = Ai[j+2]
			Bi[j+3] = Ai[j+3]
		}
	}
}
func factor(A [][]float64, pivot []int) {
	M := len(A)
	N := len(A)
	min := func(M, N int) int {
		if M > N {
			return N
		}
		return M
	}(M, N)
	for j := 0; j < min; j++ {
		jp := j
		t := math.Abs(A[j][j])
		for i := 0; i < M; i++ {
			ab := math.Abs(A[i][j])
			if ab > t {
				jp = i
				t = ab
			}
		}
		pivot[j] = jp
		if jp != j {
			A[j], A[jp] = A[jp], A[j]
		}
		if j < M-1 {
			recp := 1.0 / A[j][j]
			for k := j + 1; k < M; k++ {
				A[k][j] *= recp
			}
		}
		if j < min-1 {
			for ii := j + 1; ii < M; ii++ {
				Aii := A[ii]
				Aj := A[j]
				AiAiJ := Aii[j]
				for jj := j + 1; jj < N; jj++ {
					Aii[jj] -= AiAiJ * Aj[jj]
				}
			}
		}
	}
}

func measureLU(N int) {
	A := randommatrix(N)
	lu := make([][]float64, N)
	pivot := make([]int, N)
	for i := range lu {
		lu[i] = make([]float64, N)
	}
	copymatrix(lu, A)
	t1 := time.Now().UnixNano()
	factor(lu, pivot)
	t2 := time.Now().UnixNano()
	fmt.Printf("====start lu====\n %d \n==== lu end ====\n", t2-t1)
}

func Lu() {
	N := 1000
	measureLU(N)
}

//sor
func measureSOR(omega float64, G [][]float64, num_iterations int) {
	M := len(G)
	N := len(G)
	omega_over_four := omega * 0.25
	one_minus_omega := 1.0 - omega
	Gi := make([]float64, N)
	Gi_sum := 0.0
	Mm1 := M - 1
	Nm1 := N - 1
	for p := 0; p < num_iterations; p++ {
		for i := 1; i < Mm1; i++ {
			Gim1 := G[i-1]
			Gip1 := G[i-1]
			for j := 1; j < Nm1; j++ {
				Gi[j] = omega_over_four*(Gim1[j]+Gip1[j]+Gi[j-1]+Gi[j+1]) + one_minus_omega*Gi[j]
			}
		}
	}
	for k := 0; k < len(Gi); k++ {
		Gi_sum += Gi[k]
	}
}
func Sor() {
	N := 1000
	G := make([][]float64, N)
	for i := range G {
		G[i] = make([]float64, N)
	}
	G = randommatrix(N)
	cycles := 256
	t1 := time.Now().UnixNano()
	measureSOR(1.25, G, cycles)
	t2 := time.Now().UnixNano()
	fmt.Printf("====start sor====\n %d \n==== sor end ====\n", t2-t1)
}

//sparse
func matmult(y, val, x []float64, row, col []int, NUM_ITERATIONS, call_count int) {
	//total := 0.0
	M := len(row) - 1
	for reps := 0; reps < NUM_ITERATIONS; reps++ {
		for r := 0; r < M; r++ {
			sum := 0.0
			rowR := row[r]
			rowRp1 := row[r+1]
			for i := rowR; i < rowRp1; i++ {
				sum += x[col[i]] * val[i]
				y[r] = sum
			}
		}
	}
}

func Sparse() {
	//N := 1000
	nz := 1000000
	N := 100000
	x := make([]float64, N)
	y := make([]float64, N)
	for i := range x {
		x[i] = randR()
	}
	nr := nz / N
	anz := nr * N
	val := make([]float64, anz)
	for i := range val {
		val[i] = randR()
	}
	col := make([]int, anz)
	row := make([]int, N+1)
	row[0] = 0
	for r := 0; r < N; r++ {
		rowr := row[r]
		row[r+1] = rowr + nr
		step := r / nr
		if step < 1 {
			step = 1
		}
		for i := 0; i < nr; i++ {
			col[rowr+i] = i * step
		}
	}
	cycles := 512
	count := 2
	t1 := time.Now().UnixNano()
	matmult(y, val, x, row, col, cycles, count)
	t2 := time.Now().UnixNano()
	fmt.Printf("====start sparse====\n %d \n==== sparse end ====\n", t2-t1)
}
