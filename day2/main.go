package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	var filename = flag.String("file", "day2/input.txt", "day2 input file")
	flag.Parse()
	tourneyP1 := loadDataP1(*filename)
	log.Println(tourneyP1.Score())
	tourneyP2 := loadDataP2(*filename)
	log.Println(tourneyP2)
	log.Println(tourneyP2.Score())
}

func loadDataP1(filename string) *tournament {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	t := NewTournament()
	for {
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		opp, err := getThrow(int(line[0]))
		if err != nil {
			log.Fatal(err)
		}
		pl, err := getThrow(int(line[2]))
		if err != nil {
			log.Fatal(err)
		}
		t.rounds = append(t.rounds, round{opponent: opp, player: pl})
	}

	return t
}

func loadDataP2(filename string) *tournament {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	t := NewTournament()
	for {
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		opp, err := getThrow(int(line[0]))
		if err != nil {
			log.Fatal(err)
		}
		res, err := getResult(int(line[2]))
		if err != nil {
			log.Fatal(err)
		}
		pl := getRequiredThrow(opp, res)
		t.rounds = append(t.rounds, round{opponent: opp, player: pl})
	}

	return t
}

type round struct {
	player   throw
	opponent throw
}

func (r round) outcome() int {
	if r.opponent == r.player {
		return Tie
	}

	out := Win

	switch r.player {
	case Rock:
		if r.opponent == Paper {
			out = Loss
		}
	case Paper:
		if r.opponent == Scissors {
			out = Loss
		}
	case Scissors:
		if r.opponent == Rock {
			out = Loss
		}
	}

	return out
}

type tournament struct {
	rounds []round
}

func NewTournament() *tournament {
	return &tournament{
		rounds: make([]round, 0),
	}
}

func (t *tournament) Score() int {
	sum := 0
	for _, r := range t.rounds {
		sum += r.outcome()
		sum += int(r.player)
	}

	return sum
}

type throw int

const (
	Unknown throw = iota
	Rock
	Paper
	Scissors

	Loss int = 0
	Tie  int = 3
	Win  int = 6
)

var errByteChar = errors.New("invalid byte for character input")

func getThrow(input int) (throw, error) {
	t := Unknown
	if input < 65 {
		return t, errByteChar
	}

	if 67 < input && input < 88 {
		return t, errByteChar
	}

	if 90 < input {
		return t, errByteChar
	}

	switch input {
	case 65, 88:
		t = Rock
	case 66, 89:
		t = Paper
	case 67, 90:
		t = Scissors
	}

	return t, nil
}

func getResult(input int) (int, error) {
	if input < 88 {
		return 0, errByteChar
	}

	if 90 < input {
		return 0, errByteChar
	}

	var out int
	switch input {
	case 88:
		out = Loss
	case 89:
		out = Tie
	case 90:
		out = Win
	}
	return out, nil
}

func getRequiredThrow(opp throw, result int) throw {
	if result == Tie {
		return opp
	}

	t := Unknown
	switch opp {
	case Rock:
		if result == Loss {
			t = Scissors
		} else {
			t = Paper
		}
	case Paper:
		if result == Loss {
			t = Rock
		} else {
			t = Scissors
		}
	case Scissors:
		if result == Loss {
			t = Paper
		} else {
			t = Rock
		}
	}
	return t
}
