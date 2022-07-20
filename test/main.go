package main

import (
	"fmt"
	"modtest"
	"os"
	"strconv"
	"time"
)

func getinput() string {
	input := ""
	fmt.Scanln(&input)
	return input
}
func sexit() {
	c := make(chan os.Signal)
	select {
	case signal := <-c:
		fmt.Printf("got %s signal\n", signal)
		time.Sleep(time.Second * 3)
		os.Exit(1)
	}
}

var reminder string = `1 provide mem/cpu/disk monitor 
2 provide fft/mont_carlo/lu/sor/sparse test
3 db saving support
4 monitorlog saving support
5 web access support
========================================
args such as getinfo/lscpu/benchtest(fft/montcarlo) etc
========================================
  ／ ￣ \
＜  ･    \   …？
 |   ３ ） \   
＜  ･        \
  \＿＿／u  _u  )
            Ｕ Ｕ
doubear in 2022 version 0.1.5

input:getinfo/lscpu/bench(fft/montcarlo/lu/sor/sparse) etc`

var help string = `========================================
args such as getinfo/lscpu/bench(fft/montcarlo/lu/sor/sparse) etc
========================================
input:getinfo/lscpu/bench etc`

func benchchoose(args string) {
	fmt.Printf("input test:fft/montcarlo/lu/sor/sparse\n")
	switch args {
	case "fft":
		//modtest.InvFft()
		modtest.Multi_fft()
		fmt.Printf("continue input test:\n")
	case "montcarlo":
		//modtest.Monte_carlo()
		modtest.Mult_mont()
		fmt.Printf("continue input test:\n")
	case "lu":
		modtest.Lu()
		fmt.Printf("continue input test:\n")
	case "sor":
		modtest.Sor()
		fmt.Printf("continue input test:\n")
	case "sparse":
		modtest.Sparse()
		fmt.Printf("continue input test:\n")
	case "all":
		//	modtest.InvFft()
		//	modtest.Monte_carlo()
		modtest.Multi_fft()
		modtest.Mult_mont()
		modtest.Lu()
		modtest.Sor()
		modtest.Sparse()
		fmt.Printf("continue input test:\n")

	case "back":
		choose()
	case "help":
		fmt.Printf("fft/montcarlo/all/exit cmd\n")
	case "exit":
		fmt.Printf("loading stop...:Exit\n")
		time.Sleep(time.Second)
		os.Exit(1)
	default:
		fmt.Printf("input reset:\n")
		benchchoose(getinput())
	}
}

func choose() {
	for {
		cmd := getinput()
		//fmt.Println(cmd)
		switch cmd {
		case "getinfo":
			fmt.Printf("input monitor seconds\n")
			ns, _ := strconv.Atoi(getinput())
			modtest.Getinfo(ns)
			fmt.Printf("continue input:\n")

		case "bench":
			fmt.Printf("input u want to test all or single test\n")
			benchchoose(getinput())
			fmt.Printf("test end pls continue input getinfo/bench etc:\n")

		case "lscpu":
			modtest.Lscpu()
			fmt.Printf("continue input:\n")

		case "help":
			fmt.Println(help)
		case "exit":
			fmt.Printf("loading stop...:Exit\n")
			time.Sleep(time.Second)
			os.Exit(1)
		default:
			fmt.Printf("error cmd restart inpout:\n")
			choose()
		}
	}
}

func main() {
	fmt.Println("hello doubear")
	//r := routers.Addrouter()
	//db.Dbinit()
	//db.Dbclose()
	//r.Run()
	//modtest.InvFft()
	//modtest.Monte_carlo()
	//modtest.Lscpu()

	//modtest.Getinfo(2)

	fmt.Println(reminder)
	choose()

}
