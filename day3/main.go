package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var file = flag.String("file", "day3/input.txt", "day3 inputs")
	flag.Parse()
	lines := loadData(*file)
	log.Println(part1(lines))
	log.Println(part2(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		s1 := NewSet[rune]()
		s2 := NewSet[rune]()
		for i, ch := range line {
			if i < len(line)/2 {
				s1.Add(ch)
				continue
			}
			s2.Add(ch)
		}

		inter := s1.Intersect(s2)
		prio, err := priority(inter)
		if err != nil {
			log.Fatal(err)
		}
		sum += prio
	}
	return sum
}

func part2(lines []string) int {
	s := NewSet[rune]()
	sum := 0
	for index, line := range lines {
		// if index == 0 {
		// 	for _, rn := range line {
		// 		s.Add(rn)
		// 	}
		// 	log.Println("initialize")
		// 	continue
		// }

		if (index%3) == 0 && index != 0 {
			prio, err := priority(s)
			if err != nil {
				log.Fatal(err)
			}
			sum += prio
			log.Println("reset outer set")
			s = NewSet[rune]()
		}

		temp := NewSet[rune]()
		for _, rn := range line {
			temp.Add(rn)
		}

		if len(s.m) != 0 {
			temp = s.Intersect(temp)
		}
		s = temp
		log.Println(s)
	}

	// final priority
	prio, err := priority(s)
	if err != nil {
		log.Fatal(err)
	}

	return sum + prio
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

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	temp := NewSet[T]()

	for k, _ := range s.m {
		if _, ok := s2.m[k]; ok {
			temp.Add(k)
		}
	}

	return temp
}

var errRuneValue = errors.New("invalid rune")

func priority(s *Set[rune]) (int, error) {
	for rn := range s.m {
		if rn < 65 {
			return 0, errRuneValue
		}

		if 90 < rn && rn < 97 {
			return 0, errRuneValue
		}

		if 122 < rn {
			return 0, errRuneValue
		}

		temp := rn - 65

		if temp < 26 {
			return int(27 + temp), nil
		}
		return int(rn - 96), nil
	}

	return 0, fmt.Errorf("empty set")
}
