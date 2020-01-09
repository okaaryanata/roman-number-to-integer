package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	console "../console"
	convert "../convert"
)

// Menu - Funtion to create Main Menu
func Menu() {
	console.CallClear()
	fmt.Println("======================================================")
	fmt.Println()
	fmt.Println("Menu : ")
	fmt.Println("1. Input Data")
	fmt.Println("2. Get Data")
	fmt.Println("Q. Quit")
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Print("Enter menu : ")
}

// InputData - menu input / add data
func InputData() {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	console.CallClear()
	fmt.Println("======================================================")
	fmt.Println()
	fmt.Println("INPUT DATA")
	fmt.Println("Press 0 and Enter for back to Main Menu")
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Print("Enter new data : ")
	for input != "0" {
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		arr := strings.Split(input, " ")
		lenArr := len(arr)
		for idx, elm := range arr {
			if idx != lenArr-1 || elm == "Credits" {
				arr[idx] = strings.ToLower(elm)
			}
		}
		switch input {
		case "0":
			Menu()
		default:
			convert.ValidateInputData(arr)
		}
	}
}

// ReadData - menu read data
func ReadData() {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	console.CallClear()
	fmt.Println("======================================================")
	fmt.Println()
	fmt.Println("READ DATA")
	fmt.Println("Press 0 and Enter for back to Main Menu")
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Print("Enter your words : ")
	for input != "0" {
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.ToLower(input)
		arr := strings.Split(input, " ")
		switch input {
		case "0":
			Menu()
		default:
			convert.ValidateReadData(arr)
		}
	}
}

// MainMenu - function to call Menu()
func MainMenu() {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	Menu()
	for input != "Q" && input != "q" {
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		switch input {
		case "1":
			InputData()
		case "2":
			ReadData()
		default:
			if strings.Compare(input, "Q") != 0 && strings.Compare(input, "q") != 0 {
				fmt.Println("Invalid Menu")
				time.Sleep(1 * time.Second)
				Menu()
			}
		}
	}
	fmt.Println("======================================================")
}
