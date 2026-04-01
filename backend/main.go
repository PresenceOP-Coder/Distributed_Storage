package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const chunkSize = 1024 * 1024

func splitfile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	chunkIndex := 1

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		chunkFileName := fmt.Sprintf("data/file_%d.chunk", chunkIndex)
		chunkFile, err := os.Create(chunkFileName)

		if err != nil {
			return err
		}

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			chunkFile.Close()
			return err
		}

		chunkFile.Close()
		fmt.Println("Created: ", chunkFileName)
		chunkIndex++

		if err == io.EOF {
			break
		}
	}
	return nil

}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//*
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//*/
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("myfile")

	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll("./upload", os.ModePerm)
	dstPath := filepath.Join("./upload", header.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Could not create destination file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Could not save uploaded file", http.StatusInternalServerError)
		return
	}

	if err := os.MkdirAll("./data", os.ModePerm); err != nil {
		http.Error(w, "Could not prepare chunk directory", http.StatusInternalServerError)
		return
	}

	if err := splitfile(dstPath); err != nil {
		http.Error(w, "Could not split uploaded file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Uploaded : %s", header.Filename)

}
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")

	if fileName == "" {
		http.Error(w, "Plz give a file name", http.StatusBadRequest)
		return
	}

	fileDir := "./storage"
	filePath := filepath.Join(fileDir, filepath.Base(fileName))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}
func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	// http.Handle("/", http.FileServer(http.Dir("./test-client")))
	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello")
}
