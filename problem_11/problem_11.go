package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadData(filename string) (data string) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data = string(file)

	return
}

func prepData(data string) (array []string) {
	firstSplit := strings.Split(data, "\r\n")
	array = strings.Split(firstSplit[0], "\n")

	return
}

func splitLine(line string) (array []string) {
	array = strings.Split(line, " ")
	return
}

func stringsToInts(numberStrings []string) (intArray []int) {
	intArray = []int{}

	for i := 0; i < len(numberStrings); i++ {
		if string(numberStrings[i][0]) == "0" {
			number, err := strconv.ParseInt(string(numberStrings[i][1]), 0, 32)
			if err != nil {
				panic(err)
			}

			intArray = append(intArray, int(number))
		} else {
			number, err := strconv.ParseInt(numberStrings[i], 0, 32)
			if err != nil {
				panic(err)
			}

			intArray = append(intArray, int(number))
		}
	}

	return
}

func dataToNestedIntArray(data []string) (masterArray [][]int) {
	masterArray = [][]int{}

	for i := 0; i < len(data)-1; i++ {
		cleanedLine := splitLine(data[i])
		intArray := stringsToInts(cleanedLine)

		masterArray = append(masterArray, intArray)
	}

	return
}

func genGrid(filename string) (masterArray [][]int) {
	data := loadData(filename)
	arrays := prepData(data)
	masterArray = dataToNestedIntArray(arrays)
	return
}

func rotateGrid(grid [][]int) (newGrid [][]int) {
	newGrid = [][]int{}

	for i := 0; i < len(grid); i++ {
		column := getColumn(grid, i)
		newGrid = append(newGrid, column)
	}

	return
}

func flipGrid(grid [][]int) (newGrid [][]int) {
	newGrid = [][]int{}

	for i := 0; i < len(grid); i++ {
		row := []int{}
		for j := len(grid[i]) - 1; j > -1; j-- {
			row = append(row, grid[i][j])
		}
		newGrid = append(newGrid, row)
	}

	return
}

func genDiagonalGrid(grid [][]int) (newGrid [][]int) {
	newGrid = [][]int{}

	for i := 0; i < len(grid); i++ {
		diagonal := getDiagonal(grid, i)
		newGrid = append(newGrid, diagonal)
	}

	rotatedGrid := rotateGrid(grid)

	for i := 0; i < len(rotatedGrid); i++ {
		diagonal := getDiagonal(rotatedGrid, i)
		newGrid = append(newGrid, diagonal)
	}

	return
}

func getColumn(grid [][]int, columnIndex int) (column []int) {
	column = []int{}

	for i := 0; i < len(grid); i++ {
		column = append(column, grid[i][columnIndex])
	}

	return
}

func getDiagonal(grid [][]int, startIndex int) (diagonal []int) {
	diagonal = []int{}

	j := 0

	for i := startIndex; i < len(grid); i++ {

		diagonal = append(diagonal, grid[j][i])

		j++
	}

	// grid[0][0]
	// grid[1][1]
	// 		.
	// 		.
	// 		.
	// grid[19][19]

	// grid[1][0]
	// grid[2][1]
	// grid[3][2]
	// 		.
	// 		.
	// 		.
	// grid[19][18]

	// grid[2][0]
	// grid[3][1]
	// grid[4][2]
	// 		.
	// 		.
	// 		.
	// grid[19][17]

	return
}

func findGreatestOfNValues(grid [][]int, nValues int) (value int, products []int) {
	value = 0

	for i := 0; i < len(grid); i++ {
		column := grid[i]

		tempValue, tempProducts := multiplyNValues(column, nValues)

		if tempValue > value {
			value, products = tempValue, tempProducts
		}
	}

	return
}

func multiplyNValues(listOfData []int, nValues int) (value int, products []int) {
	value = 0
	products = []int{}

	for i := nValues - 1; i < len(listOfData); i++ {
		tempValue := 1
		tempProducts := []int{}
		for n := nValues - 1; n > -1; n-- {
			tempValue *= listOfData[i-n]
			tempProducts = append(tempProducts, listOfData[i-n])
		}
		if tempValue > value {
			value = tempValue
			products = tempProducts
		}
	}

	return
}

func main() {
	grid := genGrid("./problem_11/data")
	rotatedGrid := rotateGrid(grid)

	greatest, products := findGreatestOfNValues(rotatedGrid, 4)
	greatestRow, productsRow := findGreatestOfNValues(grid, 4)

	fmt.Println("Greatest Multiple in Columns: ", greatest, " Products: ", products)
	fmt.Println("Greatest Multiple in Rows: ", greatestRow, " Products: ", productsRow)

	diagonalGrid := genDiagonalGrid(grid)

	greatestDiagonal, productsDiagonal := findGreatestOfNValues(diagonalGrid, 4)
	fmt.Println("Greatest Multiple in Diagonals: ", greatestDiagonal, " Products: ", productsDiagonal)

	flippedGrid := flipGrid(grid)

	flippedDiagonal := genDiagonalGrid(flippedGrid)

	greatestFlippedDiagonal, productsFlippedDiagonal := findGreatestOfNValues(flippedDiagonal, 4)
	fmt.Println("Greatest Multiple in Flipped Diagonals: ", greatestFlippedDiagonal, " Products: ", productsFlippedDiagonal)
}
