package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Username  string   `json:"username"`
	Password  string   `json:"-"`
	Following []string `json:"following"`
	// omitempty neglect the empty field of the data
	Followers []string `json:"followers,omitempty"`
}

func main() {
	// encodeJsonData()
	decodeJsonData()
}

func encodeJsonData() {
	users := []user{
		{"Aman", "aman@example.com", 23, "aman@123", "123@aman", []string{"raj", "shubham", "amit"}, []string{"shlok", "Ravi", "Arvind"}},
		{"Sohan", "shlok@example.com", 23, "shlok@123", "123@shlok", []string{"raj", "shubham", "amit"}, []string{"shlok", "Ravi", "Arvind"}},
		{"Ravi", "ravi@example.com", 23, "ravi@123", "123@ravi", []string{"raj", "shubham", "amit"}, nil},
	}

	// convert the users data into json data

	// finalJson,err:=json.Marshal(users)
	finalJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func decodeJsonData() {
	jsonData := []byte(`{
	"name": "Sohan",
                "email": "shlok@example.com",
                "age": 23,
                "username": "shlok@123",
                "following": [
                        "raj",
                        "shubham",
                        "amit"
                ],
                "followers": [
                        "shlok",
                        "Ravi",
                        "Arvind"
                ]
	}`)

	var jsonUser user

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("\n Json was valid!\n")
		json.Unmarshal(jsonData, &jsonUser)
		fmt.Printf("%#v\n", jsonUser)
	} else {
		fmt.Println("Json was invalid!\n")
	}

	// for key value pairs
	// if we don't know what kind of data it is so we just populat it like this

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)

	fmt.Println("\n", myData)

	for key, val := range myData {
		fmt.Println(key, " : ", val)
	}
}
