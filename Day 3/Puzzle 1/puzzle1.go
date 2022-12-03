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
	
	sum := 0
	score := 0
	
    for scanner.Scan() {
        	
		line := scanner.Text()
		//fmt.Println(line)
		
		lineLength := len(line);
		
		firstpart := line[0:(lineLength / 2)]
		secondpart := line[(lineLength / 2):lineLength]
		
		for i := 0; i < len(firstpart); i++ {
			for j := 0; j < len(secondpart); j++ {
				if firstpart[i] == secondpart[j] {
					score = getscore(firstpart[i])
				}
			}
		}
		sum += score
		
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