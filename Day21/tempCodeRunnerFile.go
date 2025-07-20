func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)
	go increment(&wg,"first",&mu)
	go increment(&wg,"second",&mu)

	wg.Wait()
	fmt.Println(count)
}