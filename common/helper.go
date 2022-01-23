package helper

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool, bool) {
	// could use some appropriate validation package here
	isValidFirstName := len(firstName) > 2
	isValidLastName := len(lastName) > 2
	isValidEmail := len(email) > 2

	return isValidFirstName, isValidLastName, isValidEmail
}
