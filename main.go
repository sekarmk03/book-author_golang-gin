package main

import "book-author/routers"

func main() {
	r := routers.StartApp()
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
