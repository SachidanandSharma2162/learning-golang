package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	go startServer()
	time.Sleep(1 * time.Second) // ✅ Wait for server to start
	fmt.Println("Lets learn form data post!")
	PerformPostFormDataSubmit()
}

func startServer() {
	http.HandleFunc("/formpost", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")

		// Respond with received data
		fmt.Fprintf(w, "Received:\nFirst Name: %s\nLast Name: %s\nEmail: %s\n", firstName, lastName, email)
	})

	fmt.Println("🟢 Server running on http://127.0.0.1:3000")
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		panic(err)
	}
}

func PerformPostFormDataSubmit() {
	myurl := "http://localhost:3000/formpost"

	data := url.Values{}
	data.Add("firstName", "Hanuman")
	data.Add("lastName", "Sharma")
	data.Add("email", "email@gmail.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("✅ Server Response:")
	fmt.Println(string(content))
}
