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
	// array := []int{2, 3, 4, 5, 1}
	// fmt.Println("array: ", array)
	// InsertSort(array)
	// fmt.Println("array: ", array)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert Sort")
	fmt.Println("Digite os números separados por espaços")
	fmt.Println("---------------------")
	array := []int{}
loop: //label do loop
	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n') // lendo as entradas
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		s := strings.Fields(text)
		for i := range s {
			j, err := strconv.Atoi(s[i]) //convertendo para int
			if err != nil {
				continue loop
			}
			array = append(array, j)
		}
		fmt.Println("----------------------")
		fmt.Println("Array antes do InsertSort: ", array)
		InsertSort(array)
		fmt.Println("Array depois do InsertSort: ", array)
		fmt.Println("----------------------")
	}
}
