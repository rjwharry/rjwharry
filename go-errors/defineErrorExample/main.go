package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() functions")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("End normally")
	} else {
		fmt.Println(err)
	} //"End normally" 출력

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("End normally")
	} else {
		fmt.Println(err)
	} //"Error in returnError() function" 출력
}
