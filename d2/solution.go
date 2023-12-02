package d2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution(fileName string) int {
	val := 0
	fl, err := os.OpenFile(fmt.Sprintf("./%s", fileName), os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer fl.Close()
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		if power, _, ok := isPossible(scanner.Text()); ok || !ok {
			val += power
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return val
}

func isPossible(str string) (power, id int, ok bool) {
	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	strarr1 := strings.Split(str, ":")
	idOfGame := strings.Split(strarr1[0], " ")[1]
	shown := strings.Split(strarr1[1], ";")
	mp := make(map[string]int)
	for _, v := range shown {
		for _, v1 := range strings.Split(v, ",") {
			v1 = strings.TrimSpace(v1)
			v1arr := strings.Split(v1, " ")
			num, err := strconv.Atoi(v1arr[0])
			if err != nil {
				fmt.Println(err)
			}
			if num > mp[v1arr[1]] {
				mp[v1arr[1]] = num
			}
		}
	}
	intIdOfGame, err := strconv.Atoi(idOfGame)
	if err != nil {
		fmt.Println(err)
	}
	power = mp["red"] * mp["green"] * mp["blue"]
	// fmt.Println(mp["red"] <= 12 && mp["green"] <= 13 && mp["blue"] <= 14, idOfGame, mp)
	fmt.Println(idOfGame, power)
	//only 12 red cubes, 13 green cubes, and 14 blue cubes
	return power, intIdOfGame, mp["red"] <= 12 && mp["green"] <= 13 && mp["blue"] <= 14
}
