package inputs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ParseDecInt(str string) int {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func ParseDecInt64(str string) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func GetToken() string {
	token, err := os.ReadFile("./.token")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(token), "\n")[0]
}

func GetInput(day int, token string) string {

	client := http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	if err != nil {
		panic(err)
	}
	request.AddCookie(&http.Cookie{Name: "session", Value: token})

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return strings.Trim(string(data), "\n")
}

func InputToStrList(input string) []string {
	input = strings.Trim(input, "\n")
	return strings.Split(input, "\n")
}

func StrListToInt64List(input []string) []int64 {
	intlist := make([]int64, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
		intlist = append(intlist, i)
	}
	return intlist
}

func StrListToUInt64List(input []string) []uint64 {
	intlist := make([]uint64, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
		intlist = append(intlist, uint64(i))
	}
	return intlist
}

func StrListToIntList(input []string) []int {
	intlist := make([]int, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
		intlist = append(intlist, int(i))
	}
	return intlist
}

func InputToIntList(input string) []int {
	return StrListToIntList(strings.Split(strings.Trim(input, "\n"), "\n"))
}
func InputToInt64List(input string) []int64 {
	return StrListToInt64List(strings.Split(strings.Trim(input, "\n"), "\n"))
}

func InputToIntMatrice(input string) [][]int {
	input = strings.Trim(input, "\n")
	split := strings.Split(input, "\n")
	matrice := make([][]int, len(split))
	for i, s := range split {
		row := make([]int, len(s))
		for j, c := range s {
			row[j] = int(c - '0')
		}
		matrice[i] = row
	}
	return matrice
}
