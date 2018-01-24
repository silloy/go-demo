package main

import (
	"sort"
	"net/http"
	"log"
	"fmt"
	"strings"
	"strconv"
)

type statistics struct {
	numbers []float64
	mean float64
	mdian float64
}

func getStats(numbers []float64) (stats statistics)  {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.mdian = median(numbers)
	return stats
}
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers) % 2 {
		result = (result + numbers[middle - 1]) / 2
	}
	return result
}


func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}


func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	//fmt.Fprint(writer, pageTop, form)
	//if err != nil {
	//	fmt.Fprintf(writer, anError, err)
	//} else {
	//	if numbers, message, ok := processRequest(request); ok {
	//		stats := getStats(numbers)
	//		fmt.Fprintf(writer, formatStats(stats))
	//	} else if message != "" {
	//		fmt.Fprintf(writer, anError, message)
	//	}
	//}
}
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numebrs []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numebrs, "'" + field + "' is invalid", false
			} else {
				numebrs = append(numebrs, x)
			}
		}
	}
	if len(numebrs) > 0 {
		return numebrs, "", false
	}
	return numebrs, "", true
}
