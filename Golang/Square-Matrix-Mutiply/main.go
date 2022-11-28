package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	array  [][]float64
	rows   int
	colums int
}

// func square_matrix_mutiply_recursive(a Matrix, b Matrix) [][]float64 {
// 	n := a.rows
// 	c := Matrix{make([][]float64, n*n), n, n}
// 	if n == 1 {
// 		c.array[0][0] = a.array[0][0] * b.array[0][0]
// 	} else {

//			c.array[0][0] = square_matrix_mutiply_recursive(a[0][0], b[0][0]) + square_matrix_mutiply_recursive(a[1][2], b[2][1])
//		}
//		return c.array
//	}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Square Matrix Mutiply")
	fmt.Println("Enter the number of square matrix separated by spaces\n* type quit to quit")
	fmt.Println("---------------------")
	first := []float64{}
	second := []float64{}
	a_matrix := Matrix{}
	b_matrix := Matrix{}
loop: //label loop
	for {
	first_loop:
		for {
			fmt.Print("First Matrix => ")
			text, _ := reader.ReadString('\n') // reading the entries
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			if text == "quit" {
				break loop
			}
			s := strings.Fields(text)
			for i := range s {
				j, err := strconv.ParseFloat(s[i], 64) //converting to float64
				if err != nil {
					continue
				}
				first = append(first, j)
			}
			fmt.Println("----------------------")
			sqrt := math.Sqrt(float64(len(first)))
			trunc := math.Trunc(sqrt)
			if sqrt == trunc {
				fmt.Println("Square matrix is a square matrix")
				fmt.Println("First Matrix:")
				a_matrix.array = make([][]float64, int(sqrt))
				a_matrix.rows = int(sqrt)
				a_matrix.colums = int(sqrt)
				count := 0
				it := 0
				for i := 0; i < len(first); i++ {
					if count > int(sqrt)-1 {
						it++
						count = 0
					}
					a_matrix.array[it] = append(a_matrix.array[it], first[i])
					count++
				}
				fmt.Println("matrix", a_matrix.array)
				break first_loop
			} else {
				fmt.Println("Square matrix is not a square matrix", first)
			}
		}
	second_loop:
		for {
			fmt.Print("Second Matrix => ")
			text, _ := reader.ReadString('\n') // reading the entries
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			if text == "quit" {
				break loop
			}
			s := strings.Fields(text)
			for i := range s {
				j, err := strconv.ParseFloat(s[i], 64) //converting to float64
				if err != nil {
					continue
				}
				second = append(second, j)
			}
			fmt.Println("----------------------")
			sqrt := math.Sqrt(float64(len(second)))
			trunc := math.Trunc(sqrt)
			if sqrt == trunc {
				fmt.Println("Second Matrix: ", second)
				b_matrix.array = make([][]float64, int(sqrt))
				b_matrix.rows = int(sqrt)
				b_matrix.colums = int(sqrt)
				count := 0
				it := 0
				for i := 0; i < len(second); i++ {
					if count > int(sqrt)-1 {
						it++
						count = 0
					}
					b_matrix.array[it] = append(b_matrix.array[it], second[i])
					count++
				}
				fmt.Println("matrix", b_matrix.array)
				break second_loop
			} else {
				fmt.Println("Square matrix is not a square matrix", second)
			}
		}
		fmt.Println("----------------------")
		first = first[:0] //clear the matrix
		second = second[:0]
	}
}
