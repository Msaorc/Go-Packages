package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "jwt", "asdfasdfasdfjklçjklçjklç")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	jwt := ctx.Value("jwt")
	fmt.Println(jwt)
}