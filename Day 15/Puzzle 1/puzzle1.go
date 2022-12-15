package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var l []pair //list.New()
	biggestX := 0
	biggestY := 0
	smallestX := math.MaxInt32
	smallestY := math.MaxInt32
	MarginX := 0
	MarginY := 0

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

		if sensorx < smallestX {
			smallestX = sensorx

			if Abs(sensorx-beaconx) > MarginX {
				MarginX = Abs(sensorx - beaconx)
			}
		}

		if sensory < smallestY {
			smallestY = sensory

			if Abs(sensory-beacony) > MarginY {
				MarginY = Abs(sensory - beacony)
			}
		}

		if sensorx > biggestX {
			biggestX = sensorx

			if Abs(sensorx-beaconx) > MarginX {
				MarginX = Abs(sensorx - beaconx)
			}
		}
		if sensory > biggestY {
			biggestY = sensory

			if Abs(sensory-beacony) > MarginY {
				MarginY = Abs(sensory - beacony)
			}
		}

		if beaconx < smallestX {
			smallestX = beaconx

			if Abs(sensorx-beaconx) > MarginX {
				MarginX = Abs(sensorx - beaconx)
			}
		}
		if beacony < smallestY {
			smallestY = beacony

			if Abs(sensory-beacony) > MarginY {
				MarginY = Abs(sensory - beacony)
			}
		}

		if beaconx > biggestX {
			biggestX = beaconx

			if Abs(sensorx-beaconx) > MarginX {
				MarginX = Abs(sensorx - beaconx)
			}
		}
		if beacony > biggestY {
			biggestY = beacony

			if Abs(sensory-beacony) > MarginY {
				MarginY = Abs(sensory - beacony)
			}
		}

		tmp := pair{sensorX: sensorx, sensorY: sensory, beaconX: beaconx, beaconY: beacony}
		l = append(l, tmp)
		//l.PushBack(tmp)

	}

	//rows := biggestY - smallestY + MarginY + 1
	cols := 4*(biggestX-smallestX) + 1 // + 2*MarginX + 2
	offset := 2 * (biggestX - smallestX)
	rowNum := 2000000

	grid := make([]string, cols) // Make one inner slice per iteration and give it size 10
	for j := 0; j < cols; j++ {
		grid[j] = "."
	}

	for _, element := range l {

		if element.sensorY == rowNum {
			grid[element.sensorX+offset] = "S"
		}

		if element.beaconY == rowNum {
			grid[element.beaconX+offset] = "B"
		}

	}

	for _, element := range l {

		manhatan := Abs(element.sensorX-element.beaconX) + Abs(element.sensorY-element.beaconY)

		for j := 0; j < cols; j++ {

			if Abs(element.sensorX-(j-offset))+Abs(element.sensorY-rowNum) <= manhatan {
				if grid[j] == "." {
					grid[j] = "#"
				}
			}

		}
	}

	count := 0
	for j := 0; j < cols; j++ {
		if grid[j] == "#" {
			count += 1
		}
	}

	fmt.Println(count)
	//fmt.Println(cols)

	//for j := 0; j < cols; j++ {
	//	fmt.Print(grid[j])
	//}

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
