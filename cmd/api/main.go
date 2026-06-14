package main

import (
	"github.com/Habeebamoo/tunnl-backend/internal/app"
)

func main() {
	application := app.New()
	
	if err := application.Run(); err != nil {
		panic(err)
	}
}