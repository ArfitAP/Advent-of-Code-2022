package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("..\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var chamber [7][]int
	for i := 0; i < 7; i++ {
		chamber[i] = []int{0}
	}
	rockType := 0
	cycleIndex := 0
	//currentFloor := 0

	scanner.Scan()
	line := scanner.Text()
	gasCycle := len(line)

	for i := 0; i < 2022; i++ {

		currY := getHighest(chamber) + 4
		currX := 2

		for true {

			if line[cycleIndex] == '>' && currX+getRockWidth(rockType) < 7 && canGoRight(chamber, rockType, currX, currY) {
				currX += 1
			} else if line[cycleIndex] == '<' && currX > 0 && canGoLeft(chamber, rockType, currX, currY) {
				currX -= 1
			}
			cycleIndex = (cycleIndex + 1) % gasCycle

			if rockCanFall(chamber, rockType, currX, currY) {
				currY -= 1
			} else {
				// Increase chamber heights
				dropRock(&chamber, rockType, currX, currY)

				rockType = (rockType + 1) % 5

				break
			}
		}
	}

	/*for i := 0; i < 7; i++ {
		fmt.Println(chamber[i])
	}*/

	fmt.Println(getHighest(chamber))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func dropRock(chamber *[7][]int, rockType int, currX int, currY int) {

	if rockType == 0 {
		for i := 0; i < 4; i++ {
			chamber[currX+i] = append(chamber[currX+i], currY)
		}
	} else if rockType == 1 {
		chamber[currX] = append(chamber[currX], currY+1)
		chamber[currX+1] = append(chamber[currX+1], currY)
		chamber[currX+1] = append(chamber[currX+1], currY+1)
		chamber[currX+1] = append(chamber[currX+1], currY+2)
		chamber[currX+2] = append(chamber[currX+2], currY+1)
	} else if rockType == 2 {
		chamber[currX] = append(chamber[currX], currY)
		chamber[currX+1] = append(chamber[currX+1], currY)
		chamber[currX+2] = append(chamber[currX+2], currY)
		chamber[currX+2] = append(chamber[currX+2], currY+1)
		chamber[currX+2] = append(chamber[currX+2], currY+2)
	} else if rockType == 3 {
		chamber[currX] = append(chamber[currX], currY)
		chamber[currX] = append(chamber[currX], currY+1)
		chamber[currX] = append(chamber[currX], currY+2)
		chamber[currX] = append(chamber[currX], currY+3)
	} else if rockType == 4 {
		chamber[currX] = append(chamber[currX], currY)
		chamber[currX] = append(chamber[currX], currY+1)
		chamber[currX+1] = append(chamber[currX+1], currY)
		chamber[currX+1] = append(chamber[currX+1], currY+1)
	}
}

func rockCanFall(chamber [7][]int, rockType int, currX int, currY int) bool {

	if rockType == 0 {
		for i := 0; i < 4; i++ {
			for j := len(chamber[currX+i]) - 1; j >= 0; j-- {
				if chamber[currX+i][j] == currY-1 {
					return false
				}
			}
		}
		return true
	} else if rockType == 1 {
		for j := len(chamber[currX]) - 1; j >= 0; j-- {
			if chamber[currX][j] == currY {
				return false
			}
		}
		for j := len(chamber[currX+1]) - 1; j >= 0; j-- {
			if chamber[currX+1][j] == currY-1 {
				return false
			}
		}
		for j := len(chamber[currX+2]) - 1; j >= 0; j-- {
			if chamber[currX+2][j] == currY {
				return false
			}
		}
		return true
	} else if rockType == 2 {
		for i := 0; i < 3; i++ {
			for j := len(chamber[currX+i]) - 1; j >= 0; j-- {
				if chamber[currX+i][j] == currY-1 {
					return false
				}
			}
		}
		return true
	} else if rockType == 3 {
		for i := 0; i < 1; i++ {
			for j := len(chamber[currX+i]) - 1; j >= 0; j-- {
				if chamber[currX+i][j] == currY-1 {
					return false
				}
			}
		}
		return true
	} else if rockType == 4 {
		for i := 0; i < 2; i++ {
			for j := len(chamber[currX+i]) - 1; j >= 0; j-- {
				if chamber[currX+i][j] == currY-1 {
					return false
				}
			}
		}
		return true
	}

	return false
}

func canGoRight(chamber [7][]int, rockType int, currX int, currY int) bool {

	if rockType == 0 {
		for i := len(chamber[currX+4]) - 1; i >= 0; i-- {
			if chamber[currX+4][i] == currY {
				return false
			}

			/*if chamber[currX+4][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 1 {
		for i := len(chamber[currX+2]) - 1; i >= 0; i-- {
			if chamber[currX+2][i] == currY || chamber[currX+2][i] == currY+2 {
				return false
			}
			/*if chamber[currX+2][i] < currY {
				break
			}*/
		}
		for i := len(chamber[currX+3]) - 1; i >= 0; i-- {
			if chamber[currX+3][i] == currY+1 {
				return false
			}
			/*if chamber[currX+3][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 2 {
		for i := len(chamber[currX+3]) - 1; i >= 0; i-- {
			if chamber[currX+3][i] == currY || chamber[currX+3][i] == currY+1 || chamber[currX+3][i] == currY+2 {
				return false
			}
			/*if chamber[currX+3][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 3 {
		for i := len(chamber[currX+1]) - 1; i >= 0; i-- {
			if chamber[currX+1][i] == currY || chamber[currX+1][i] == currY+1 || chamber[currX+1][i] == currY+2 || chamber[currX+1][i] == currY+3 {
				return false
			}
			/*if chamber[currX+1][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 4 {
		for i := len(chamber[currX+2]) - 1; i >= 0; i-- {
			if chamber[currX+2][i] == currY || chamber[currX+2][i] == currY+1 {
				return false
			}
			/*if chamber[currX+2][i] < currY {
				break
			}*/
		}
		return true
	}

	return false
}

func canGoLeft(chamber [7][]int, rockType int, currX int, currY int) bool {

	if rockType == 0 {
		for i := len(chamber[currX-1]) - 1; i >= 0; i-- {
			if chamber[currX-1][i] == currY {
				return false
			}

			/*if chamber[currX-1][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 1 {
		for i := len(chamber[currX]) - 1; i >= 0; i-- {
			if chamber[currX][i] == currY || chamber[currX][i] == currY+2 {
				return false
			}
			/*if chamber[currX][i] < currY {
				break
			}*/
		}
		for i := len(chamber[currX-1]) - 1; i >= 0; i-- {
			if chamber[currX-1][i] == currY+1 {
				return false
			}
			/*if chamber[currX-1][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 2 {
		for i := len(chamber[currX+1]) - 1; i >= 0; i-- {
			if chamber[currX+1][i] == currY+1 || chamber[currX+1][i] == currY+2 {
				return false
			}
			/*if chamber[currX+1][i] < currY {
				break
			}*/
		}
		for i := len(chamber[currX-1]) - 1; i >= 0; i-- {
			if chamber[currX-1][i] == currY {
				return false
			}
			/*if chamber[currX-1][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 3 {
		for i := len(chamber[currX-1]) - 1; i >= 0; i-- {
			if chamber[currX-1][i] == currY || chamber[currX-1][i] == currY+1 || chamber[currX-1][i] == currY+2 || chamber[currX-1][i] == currY+3 {
				return false
			}
			/*if chamber[currX-1][i] < currY {
				break
			}*/
		}
		return true
	} else if rockType == 4 {
		for i := len(chamber[currX-1]) - 1; i >= 0; i-- {
			if chamber[currX-1][i] == currY || chamber[currX-1][i] == currY+1 {
				return false
			}
			/*if chamber[currX-1][i] < currY {
				break
			}*/
		}
		return true
	}

	return false
}

func getHighest(chamber [7][]int) int {

	tmphighest := chamber[0][0]
	for i := 0; i < 7; i++ {
		for j := len(chamber[i]) - 1; j >= 0; j-- {
			if chamber[i][j] > tmphighest {
				tmphighest = chamber[i][j]
			}
		}
	}
	return tmphighest
}

func getRockWidth(rockType int) int {

	switch rockType {
	case 0:
		return 4
	case 1:
		return 3
	case 2:
		return 3
	case 3:
		return 1
	case 4:
		return 2
	default:
		return 0
	}

}

func getRockHeight(rockType int) int {

	switch rockType {
	case 0:
		return 1
	case 1:
		return 3
	case 2:
		return 3
	case 3:
		return 4
	case 4:
		return 2
	default:
		return 0
	}

}
