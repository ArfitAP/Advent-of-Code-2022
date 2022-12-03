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
	currLine := 0
	currItems := ""
	possibleItems := ""
	
    for scanner.Scan() {
        	
		line := scanner.Text()
		currLine++

		
		if len(currItems) == 0 {
			currItems = line
		} else {
			for i := 0; i < len(line); i++ {
				if strings.Contains(currItems, string(line[i])) && strings.Contains(possibleItems, string(line[i])) == false {
					possibleItems += string(line[i])
				}							
			}
		}
			
		if currLine == 3 {
			for i := 0; i < len(possibleItems); i++ {
				if strings.Contains(line, string(possibleItems[i])) {
					score = getscore(possibleItems[i])
					break
				}							
			}
			currLine = 0
			currItems = ""
			possibleItems = ""
			sum += score
		}		
		
    }
	
	fmt.Println(sum)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func getscore(c byte) int {
	
	if c >= 'a' && c <= 'z'{
		return int(c) - 'a' + 1
	} else {
		return int(c) - 'A' + 27
	}
	return 0
}