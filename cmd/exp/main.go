package main

import (
	"errors"
	"fmt"
)

func main() {
	err := B()
	// TODO: Determine if the `err` variable is an `ErrNotFound`
	if errors.Is(err, ErrNotFound) {
		fmt.Println("It was ErrNotFound")
	} else {
		fmt.Println("Not it was not.")
	}
}

var ErrNotFound = errors.New("not found")

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
