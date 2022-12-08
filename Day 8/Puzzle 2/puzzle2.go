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
	score := make([][]int,0)

    for scanner.Scan() {
        	
		line := scanner.Text()
		//fmt.Println(line)
		
		tmpTrees := make([]int, 0)
		tmpScore := make([]int, 0)
		
		for i := 0; i < len(line); i++ {

			tmpTrees = append(tmpTrees, int(line[i]) - 48)
		    tmpScore = append(tmpScore, 0)
		}
		
		grid = append(grid, tmpTrees)
		score = append(score, tmpScore)					
    }
	
	
	for indexRow, row := range grid {
		for indexColumn, height := range row {
			
			scoreD := 0
			scoreU := 0
			scoreR := 0
			scoreL := 0
			
			for i := indexRow + 1; i < len(grid); i++ {
				scoreD++
				if grid[i][indexColumn] >= height {
					break
				}
				
			}
			for i := indexRow - 1; i >= 0; i-- {
				scoreU++
				if grid[i][indexColumn] >= height {
					break
				}
				
			}
			for i := indexColumn + 1; i < len(row); i++ {
				scoreR++
				if grid[indexRow][i] >= height {
					break
				}
				
			}
			for i := indexColumn - 1; i >= 0; i-- {
				scoreL++
				if grid[indexRow][i] >= height {
					break
				}
				
			}
							
			score[indexRow][indexColumn] = scoreD * scoreU * scoreR * scoreL
			
		}
	}
	
	max := 0
	for _, row := range score {
		for _, element := range row {
			
			if element > max {
				max = element
			}
		}
	}

	fmt.Println(max)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}