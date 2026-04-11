package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)
var nodes = []string{
	"http://localhost:8001",
	"http://localhost:8002",
	"http://localhost:8003",
}
const chunkSize = 1024 * 1024

func splitAndDistribute(filePath, fileID string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	chunkIndex := 0

	for {
		n, readErr := file.Read(buffer)
		if readErr != nil && readErr != io.EOF {
			return 0, readErr
		}
		if n == 0 {
			break
		}

		node := nodes[chunkIndex%len(nodes)]
		chunkName := fmt.Sprintf("%s_%d.chunk", fileID, chunkIndex)
		sendErr := sendChunkToNode(node, chunkName, buffer[:n])
		if sendErr != nil{
			return 0,sendErr
		}
		chunkIndex++

		if readErr == io.EOF{
			break
		}
	}
	return chunkIndex, nil

}
func fetchChunk(nodeURL, chunkName string)([]byte, error){
	resp, err := http.Get(nodeURL + "/chunk?name=" + chunkName)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s from %s: status %d", chunkName, nodeURL, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
func mergeDistributed(outputFile string , totalChunks int, fileID string) error{
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	for i:= 0 ; i<totalChunks; i++{
		node := nodes[i%len(nodes)]
		chunkName := fmt.Sprintf("%s_%d.chunk", fileID, i)

		data, err := fetchChunk(node, chunkName)
		if err != nil{
			return err
		}
		out.Write(data)
		fmt.Println("Fetched from:", node, chunkName)
	}
	return nil
}
func storeChunk(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only post allowed",http.StatusMethodNotAllowed)
		return
	}

	file , header , err := r.FormFile("chunk")
	if err != nil {
		http.Error(w,"Invalid Check",http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll("data",os.ModePerm)
	path := filepath.Join("data",header.Filename)
	dst, err := os.Create(path)

	if err != nil{
		http.Error(w,"Cannot Save Chunk",http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst,file)
	fmt.Println("Stored:", header.Filename)
	w.Write([]byte("OK"))
}
func getChunk(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == ""{
		http.Error(w,"Missing Name",  http.StatusBadRequest)
		return
	}
	path := filepath.Join("data", filepath.Base(name))
	http.ServeFile(w,r,path)
}
func mergeChunks(outputFile string, totalChunks int) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer outFile.Close()
	for i := 0; i < totalChunks; i++ {
		chunkFileName := fmt.Sprintf("data/file_%d.chunk", i)

		chunkFile, err := os.Open(chunkFileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, chunkFile)
		if err != nil {
			chunkFile.Close()
			return err
		}
		chunkFile.Close()
		fmt.Println("Merged", chunkFileName)
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

	fileID := fmt.Sprintf("%d", time.Now().UnixNano())

	totalChunks, err := splitAndDistribute(dstPath, fileID)

	if err != nil {
		http.Error(w, "Could not split uploaded file", http.StatusInternalServerError)
		return
	}
	err = mergeDistributed("reconstructed.mp4", totalChunks, fileID)
	if err != nil {
		http.Error(w, "Could not reconstruct distributed file", http.StatusInternalServerError)
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

	filePath := "./reconstructed.mp4"
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
}

func sendChunkToNode(nodeURL, fileName string, data[]byte) error{
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("chunk",fileName)
	if err != nil{
		return err
	}
	part.Write(data)
	writer.Close()

	req, err := http.NewRequest("POST",nodeURL+"/store",&body)

	if err != nil{
		return err
	}

	req.Header.Set("Content-Type",writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to store %s on %s: status %d", fileName, nodeURL, resp.StatusCode)
	}

	fmt.Println("Sent to:",nodeURL ,fileName)
	return nil
}
func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/store", storeChunk)
	http.HandleFunc("/chunk", getChunk)
	// http.Handle("/", http.FileServer(http.Dir("./test-client")))
	if err := http.ListenAndServe(":8005", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello")
}
