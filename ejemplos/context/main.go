package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	newCtx := addValue(ctx)

	s := newCtx.Value("llave")

	fmt.Println(s)

}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "llave", "valor")
}
