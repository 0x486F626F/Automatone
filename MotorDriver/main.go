package main

import (
	"fmt"
	"time"
)

func setAcc(acc float64) {
	cmd := "sendjson {\"P\":\"/dev/ttyUSB0\",\"Data\":"
	cmd += fmt.Sprintf("[{\"D\":\"{\\\"xjm\\\": %f}\\n\",\"Id\":\"console1\"}]}", acc)
	fmt.Println(cmd)
	spWriteJson(cmd)
}
func setSpeed(speed float64) {
	cmd := "sendjson {\"P\":\"/dev/ttyUSB0\",\"Data\":"
	cmd += fmt.Sprintf("[{\"D\":\"{\\\"xvm\\\": %f}\\n\",\"Id\":\"console1\"}]}", speed)
	fmt.Println(cmd)
	spWriteJson(cmd)
}
func runMotor(dis float64) {
	cmd := "sendjson {\"P\":\"/dev/ttyUSB0\",\"Data\":"
	cmd += fmt.Sprintf("[{\"D\":\"G91 G0 X-%f\\n\",\"Id\":\"console1\"}]}", dis)
	fmt.Println(cmd)
	spWriteJson(cmd)
}

func note(a int, b float64) {
	var freq = [...]float64{0.1, 98.0, 110.0, 124.0, 130.0, 146.0, 165.0, 186.0,
		92.0, 82.0, 73.0}
	setSpeed(freq[a])
	runMotor(freq[a] / 40.0 / b)
	setSpeed(1)
	runMotor(0.0005)
}

func play(n int, a []int, len []float64) {
	for i := 0; i < n; i++ {
		note(a[i], len[i])
	}
}

func loadScript() {
	setAcc(200)

	var a []int
	var b []float64
	a = append(a, 0)
	b = append(b, 4)

	a = append(a, 1, 2)
	b = append(b, 4, 4)
	a = append(a, 1, 2, 1)
	b = append(b, (4 / 1.5), 8, 2)
	a = append(a, 0, 10, 4, 3, 2, 1)
	b = append(b, (4 / 1.5), 8, (4 / 1.5), (8 / 3), 8, 2)
	a = append(a, 0, 1, 2)
	b = append(b, 1, 4, 4)
	a = append(a, 3, 2, 3, 4, 3, 1, 1, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 3, 3, 3, 4, 3, 2, 1)
	b = append(b, 8, 8, 8, 8, 8, 4, (8 / 5))

	a = append(a, 0, 1, 10, 10)
	b = append(b, (4 / 7.0), (4 / 1.5), 8, 2)
	a = append(a, 0, 10, 4, 3, 2, 1)
	b = append(b, (4 / 1.5), 8, (4 / 1.5), (8 / 3), 8, 2)
	a = append(a, 0, 1, 2)
	b = append(b, 1, 4, 4)
	a = append(a, 3, 2, 3, 4, 3, 2, 1, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 3, 3, 3, 4, 3, 2, 1)
	b = append(b, 8, 8, 8, 8, 4, 8, (8 / 5))
	a = append(a, 0, 1, 2)
	b = append(b, 1, 4, 4)
	a = append(a, 3, 2, 3, 4, 3, 2, 1, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 3, 3, 3, 4, 3, 2, 1, 0)
	b = append(b, 8, 8, 8, 8, 4, 8, (8 / 7), 0.8)

	a = append(a, 1, 2, 3, 1, 2, 3, 1, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 3, 1, 2, 3, 1, 2, 3, 3)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 5, 2, 2, 0)
	b = append(b, (4 / 3), 8, (8 / 5), 2)
	a = append(a, 1, 2, 3, 1, 2, 3, 1, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 3, 1, 2, 3, 1, 2, 3, 1)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 5, 2, 2, 0)
	b = append(b, (4 / 3), 8, (8 / 5), 4)

	a = append(a, 8, 9, 8, 9, 8, 9, 1, 9, 9, 0)
	b = append(b, 8, 8, 4, 8, 4, 8, 2, 4, 8, 4)
	a = append(a, 10, 2, 2, 2, 3, 4, 3, 0)
	b = append(b, 8, 4, 8, 4, 8, 4, 4, (8 / 5))
	a = append(a, 9, 8, 9, 8, 9, 1, 9, 10, 0)
	b = append(b, 8, 4, 8, 4, 8, 2, 4, 4, 4)
	a = append(a, 3, 2, 2, 2, 1, 1, 2, 2, 0)
	b = append(b, (4 / 3), 8, 8, (8 / 7), 8, (8 / 7), 8, 2, 8)

	a = append(a, 1, 1, 1)
	b = append(b, 8, 8, 8)
	a = append(a, 5, 5, 5, 3, 2)
	b = append(b, 4, 4, 4, 8, 8)
	a = append(a, 2, 2, 0, 1, 1, 1, 1)
	b = append(b, 8, 8, 8, 4, 8, 8, 8)
	a = append(a, 5, 5, 0, 5, 5, 3, 3, 2)
	b = append(b, 8, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 0, 0, 1, 1, 2, 3, 2)
	b = append(b, 4, 8, 8, 8, 8, 8, 8)
	a = append(a, 0, 1, 1, 1, 1, 2, 3)
	b = append(b, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 2, 1, 1, 9, 1, 9, 1, 0)
	b = append(b, 4, 8, 8, 8, 8, 8, 8, 8)
	a = append(a, 5, 3, 5, 3, 2)
	b = append(b, 4, 8, 4, 4, (4 / 1.5))

	a = append(a, 1, 1, 1)
	b = append(b, 8, 8, 8)
	a = append(a, 5, 5, 5, 3, 2)
	b = append(b, 4, 4, 4, 8, 8)
	a = append(a, 2, 2, 0, 1, 1, 1, 1)
	b = append(b, 8, 8, 8, 4, 8, 8, 8)
	a = append(a, 5, 5, 5, 0, 3, 0, 2)
	b = append(b, 4, 8, 8, 8, 8, 8, 8)
	a = append(a, 0, 0, 1, 1, 2, 3, 2)
	b = append(b, 4, 8, 8, 8, 8, 8, 4)
	a = append(a, 1, 1, 1, 1, 2, 3)
	b = append(b, 8, 8, 8, 8, 8, 8)
	a = append(a, 2, 1, 1, 9, 1, 2, 3, 2)
	b = append(b, 4, 8, 8, 8, 8, 8, 8, 4)
	a = append(a, 1, 1)
	b = append(b, 8, 2)
	fmt.Println(len(a), len(b))
	play(len(a), a, b)
}

func main() {
	go h.run()
	go sh.run()

	go spHandlerOpen("/dev/ttyUSB0", 115200, "tinyg", false)

	time.Sleep(1 * time.Second)
	loadScript()

	ch := make(chan bool)
	<-ch
}
