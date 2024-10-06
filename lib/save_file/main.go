package main

import (
	"C"
	"encoding/json"
	"io/ioutil"
	"log"
)

//export SaveFile
func SaveFile(receivedData string) string {
    // Конвертація рядка в мапу
    var data map[string]interface{}
    if err := json.Unmarshal([]byte(receivedData), &data); err != nil {
        return "Invalid JSON"
    }

    // Конвертація отриманих даних у JSON формат
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return "Error marshaling JSON"
    }

    // Вказуємо шлях до файлу
    filePath := "out/data.json"
    
    // Записуємо JSON у файл
    if err := ioutil.WriteFile(filePath, jsonData, 0644); err != nil {
        log.Println("Error writing file:", err)
        return "Error writing file"
    }

    log.Println("File created successfully and data written.")
    return "Process stopped and reset"
}

func main() {}
