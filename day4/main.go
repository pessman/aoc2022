package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var file = flag.String("file", "day4/input.txt", "day4 inputs")
	flag.Parse()
	lines := loadData(*file)
	log.Println(run(lines))
}

func loadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func run(lines []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, line := range lines {
		splitLine := strings.Split(line, ",")
		one := strings.Split(splitLine[0], "-")
		two := strings.Split(splitLine[1], "-")

		a, err := strconv.Atoi(one[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(one[1])
		if err != nil {
			log.Fatal(err)
		}
		x, err := strconv.Atoi(two[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(two[1])
		if err != nil {
			log.Fatal(err)
		}

		if Contained(a, b, x, y) {
			part1++
		}

		if Overlap(a, b, x, y) {
			part2++
		}
	}

	return part1, part2
}

func Contained(a, b, x, y int) bool {
	return contained(a, b, x, y) || contained(x, y, a, b)
}

func contained(a, b, x, y int) bool {
	return (a <= x && x <= b) && (a <= y && y <= b)
}

func Overlap(a, b, x, y int) bool {
	return overlap(a, b, x, y) || overlap(x, y, a, b)
}

func overlap(a, b, x, y int) bool {
	return (a <= x && x <= b) || (a <= y && y <= b)
}
