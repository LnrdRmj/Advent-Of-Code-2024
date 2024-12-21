package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

    file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
    defer file.Close() // Ensure the file is closed when done

    // Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

    reports := [][]int{}
    var cont int // defaults to 0
    for scanner.Scan() && cont < 10 {
        line := scanner.Text()

        // fmt.Println(line)
        err, result := arrayOfStringToInteger(strings.Split(line, " "))
        if err != nil {
            fmt.Println("Error converting numbers from file:", err)
            return 
        }
        reports = append(reports, result)
        // cont++
    }
    // fmt.Println(reports)

    safeReportsCounter := 0
    for _, arr := range reports {
        // fmt.Println(arr)
        // fmt.Println(reportIsSafePart2(arr))
        if (reportIsSafe(arr, true)) {
            safeReportsCounter++ 
        } else {
            // fmt.Println("Report is not safe", arr)
        }
    }

    fmt.Println("Safe reports", safeReportsCounter)

    // Check for errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}

func reportIsSafe(report []int, tollerance bool) bool {
    
    if (len(report) == 1) {
        return true
    }

    var firstEl, secondEl = report[0], report[1]
    desc := firstEl > secondEl

    for i := 0; i < len(report) - 1; i++ {
        if report[i] == report[i + 1] {
            if tollerance == false { 
                return false
            } else {
                result := reportIsSafeWithProblemDampener(report)
                if result == true {
                    fmt.Println("Report is safe thx to tollerance", report)
                } else {
                    fmt.Println("Report not safe, even with tollerance", report)
                }
                return result
            }
        } else if !(((report[i] > report[i + 1] && desc) || (report[i] < report[i + 1] && !desc)) && (math.Abs(float64(report[i] - report[i + 1])) <= 3)){
            if tollerance == false { 
                return false
            } else {
                result := reportIsSafeWithProblemDampener(report)
                if result == true {
                    fmt.Println("Report is safe thx to tollerance", report)
                } else {
                    fmt.Println("Report not safe, even with tollerance", report)
                }
                return result
            }
        }
    }

    fmt.Println("Report safe immediately", report)
    return true

}

func reportIsSafeWithProblemDampener(report []int) bool{

    for i := 0; i < len(report); i++ {
        var possibleSafeReport = removeElement(report, i)
        // fmt.Println(report)
        // fmt.Println(possibleSafeReport)
        if reportIsSafe(possibleSafeReport, false) == true { return true }
    }

    return false

}

func removeElement(slice []int, i int) []int {
    ret := make([]int, 0)
    ret = append(ret, slice[:i]...)
    return append(ret, slice[i+1:]...)
}

func reportIsSafePart2(report []int) bool {
    if (len(report) == 1) {
        return true
    }

    var partialDescCounter uint
    var partialAscCounter uint
    var partialEqualCounter uint
    var toMuchIncreaseCounter uint

    for i := 0; i < len(report) - 1; i++ {
        if (math.Abs(float64(report[i] - report[i + 1])) <= 3){
            if report[i] > report[i + 1] { 
                partialDescCounter++
            } else if report[i] < report[i + 1]{ 
                partialAscCounter++ 
            } else {
                partialEqualCounter++
            }
        } else { toMuchIncreaseCounter++ }

    }

    fmt.Println("Partial desc:", partialDescCounter)
    fmt.Println("Partial asc:", partialAscCounter)
    fmt.Println("Partial equal:", partialEqualCounter)

    if (partialDescCounter >= 1 && (partialAscCounter + partialEqualCounter + toMuchIncreaseCounter) <= 1) { return true }
    if (partialAscCounter >= 1 && (partialDescCounter + partialEqualCounter + toMuchIncreaseCounter) <= 1) { return true }

    return false
    
}

func arrayOfStringToInteger(input []string) (error, []int) {
    var intArray []int
    
    for _, str := range input {
		// Convert string to integer
		num, err := strconv.Atoi(str)
		if err != nil {
			// Handle error (invalid string to int conversion)
			fmt.Println("Error converting string to int:", err)
			return err, []int{}
		}
		intArray = append(intArray, num)
	}

    return nil, intArray
}