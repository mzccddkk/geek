package main

import "github.com/gin-gonic/gin"

func main() {
	s, err := initApp()
	if err != nil {
		panic(err)
	}

	e := gin.Default()
	s.register(e)
}