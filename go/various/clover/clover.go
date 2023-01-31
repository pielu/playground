package clover

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Clover() string {
	var lines []string
	var counter int = 0
	var size int
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	size, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if counter <= size {
			l := strings.Split(line, " ")
			for i, j := 1, len(l)-2; i < j; i, j = i+1, j-1 {
				l[i], l[j] = l[j], l[i]
			}
			lines = append(lines, strings.Join(l, " "))
		}
		counter++
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return lines[len(lines)-1]
}
