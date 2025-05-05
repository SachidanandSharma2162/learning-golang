package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Sample data structure to send/receive
type PostData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Step 1: Start the server in background
	go startServer()

	// Step 2: Wait for server to start
	time.Sleep(1 * time.Second)

	// Step 3: Send POST request
	PerformPostRequest()
}

func startServer() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}

		var data PostData
		json.Unmarshal(body, &data)

		fmt.Println("🟢 Received POST data:", data)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Data received successfully!",
		})
	})

	fmt.Println("🟢 Server running on http://127.0.0.1:3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}

func PerformPostRequest() {
	url := "http://127.0.0.1:3000/post"

	// Step 1: Prepare the JSON data
	data := PostData{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Step 2: Send the POST request
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Step 3: Read the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Server Response:")
	fmt.Println(string(body))
}
