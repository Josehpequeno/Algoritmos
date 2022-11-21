package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InsertSort(a []int) {
	for j := 1; j < len(a); j++ {
		chave := a[j]
		i := j - 1
		for i > -1 && a[i] > chave {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = chave
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert Sort")
	fmt.Println("Enter the numbers separated by spaces")
	fmt.Println("---------------------")
	array := []int{}
loop: //label do loop
	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n') // reading the entries
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		s := strings.Fields(text)
		for i := range s {
			j, err := strconv.Atoi(s[i]) //converting to int
			if err != nil {
				continue loop
			}
			array = append(array, j)
		}
		fmt.Println("----------------------")
		fmt.Println("Array before Insert Sort: ", array)
		InsertSort(array)
		fmt.Println("Array after Insert Sort: ", array)
		fmt.Println("----------------------")
	}
}
