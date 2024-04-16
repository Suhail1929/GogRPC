// package main

// import (
// 	// "Souhail-Bakri/Go-gRPC/device"
// 	// "Souhail-Bakri/Go-gRPC/operations"
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {

// 	Welcome()

// }

// func Welcome() {
// 	fmt.Println("Bienvenue dans le programme de gestion d'appareils des clients")
// 	fmt.Println("Veuillez choisir une option:")
// 	fmt.Println("1. Voir les données des appareils")
// 	fmt.Println("2. Stocker les données des appareils dans la base de données")
// 	fmt.Println("3. Quitter")
// 	ChooseOption()
// }

// func ChooseOption() {
// 	var option int
// 	for {
// 		fmt.Scanln(&option)
// 		switch option {
// 		case 1:
// 			ShowDevices()
// 		case 2:
// 			StoreDevices()
// 		case 3:
// 			fmt.Println("Merci d'avoir utilisé notre programme")
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Veuillez choisir une option valide")
// 		}
// 	}
// }

// func ChooseJournnee() string {
// 	var journee string
// 	for {
// 		fmt.Println("Veuillez choisir une journée entre 1 et 5")
// 		fmt.Scanln(&journee)
// 		if journee >= "1" && journee <= "5" {
// 			break
// 		} else {
// 			fmt.Println("Veuillez choisir une journée entre 1 et 5")
// 		}
// 	}
// 	return journee
// }

// func OpenJsonFile() *os.File {
// 	Choice := ChooseJournnee()
// 	file, err := os.Open("data/journee_" + Choice + ".json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return file
// }

// func ShowDevices() {
// 	file := OpenJsonFile()
// 	defer file.Close() // Ensure file is closed after use

// 	// read lines from the opened file
// 	lines, err := readLines(file)
// 	if err != nil {
// 		log.Fatalf("readLines: %s", err)
// 	}

// 	// print file contents
// 	for i, line := range lines {
// 		fmt.Println(i, line)
// 	}
// }

// func StoreDevices() {

// }

// // read line by line into memory
// // all file contents is stores in lines[]
// func readLines(file *os.File) ([]string, error) {
// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return lines, nil
// }

// func countOperations() {
	
// }
