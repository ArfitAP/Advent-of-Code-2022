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
	currentFloor := 0

	scanner.Scan()
	line := scanner.Text()
	gasCycle := len(line)

	//lastCycleHeight := 0
	//thisCycleHeight := 0

	for c := 1; c <= 443; c++ {

		//thisCycleHeight = getHighest(chamber) + currentFloor

		runOneCycle(&chamber, &rockType, &cycleIndex, gasCycle, line, &currentFloor)

		//lastCycleHeight = getHighest(chamber) + currentFloor

		//fmt.Println(lastCycleHeight - thisCycleHeight)
	}

	/*thisCycleHeight = getHighest(chamber) + currentFloor

	for i := 1; i <= 15906670; i++ {

		if i%20 == 0 {
			var removedHieght = cut(&chamber)
			currentFloor += removedHieght
		}

		FallOneRock(&chamber, &rockType, &cycleIndex, gasCycle, line)

	}

	lastCycleHeight = getHighest(chamber) + currentFloor

	fmt.Println(lastCycleHeight - thisCycleHeight)*/

	/*for i := 0; i < 7; i++ {
		fmt.Println(chamber[i])
	}*/

	fmt.Print("End height:")
	fmt.Println(getHighest(chamber) + currentFloor)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func runOneCycle(chamber *[7][]int, rockType *int, cycleIndex *int, gasCycle int, line string, currentFloor *int) {

	numberOfRocks := gasCycle * 5

	for i := 0; i < numberOfRocks; i++ {
		if i%20 == 0 {
			var removedHieght = cut(chamber)
			*currentFloor += removedHieght
		}

		FallOneRock(chamber, rockType, cycleIndex, gasCycle, line)
	}
}

func FallOneRock(chamber *[7][]int, rockType *int, cycleIndex *int, gasCycle int, line string) {

	highestRock := getHighest(*chamber)

	currY := highestRock + 4
	currX := 2

	for true {

		if line[*cycleIndex] == '>' && currX+getRockWidth(*rockType) < 7 && canGoRight(*chamber, *rockType, currX, currY) {
			currX += 1
		} else if line[*cycleIndex] == '<' && currX > 0 && canGoLeft(*chamber, *rockType, currX, currY) {
			currX -= 1
		}
		*cycleIndex = (*cycleIndex + 1) % gasCycle

		if rockCanFall(*chamber, *rockType, currX, currY) {
			currY -= 1
		} else {
			// Increase chamber heights
			dropRock(chamber, *rockType, currX, currY)

			*rockType = (*rockType + 1) % 5

			break
		}
	}
}

func cut(chamber *[7][]int) int {

	highestCut := 0
	for i0 := len(chamber[0]) - 1; i0 >= 0; i0-- {

		if chamber[0][i0] <= highestCut {
			continue
		}

		for i1 := len(chamber[1]) - 1; i1 >= 0; i1-- {
			if chamber[1][i1] >= chamber[0][i0]-1 && chamber[1][i1] <= chamber[0][i0]+1 {
				for i2 := len(chamber[2]) - 1; i2 >= 0; i2-- {
					if chamber[2][i2] >= chamber[1][i1]-1 && chamber[2][i2] <= chamber[1][i1]+1 {
						for i3 := len(chamber[3]) - 1; i3 >= 0; i3-- {
							if chamber[3][i3] >= chamber[2][i2]-1 && chamber[3][i3] <= chamber[2][i2]+1 {
								for i4 := len(chamber[4]) - 1; i4 >= 0; i4-- {
									if chamber[4][i4] >= chamber[3][i3]-1 && chamber[4][i4] <= chamber[3][i3]+1 {
										for i5 := len(chamber[5]) - 1; i5 >= 0; i5-- {
											if chamber[5][i5] >= chamber[4][i4]-1 && chamber[5][i5] <= chamber[4][i4]+1 {
												for i6 := len(chamber[6]) - 1; i6 >= 0; i6-- {
													if chamber[6][i6] >= chamber[5][i5]-1 && chamber[6][i6] <= chamber[5][i5]+1 {
														highestCut = min(chamber[0][i0], chamber[1][i1])
														highestCut = min(highestCut, chamber[2][i2])
														highestCut = min(highestCut, chamber[3][i3])
														highestCut = min(highestCut, chamber[4][i4])
														highestCut = min(highestCut, chamber[5][i5])
														highestCut = min(highestCut, chamber[6][i6])
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if highestCut > 0 {

		for i := 0; i < 7; i++ {
			for j := len(chamber[i]) - 1; j >= 0; j-- {
				if chamber[i][j] >= highestCut {

					chamber[i][j] -= highestCut

				} else {
					chamber[i][j] = chamber[i][len(chamber[i])-1]
					chamber[i][len(chamber[i])-1] = 0
					chamber[i] = chamber[i][:len(chamber[i])-1]
				}
			}
		}

	}

	return highestCut
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
