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
	
	grid := make([][]int,0)
	visible := make([][]bool,0)

    for scanner.Scan() {
        	
		line := scanner.Text()
		//fmt.Println(line)
		
		tmpTrees := make([]int, 0)
		tmpVisible := make([]bool, 0)
		
		for i := 0; i < len(line); i++ {

			tmpTrees = append(tmpTrees, int(line[i]) - 48)
		    tmpVisible = append(tmpVisible, false)
		}
		
		grid = append(grid, tmpTrees)
		visible = append(visible, tmpVisible)					
    }
	
	
	for indexRow, row := range grid {
		for indexColumn, height := range row {
			
			if indexRow == 0 || indexRow == len(grid) - 1 || indexColumn == 0 || indexColumn == len(row) - 1 {
				visible[indexRow][indexColumn] = true
			} else {
				
				isVisibleD := true
				isVisibleU := true
				isVisibleR := true
				isVisibleL := true
				
				for i := indexRow + 1; i < len(grid); i++ {

					if grid[i][indexColumn] >= height {
						isVisibleD = false
					}
				}
				for i := indexRow - 1; i >= 0; i-- {

					if grid[i][indexColumn] >= height {
						isVisibleU = false
					}
				}
				for i := indexColumn + 1; i < len(row); i++ {

					if grid[indexRow][i] >= height {
						isVisibleR = false
					}
				}
				for i := indexColumn - 1; i >= 0; i-- {

					if grid[indexRow][i] >= height {
						isVisibleL = false
					}
				}
								
				visible[indexRow][indexColumn] = isVisibleD || isVisibleU || isVisibleR || isVisibleL
			}
		}
	}
	
	count := 0
	for _, row := range visible {
		for _, element := range row {
			
			if element == true {
				count++
			}
		}

	}

	fmt.Println(count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}