package main

import (
	"context"
	"keys_service/app"
)

func main() {
	a := app.App{}
	a.Run(context.Background())
}
