package main

import (
	"context"
	"keys_service/app"
	"os"
)

func main() {
	a := app.App{}
	a.Run(context.Background(), os.Args)
}
