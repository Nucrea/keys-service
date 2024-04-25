package main

import (
	"context"
	"keys_service/app"
)

func main() {
	ctx := context.Background()

	a := app.App{}
	a.Run(ctx)
}
