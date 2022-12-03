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
		
		if player == "X" {
			if opponent == "A" {
				score = 3
			} else if opponent == "B" {
				score = 1
			} else if opponent == "C" {
				score = 2
			}
		} else if player == "Y" {
			if opponent == "A" {
				score = 3 + 1
			} else if opponent == "B" {
				score = 3 + 2
			} else if opponent == "C" {
				score = 3 + 3
			}
		} else if player == "Z" {
			if opponent == "A" {
				score = 6 + 2
			} else if opponent == "B" {
				score = 6 + 3
			} else if opponent == "C" {
				score = 6 + 1
			}
		}
		
		sum += score
		score = 0
		
    }
	
	fmt.Println(sum)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}