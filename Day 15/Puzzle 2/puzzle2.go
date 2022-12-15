package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	sensorX int
	sensorY int
	beaconX int
	beaconY int
}

func main() {
	file, err := os.Open("..\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var l []pair

	for scanner.Scan() {

		line := scanner.Text()
		//fmt.Println(line)

		split := strings.Split(line, ":")
		//fmt.Println(split)

		sensor := split[0][12:]
		sensorx, _ := strconv.Atoi(strings.Split(sensor, ",")[0])
		sensory, _ := strconv.Atoi(strings.Split(sensor, "=")[1])

		beacon := split[1][24:]
		beaconx, _ := strconv.Atoi(strings.Split(beacon, ",")[0])
		beacony, _ := strconv.Atoi(strings.Split(beacon, "=")[1])

		tmp := pair{sensorX: sensorx, sensorY: sensory, beaconX: beaconx, beaconY: beacony}
		l = append(l, tmp)
	}

	max := 4000000

loop:
	for _, element := range l {

		manhatan := Abs(element.sensorX-element.beaconX) + Abs(element.sensorY-element.beaconY)

		x := element.sensorX - manhatan - 1
		y := element.sensorY

		for x < element.sensorX {
			found := true
			for _, tmpelement := range l {

				tmpmanhatan := Abs(tmpelement.sensorX-tmpelement.beaconX) + Abs(tmpelement.sensorY-tmpelement.beaconY)

				if Abs(tmpelement.sensorX-x)+Abs(tmpelement.sensorY-y) <= tmpmanhatan {
					found = false
					break
				}
			}
			if found && x >= 0 && x <= max && y >= 0 && y <= max {
				fmt.Println(4000000*x + y)
				break loop
			}

			y -= 1
			x += 1
		}

		for y < element.sensorY {
			found := true
			for _, tmpelement := range l {

				tmpmanhatan := Abs(tmpelement.sensorX-tmpelement.beaconX) + Abs(tmpelement.sensorY-tmpelement.beaconY)

				if Abs(tmpelement.sensorX-x)+Abs(tmpelement.sensorY-y) <= tmpmanhatan {
					found = false
					break
				}
			}
			if found && x >= 0 && x <= max && y >= 0 && y <= max {
				fmt.Println(4000000*x + y)
				break loop
			}

			y += 1
			x += 1
		}

		for x > element.sensorX {
			found := true
			for _, tmpelement := range l {

				tmpmanhatan := Abs(tmpelement.sensorX-tmpelement.beaconX) + Abs(tmpelement.sensorY-tmpelement.beaconY)

				if Abs(tmpelement.sensorX-x)+Abs(tmpelement.sensorY-y) <= tmpmanhatan {
					found = false
					break
				}
			}
			if found && x >= 0 && x <= max && y >= 0 && y <= max {
				fmt.Println(4000000*x + y)
				break loop
			}

			y += 1
			x -= 1
		}

		for y > element.sensorX {
			found := true
			for _, tmpelement := range l {

				tmpmanhatan := Abs(tmpelement.sensorX-tmpelement.beaconX) + Abs(tmpelement.sensorY-tmpelement.beaconY)

				if Abs(tmpelement.sensorX-x)+Abs(tmpelement.sensorY-y) <= tmpmanhatan {
					found = false
					break
				}
			}
			if found && x >= 0 && x <= max && y >= 0 && y <= max {
				fmt.Println(4000000*x + y)
				break loop
			}

			y -= 1
			x -= 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
