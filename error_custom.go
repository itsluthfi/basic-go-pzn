package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (v *NotFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Data tidak tervalidasi"}
	}
	if id != "luthfi" {
		return &NotFoundError{"Data tidak ditemukan"}
	}
	return nil
}

func main() {
	err := SaveData("Luthfi", nil)
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Validation error:", validationErr)
		} else if notFoundErr, ok := err.(*NotFoundError); ok {
			fmt.Println("Not found error:", notFoundErr)
		} else {
			fmt.Println("Unknown error:", err.Error())
		}

		// switch
		switch finalErr := err.(type) {
		case *ValidationError:
			fmt.Println("Validation error:", finalErr)
		case *NotFoundError:
			fmt.Println("Not found error:", finalErr)
		default:
			fmt.Println("Unknown error:", finalErr.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
