package main

import (
	"errors"
	"fmt"
	"net/mail"
)

func isMailAdressValid(email string) (string, bool) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		err := fmt.Errorf("you provided %s, valid format is xyz@somemail.com", email)
		return err.Error(), false
	}
	return addr.Address, true
}

func intDiv(num, denom int) (int, error) {
	if denom == 0 {
		return -999, errors.New("Divide by zero is not allowed")
	}
	return num / denom, nil
}

//func main() {
//	em, ok := isMailAdressValid("mobarak@appscode.com")
//	if ok {
//		fmt.Println(em)
//	} else {
//		fmt.Println("Wrong ", em)
//	}
//
//	result, err := intDiv(20, 0)
//	if err == nil {
//		fmt.Println("Result : ", result)
//	} else {
//		fmt.Println(err.Error())
//	}
//}
