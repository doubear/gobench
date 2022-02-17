package main

import (
	"fmt"
	"modtest"
	"os"
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
2 provide fft/mont_carlo test
3 db saving support (future)
4 monitorlog saving support
5 web access support (future)
========================================
args such as getinfo/lscpu/montcarlo etc
========================================
  ／ ￣ \
＜  ･    \   …？
 |   ３ ） \   
＜  ･        \
  \＿＿／u  _u  )
            Ｕ Ｕ
doubear in 2022 

input:getinfo/lscpu/montcarlo etc`

var help string = `========================================
args such as getinfo/lscpu/montcarlo etc
========================================
input:getinfo/lscpu/montcarlo etc`

func main() {
	fmt.Println("hello zx")
	//r := routers.Addrouter()
	//db.Dbinit()
	//db.Dbclose()
	//r.Run()
	//modtest.InvFft()
	//modtest.Monte_carlo()
	//modtest.Lscpu()

	//modtest.Getinfo(2)

	fmt.Println(reminder)
	for {
		cmd := getinput()
		//fmt.Println(cmd)
		switch cmd {
		case "getinfo":
			modtest.Getinfo(5)
			fmt.Printf("continue input:\n")

		case "help":
			fmt.Println(help)
		case "exit":
			fmt.Printf("loading stop...:Exit\n")
			time.Sleep(time.Second)
			os.Exit(1)
		default:
			fmt.Printf("error cmd stop exec:Exit\n")
			time.Sleep(time.Second)
			os.Exit(1)
		}
		//modtest.Getinfo(5)
		//exec

	}

}
