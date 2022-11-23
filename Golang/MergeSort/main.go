package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func MergeSort(a []int, p int, r int) {
	if r-p > 1 {
		q := int(math.Floor(float64(p+r) / 2))
		MergeSort(a, p, q)
		MergeSort(a, q, r)
		Merge(a, p, q, r)
	}
}

func Merge(a []int, p int, q int, e int) {
	n1 := q - p
	n2 := e - q
	l := make([]int, n1+1)
	r := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		l[i] = a[p+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = a[q+j]
	}
	fmt.Println("===============>*<===============")
	fmt.Println("Left array", l[0:len(l)-1])
	fmt.Println("Right array", r[0:len(r)-1])
	l[n1] = math.MaxInt64
	r[n2] = math.MaxInt64
	i := 0
	j := 0
	for k := p; k < e; k++ {
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}
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
		fmt.Println("Array before Merge Sort: ", array)
		MergeSort(array, 0, len(array)) // start, len
		fmt.Println("===============>*<===============")
		fmt.Println("Array after Merge Sort: ", array)
		fmt.Println("----------------------")
	}
}
