package error_handler

import "fmt"

func HandleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
