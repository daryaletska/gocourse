package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase - used for naming of types in Go
	// Ex. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, Interfaces, Enums

	// snake_case - commonly used for naming of variables, constants, file names in Go
	// Ex. user_id, user_name, http_request
	// variables, constants, file names

	// UPPERCASE - used only for naming of constants in Go
	// Ex. PI, MAX_VALUE, HTTP_STATUS_OK
	// constants

	// mixedCase - can be used to name variables or certain identifiers
	// especially when dealing with external libraries
	// Ex. javaScript, htmlDocument, isValid

	const MAX_RETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID:", employeeID)
}
