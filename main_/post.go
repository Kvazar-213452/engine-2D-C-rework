package main_

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func SaveFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	var receivedData interface{}
	err := json.NewDecoder(r.Body).Decode(&receivedData)
	if err != nil {
		http.Error(w, "Помилка при читанні тіла запиту", http.StatusBadRequest)
		return
	}

	filePath := "out/data.json"
	jsonData, err := json.MarshalIndent(receivedData, "", "  ")
	if err != nil {
		http.Error(w, "Помилка при серіалізації даних", http.StatusInternalServerError)
		return
	}

	if _, err := os.Stat("out"); os.IsNotExist(err) {
		err := os.Mkdir("out", 0755)
		if err != nil {
			http.Error(w, "Не вдалося створити каталог", http.StatusInternalServerError)
			return
		}
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		http.Error(w, "Сталася помилка при записі файлу", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Дані успішно збережено у файлі: %s", filePath)
}

var process *exec.Cmd

func StopProject(w http.ResponseWriter, r *http.Request) {
	if startCore == 1 {
		startCore = 0
		if process != nil {
			err := process.Process.Kill()
			if err != nil {
				http.Error(w, "Не вдалося зупинити процес", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Process stopped and reset"))
	} else {
		http.Error(w, "Процес уже зупинено", http.StatusBadRequest)
	}
}

func OpenFile(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("save", "main.json")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Не вдалося прочитати файл", http.StatusInternalServerError)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		http.Error(w, "Помилка при розборі JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonData)
}

var startCore = 0

type RequestData struct {
	TetxY string `json:"tetxY"`
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Помилка при читанні тіла запиту", http.StatusBadRequest)
		return
	}

	err = saveToFile(data.TetxY)
	if err != nil {
		http.Error(w, "Помилка при запису у файл", http.StatusInternalServerError)
		return
	}

	err = runExecutable()
	if err != nil {
		http.Error(w, "Помилка при запуску main.exe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Дані успішно отримано, записано у файл та main.exe запущено"))
}

func saveToFile(content string) error {
	err := os.MkdirAll("engene_prefab/start", os.ModePerm)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("engene_prefab/start/index.html", []byte(content), 0644)
}

func runExecutable() error {
	cmd := exec.Command("cmd", "/C", "code.bat")
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func CompiletePrg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Невірний запит", http.StatusBadRequest)
		return
	}

	outDir := "engene_prefab/comp"
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Не вдалося створити папку", http.StatusInternalServerError)
		return
	}

	sourceFile := "engene_prefab/start/main.exe"
	destinationFile := filepath.Join(outDir, "main.exe")

	err = copyFile(sourceFile, destinationFile)
	if err != nil {
		http.Error(w, "Не вдалося скопіювати файл", http.StatusInternalServerError)
		return
	}

	htmlFile := filepath.Join(outDir, "index.html")
	err = os.WriteFile(htmlFile, []byte(data.TetxY), 0644)
	if err != nil {
		http.Error(w, "Не вдалося створити файл index.html", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Файл index.html успішно створений!"))
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
