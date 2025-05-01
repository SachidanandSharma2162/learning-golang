package main

import ("fmt"
        "time"
	)

func main()  {
	fmt.Println("Sleeping for 3 seconds...")
	// time.Sleep(3 * time.Second)
    now := time.Now()
	fmt.Println(now)
    fmt.Println("Year:", now.Year())
    fmt.Println("Month:", now.Month())
    fmt.Println("Day:", now.Day())
    fmt.Println("Hour:", now.Hour())
    fmt.Println("Minute:", now.Minute())
    fmt.Println("Second:", now.Second())
	fmt.Println("Awake now!")
	formatted := now.Format("01/02/2006 15:04:05 Wednesday")
	fmt.Println(formatted)
}