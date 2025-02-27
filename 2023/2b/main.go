package main

import (
	utils "AdventsForCode/Utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution(in))
}

func solution(data string) int {
	var sum int

	for _, line := range convert(data) {
		sum += line["red"] * line["green"] * line["blue"]
	}

	return sum
}

func convert(data string) []map[string]int {
	var strLines = strings.Split(data, "\n")
	var lines []map[string]int

	for _, line := range strLines {
		var currLine = map[string]int{}

		game := strings.Split(line, ": ")[1]
		rolls := strings.Split(game, "; ")

		for _, roll := range rolls {
			dice := strings.Split(roll, ", ")

			for _, die := range dice {
				action := strings.Split(die, " ")

				if currLine[action[1]] < utils.ToInt(action[0]) {
					currLine[action[1]] = utils.ToInt(action[0])
				}
			}
		}

		lines = append(lines, currLine)
	}

	return lines
}

const in = `Game 1: 1 green, 2 red, 6 blue; 4 red, 1 green, 3 blue; 7 blue, 5 green; 6 blue, 2 red, 1 green
Game 2: 1 green, 17 red; 1 blue, 6 red, 7 green; 2 blue, 4 red, 7 green; 1 green, 6 red, 2 blue
Game 3: 6 red, 15 blue, 15 green; 1 green, 4 red, 12 blue; 14 blue, 9 red, 1 green; 2 red, 15 blue, 12 green
Game 4: 8 green, 10 blue, 6 red; 20 blue, 4 red; 17 blue, 2 green, 3 red; 4 blue, 2 green, 3 red; 10 red, 3 blue, 3 green; 5 green, 14 blue, 6 red
Game 5: 3 green, 8 blue, 2 red; 11 red, 6 green, 11 blue; 8 red, 5 blue, 2 green
Game 6: 2 blue, 12 red, 2 green; 3 green, 2 red; 3 green, 3 blue, 10 red; 7 red, 2 blue, 4 green; 1 red, 2 blue, 5 green
Game 7: 1 red, 8 blue, 2 green; 1 red, 2 blue, 12 green; 1 blue; 3 blue, 3 green
Game 8: 10 green, 4 red, 4 blue; 12 green, 1 blue; 1 red, 13 green, 2 blue; 12 green, 3 blue; 9 green, 7 red
Game 9: 1 green, 1 blue, 3 red; 3 blue, 3 red, 8 green; 6 blue, 4 red, 6 green; 2 red, 7 green; 1 red, 10 green, 13 blue; 5 red, 1 blue, 1 green
Game 10: 9 green, 3 red, 3 blue; 12 green, 2 blue; 18 green, 1 blue; 14 green; 2 blue, 9 green, 2 red
Game 11: 14 green; 2 green, 2 red, 11 blue; 9 blue, 3 green
Game 12: 9 green, 3 blue, 8 red; 1 green, 2 blue, 3 red; 4 blue, 8 red, 10 green; 3 blue, 7 red, 8 green; 3 blue, 5 red, 7 green; 2 blue, 5 green
Game 13: 6 red, 1 blue, 10 green; 7 red, 1 green; 8 red, 2 green, 1 blue
Game 14: 2 red, 4 blue, 2 green; 2 green, 5 red, 1 blue; 1 red, 6 blue
Game 15: 7 blue, 3 red; 13 blue, 8 red, 1 green; 1 green, 5 red, 13 blue; 8 red, 5 blue; 4 red, 3 blue, 1 green; 12 blue, 8 red, 1 green
Game 16: 5 blue, 1 green, 2 red; 2 blue, 20 green; 4 blue, 1 red, 17 green; 10 green, 5 blue, 2 red
Game 17: 6 red, 13 blue, 8 green; 12 blue, 7 green, 9 red; 19 blue, 5 red; 2 green, 8 red, 14 blue
Game 18: 5 red, 2 green, 1 blue; 8 blue, 17 red, 9 green; 2 blue, 1 green; 4 blue, 10 red; 5 blue, 4 red, 6 green
Game 19: 5 red, 12 green; 8 red, 13 green, 1 blue; 1 red, 1 green, 3 blue; 5 green, 5 red
Game 20: 11 red, 8 blue; 9 red, 2 green, 13 blue; 2 red, 1 green, 2 blue; 1 green, 9 blue, 13 red; 3 blue, 5 red, 1 green
Game 21: 1 red, 4 green, 11 blue; 3 green, 15 blue; 6 green, 7 red, 14 blue; 15 blue, 6 green, 10 red; 6 red, 16 blue, 2 green
Game 22: 2 blue, 15 green, 2 red; 3 blue, 6 green, 1 red; 2 blue, 5 green, 1 red; 6 green, 2 red, 2 blue; 4 green, 2 blue; 4 blue, 1 red, 15 green
Game 23: 2 blue, 1 green, 12 red; 5 blue, 11 red, 4 green; 12 red, 4 blue; 12 red, 2 green, 5 blue
Game 24: 4 blue, 7 red; 3 red, 3 blue; 1 red, 4 blue; 2 green, 6 red, 6 blue; 7 red, 1 green, 2 blue; 6 red, 1 blue, 1 green
Game 25: 5 green, 9 blue; 6 green, 7 red, 2 blue; 1 red, 3 blue, 7 green; 9 blue, 3 red; 5 green, 9 blue, 2 red
Game 26: 6 red, 4 blue; 2 blue, 4 green; 3 green, 5 red, 5 blue; 4 green, 6 red, 3 blue; 4 green, 7 red, 4 blue
Game 27: 15 green, 1 blue, 12 red; 12 red, 1 green; 1 red, 1 blue, 5 green; 13 green, 6 red, 1 blue; 5 red, 1 blue, 1 green; 11 red, 14 green
Game 28: 3 blue, 2 green, 10 red; 5 blue, 2 green; 4 green, 3 blue, 11 red
Game 29: 10 blue, 2 red; 17 green, 7 blue, 2 red; 1 blue, 8 green, 1 red; 10 green, 2 red, 3 blue
Game 30: 10 green, 8 red, 1 blue; 4 blue, 7 green, 14 red; 2 blue, 14 red, 11 green; 1 blue, 13 green, 12 red; 5 blue, 2 red, 4 green; 4 green, 5 red
Game 31: 4 green, 11 red, 11 blue; 3 blue, 11 red; 5 blue, 7 red, 3 green; 10 blue, 5 green, 1 red
Game 32: 4 red, 8 blue, 1 green; 14 red, 7 blue, 4 green; 13 red, 3 blue, 9 green; 3 red, 1 green, 8 blue; 8 green, 8 red, 5 blue
Game 33: 6 red, 10 blue, 7 green; 19 blue, 1 red; 6 green, 11 red, 11 blue; 2 green, 2 blue, 12 red; 3 red, 13 blue, 7 green; 6 green, 4 red, 2 blue
Game 34: 3 red, 3 green, 15 blue; 7 green, 15 blue; 3 red, 2 green, 8 blue; 19 green, 18 blue
Game 35: 2 green, 1 blue; 2 green, 2 blue, 1 red; 3 blue, 1 red, 1 green; 4 blue, 1 red
Game 36: 1 red, 11 green; 1 green, 1 blue; 8 blue; 2 green, 3 red; 1 red
Game 37: 4 blue, 3 red; 12 blue, 13 red; 2 red, 2 green, 8 blue
Game 38: 8 red, 2 blue; 1 green, 2 red; 8 red, 2 green, 1 blue; 16 red, 2 green; 7 red, 2 blue, 2 green
Game 39: 6 green, 1 blue, 5 red; 14 green, 8 blue, 6 red; 8 red, 10 blue, 1 green; 14 green, 9 red; 17 blue, 5 red; 1 blue, 7 green, 1 red
Game 40: 4 red, 8 blue, 3 green; 13 blue, 1 red; 3 blue, 7 red, 3 green; 3 green, 8 blue, 10 red; 3 green, 20 blue, 5 red
Game 41: 1 blue, 2 green; 11 green, 2 blue; 5 blue; 15 red, 8 green, 5 blue
Game 42: 1 green, 12 blue, 1 red; 6 blue, 1 green, 5 red; 1 red, 11 blue, 4 green; 3 red, 17 blue, 1 green; 1 red, 11 blue; 9 blue, 6 green, 3 red
Game 43: 16 blue, 13 green, 1 red; 17 blue, 7 red, 10 green; 13 green, 5 red, 7 blue
Game 44: 2 blue, 4 red; 15 green, 7 red; 2 green, 1 blue; 6 red, 13 green
Game 45: 5 green, 1 blue, 8 red; 4 red, 1 blue, 5 green; 1 green, 3 red; 1 green, 2 blue, 6 red; 4 red, 3 green, 2 blue; 2 red, 2 blue, 5 green
Game 46: 1 green, 1 red, 6 blue; 11 blue; 1 red, 1 green, 7 blue; 8 blue; 1 green, 7 blue, 2 red
Game 47: 7 green, 9 blue, 7 red; 11 red, 13 blue, 5 green; 12 green, 12 blue, 5 red; 4 blue, 8 green, 7 red
Game 48: 11 green, 7 red, 2 blue; 2 blue, 10 green, 3 red; 1 blue, 2 red, 1 green; 4 green, 2 red, 7 blue; 7 red, 4 green, 2 blue
Game 49: 1 red, 2 blue, 5 green; 2 green, 4 blue; 5 blue, 2 green, 1 red; 9 blue, 1 green; 7 blue
Game 50: 8 green, 9 blue, 2 red; 2 green, 5 blue; 14 green, 1 red, 8 blue
Game 51: 1 green, 2 blue; 12 blue, 1 red; 2 blue
Game 52: 3 red, 2 blue, 2 green; 4 red, 4 green, 7 blue; 2 blue, 4 red, 1 green; 3 green; 1 red, 9 green, 7 blue
Game 53: 9 blue, 12 red, 7 green; 8 blue, 6 green; 1 green, 8 blue, 9 red; 12 red, 6 green; 9 blue, 14 red, 10 green; 7 red, 3 green, 5 blue
Game 54: 8 green, 5 blue, 5 red; 4 green, 13 blue, 2 red; 2 blue, 5 red, 1 green; 3 red, 3 green, 10 blue
Game 55: 17 red, 15 green, 17 blue; 6 red, 5 green, 7 blue; 17 green, 6 blue, 5 red
Game 56: 7 blue, 6 red, 7 green; 10 green, 3 red; 9 red, 3 blue, 5 green
Game 57: 5 blue, 11 red, 1 green; 13 red, 1 green, 2 blue; 2 blue, 4 red; 1 green, 10 red, 1 blue; 1 green, 8 red
Game 58: 1 red, 2 green, 9 blue; 1 green, 1 blue, 1 red; 2 red, 6 blue, 2 green; 14 blue, 1 green, 1 red; 5 blue, 1 red, 2 green; 14 blue, 2 green
Game 59: 9 green, 2 blue, 5 red; 9 red, 5 green; 10 red, 1 blue, 8 green
Game 60: 8 blue, 6 red, 4 green; 3 red, 12 green, 9 blue; 4 blue, 5 red, 5 green; 4 red, 8 blue; 7 green, 12 blue, 6 red
Game 61: 5 blue, 13 red, 1 green; 5 red, 5 blue; 1 red, 3 blue; 1 green, 9 red; 10 red, 3 blue, 1 green
Game 62: 1 blue, 13 red; 4 blue, 5 red; 11 blue, 8 red, 1 green
Game 63: 14 blue, 5 red; 9 blue, 14 green, 5 red; 3 red, 8 green, 15 blue; 4 blue, 15 green, 6 red
Game 64: 13 red, 6 blue, 11 green; 12 red, 1 blue, 8 green; 1 red, 17 green; 13 red, 12 green, 7 blue
Game 65: 4 red, 17 blue, 3 green; 2 green, 12 blue, 9 red; 2 green, 17 blue, 5 red; 1 red, 1 green, 4 blue; 9 red, 16 blue; 7 blue, 9 red
Game 66: 10 blue, 10 green, 5 red; 10 green, 3 blue, 5 red; 1 red, 1 green, 10 blue; 2 green, 5 red, 20 blue; 8 blue, 11 green, 13 red; 2 green, 18 blue, 2 red
Game 67: 6 red, 1 green; 5 red, 10 blue; 6 blue, 6 red
Game 68: 4 green, 1 red, 5 blue; 5 red, 5 blue; 7 red, 6 green; 8 red, 1 blue
Game 69: 2 blue, 11 red; 4 red, 6 green, 1 blue; 4 red, 1 blue, 14 green
Game 70: 15 red, 8 blue, 5 green; 5 green, 2 red, 8 blue; 8 red, 3 green, 10 blue
Game 71: 4 blue, 2 red; 12 green, 4 blue; 10 green
Game 72: 3 blue, 4 green, 6 red; 6 red, 5 green, 8 blue; 10 red, 6 green, 5 blue; 1 green, 2 blue; 10 red, 5 blue, 4 green
Game 73: 5 blue, 1 red; 1 green, 11 blue; 10 blue; 12 blue, 1 red; 1 red, 9 blue; 7 blue, 1 green, 1 red
Game 74: 7 green, 6 blue, 7 red; 7 blue, 6 green, 15 red; 7 red, 5 blue, 1 green; 1 blue, 6 red, 8 green; 8 green, 14 red, 3 blue
Game 75: 8 green, 3 red, 3 blue; 1 blue, 6 red, 7 green; 9 green, 3 blue; 3 blue, 9 green, 6 red; 4 blue, 1 red, 3 green; 4 green, 1 blue, 16 red
Game 76: 4 blue, 3 green; 2 blue, 1 red, 6 green; 12 blue; 1 green, 14 blue
Game 77: 5 green, 10 red, 11 blue; 3 red; 8 green, 6 red, 9 blue
Game 78: 7 red, 7 green; 8 blue; 6 green, 7 red, 5 blue
Game 79: 11 blue, 2 red, 4 green; 2 green, 3 red, 15 blue; 1 green, 15 blue, 1 red
Game 80: 3 red, 17 green, 8 blue; 8 green, 10 blue; 4 green, 1 red, 14 blue
Game 81: 17 green, 10 red, 10 blue; 9 green, 9 blue, 7 red; 11 red, 11 green, 4 blue; 15 blue, 5 red; 11 blue, 8 green, 15 red; 3 green, 16 red
Game 82: 8 green, 9 blue, 1 red; 1 red, 8 green, 9 blue; 2 green, 12 blue
Game 83: 2 green, 11 red, 20 blue; 20 blue, 1 green, 4 red; 2 green, 6 red, 20 blue; 17 blue, 10 red
Game 84: 1 green, 9 red; 4 blue, 4 green; 1 green, 6 red, 14 blue
Game 85: 5 red, 10 green, 9 blue; 8 blue, 3 green, 2 red; 4 blue, 14 green, 3 red; 5 red, 4 blue
Game 86: 8 blue, 9 green, 5 red; 5 red, 10 green, 1 blue; 15 blue, 1 red, 2 green; 8 red, 8 blue, 10 green
Game 87: 13 green, 2 red, 4 blue; 3 red, 11 green, 9 blue; 6 blue, 3 red, 12 green
Game 88: 2 red, 7 blue, 3 green; 2 blue, 9 red; 9 red, 6 blue, 7 green; 6 green, 13 blue, 9 red; 6 green, 2 red, 15 blue; 1 red, 8 green, 7 blue
Game 89: 11 red, 1 blue, 2 green; 6 blue, 5 green, 4 red; 15 red, 4 green, 5 blue; 11 red, 3 blue, 10 green; 6 blue, 13 green, 12 red
Game 90: 2 red, 2 blue, 4 green; 2 red, 2 blue; 9 green, 1 red, 1 blue; 5 green, 1 red; 7 green, 2 red; 2 green, 1 blue
Game 91: 5 blue, 3 red, 1 green; 1 red, 4 blue, 6 green; 6 blue, 6 green, 5 red
Game 92: 16 green, 1 blue, 12 red; 18 green, 14 red, 1 blue; 16 red, 1 green; 4 blue, 16 red, 18 green
Game 93: 9 red, 8 blue, 14 green; 1 blue, 1 green, 6 red; 4 blue, 4 red, 14 green
Game 94: 11 green, 4 blue, 2 red; 1 red, 1 green, 1 blue; 4 red, 1 blue, 2 green
Game 95: 5 blue, 2 red, 9 green; 5 blue, 8 green; 1 green, 15 blue; 5 red, 9 green, 5 blue; 3 green, 17 blue, 5 red
Game 96: 2 green, 14 blue, 1 red; 3 green, 3 red, 14 blue; 2 red, 2 green, 13 blue
Game 97: 2 green, 2 red; 2 blue, 1 green; 7 blue, 3 red
Game 98: 2 red, 1 blue, 12 green; 2 blue, 10 green, 5 red; 11 green, 9 blue; 6 blue, 17 green; 7 blue, 9 green, 9 red; 1 red, 11 green, 9 blue
Game 99: 2 green, 9 red, 1 blue; 3 green, 1 blue, 14 red; 5 green, 6 blue; 1 blue, 2 green, 3 red; 4 blue, 10 red, 1 green
Game 100: 4 green, 4 blue, 15 red; 3 green, 1 red, 13 blue; 5 green, 5 blue, 10 red`
