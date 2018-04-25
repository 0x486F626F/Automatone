package main

import (
	"fmt"
	"time"
)

var notes = [...]float64{0.1, 73, 82, 92, 98, 110, 124, 130, 146, 165, 186}

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

func playNote(speed, duration float64) {
	setSpeed(speed)
	runMotor(speed * duration / 58.0)
	setSpeed(1)
	runMotor(0.0001)
}

func initServ() {
	go h.run()
	go sh.run()
	go spHandlerOpen("/dev/ttyUSB0", 115200, "tinyg", false)
	time.Sleep(1 * time.Second)
	setAcc(200)
}

func main() {
	initServ()

	for i := 1; i < len(notes); i++ {
		playNote(notes[i], 0.1)
	}

	ch := make(chan bool)
	<-ch
}
