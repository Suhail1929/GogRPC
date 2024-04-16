// package main

// import (
// 	// "Souhail-Bakri/Go-gRPC/device"
// 	// "Souhail-Bakri/Go-gRPC/operations"
// 	"Bakri-Souhail/GoGrpcClient/data"
// 	"fmt"
// 	"log"
// 	"os"
// )

// // func main() {

// // 	Welcome()

// // }

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

// func OpenJsonFile() (*os.File, error) {
// 	Choice := ChooseJournnee()
// 	filePath := "data/journee_" + Choice + ".json"
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }


// func ShowDevices() {
// 	file, err := OpenJsonFile()
// 	if err != nil {
// 		log.Fatalf("Error opening JSON file: %v", err)
// 	}
// 	defer file.Close()

// 	// Read the content of the JSON file
// 	fileData, err := data.ReadFiles(file)
// 	if err != nil {
// 		log.Fatalf("Error reading JSON file: %v", err)
// 	}

// 	// Now you can access the JSON content as a string using fileData.Content
// 	fmt.Println("File content:")
// 	fmt.Println(fileData.Json)
// }

// func StoreDevices() {

// }



