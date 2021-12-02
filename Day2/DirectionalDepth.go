package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	forward := "f"
	up := "u"
	down := "d"

	horizontal := int64(0)
	vertical := int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text()[0:1] == forward {
			f, err := strconv.ParseInt(scanner.Text()[len(scanner.Text())-1:len(scanner.Text())], 10, 64)
			if err != nil {
				fmt.Printf("wrong")
			}
			horizontal = horizontal + f
		} else if scanner.Text()[0:1] == up {
			u, err := strconv.ParseInt(scanner.Text()[len(scanner.Text())-1:len(scanner.Text())], 10, 64)
			if err != nil {
				fmt.Printf("wrong")
			}
			vertical = vertical - u
		} else if scanner.Text()[0:1] == down {
			d, err := strconv.ParseInt(scanner.Text()[len(scanner.Text())-1:len(scanner.Text())], 10, 64)
			if err != nil {
				fmt.Printf("wrong")
			}
			vertical = vertical + d
		}
	}
	fmt.Println(vertical)
	fmt.Println(horizontal)
	fmt.Println(vertical * horizontal)
}
