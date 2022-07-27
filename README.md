实现在go上的scimark及其他测试：
	1、测试内容主要包括fft/monte/lu/sor/sparse,fft支持根据corenum多线程测试
	2、可以定时检测系统环境，cpu占用率等和log保存
未来实现基于mysql的事务测试等。
cd test/
go run main.go
choose which test (such as getinfo/lscpu/bench(fft/montcarlo/lu/sor/sparse) etc)
