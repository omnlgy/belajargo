package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	website := map[string]string{
		"google":   "www.google.com",
		"youtube":  "www.youtube.com",
		"facebook": "www.facebook.com",
	}

	fmt.Println(website)
	delete(website, "facebook")
	fmt.Println(website)
	website["twitter"] = "www.twitter.com"
	fmt.Println(website)

	makeSlice := make([]string, 3, 3)

	makeSlice = append(makeSlice, "string 1", "string 2")

	fmt.Println(makeSlice)
	makeSlice[0] = "string 4"
	makeSlice[1] = "string 3"
	makeSlice[2] = "string 3"
	fmt.Println(makeSlice)

	for key, value := range website {
		fmt.Println(key, value)
	}

	for index, value := range makeSlice {
		fmt.Println(index, value)
	}

}
