package main
import (
	"errors"
	"fmt"
	"log"
)

func division(x,y int)(int, error, error){
	if y == 0 {
		return 0, nil, errors.New("Can not divide by zero")
	}
	if x%y !=0 {
		reminder := errors.New("This division has reminder!...")
		return x / y, remainder, nil
	}else{
		return x/y, nil, nil
	}

}
