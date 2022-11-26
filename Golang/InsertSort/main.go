package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insert_sort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i > -1 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert Sort")
	fmt.Println("Enter the numbers separated by spaces\n* type clear to clear array, quit to quit")
	fmt.Println("---------------------")
	array := []int{}
loop: //label loop
	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n') // reading the entries
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "clear" {
			array = array[:0] //clear the array
			continue loop
		}
		if text == "quit" {
			break loop
		}
		s := strings.Fields(text)
		for i := range s {
			j, err := strconv.Atoi(s[i]) //converting to int
			if err != nil {
				continue
			}
			array = append(array, j)
		}
		fmt.Println("----------------------")
		fmt.Println("Array before Insert Sort: ", array)
		insert_sort(array)
		fmt.Println("Array after Insert Sort: ", array)
		fmt.Println("----------------------")
	}
}
