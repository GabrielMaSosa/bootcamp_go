package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	newct := Addcontext(ctx)
	s := newct.Value("Personas")
	fmt.Println(s)
}

func Addcontext(c context.Context) context.Context {
	return context.WithValue(c, "Personas", "Gabriel")

}
