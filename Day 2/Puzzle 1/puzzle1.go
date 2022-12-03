package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
)

func main() {
    file, err := os.Open("..\\input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	
	sum := 0
	score := 0

    for scanner.Scan() {
        	
		line := scanner.Text()
		//fmt.Println(line)
		
		split := strings.Split(line, " ")
		//fmt.Println(split)
		
		
		opponent := split[0]
		player := split[1]
		
		if opponent == "A" {
			if player == "Y" {
				score = 6
			} else if player == "X" {
				score = 3
			}
		} else if opponent == "B" {
			if player == "Z" {
				score = 6
			} else if player == "Y" {
				score = 3
			}
		} else if opponent == "C" {
			if player == "X" {
				score = 6
			} else if player == "Z" {
				score = 3
			}
		}
		
		if player == "X" {
			score += 1
		} else if player == "Y" {
			score += 2
		} else if player == "Z" {
			score += 3
		}
		
		sum += score
		score = 0
		
    }
	
	fmt.Println(sum)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}