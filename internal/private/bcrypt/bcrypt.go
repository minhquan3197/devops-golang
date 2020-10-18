package bcrypt

import "golang.org/x/crypto/bcrypt"

// HashPassword func encrpt password
func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashedPassword)
}

// ComparePassword func compare two password
func ComparePassword(givenPwd string, storedPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPwd), []byte(givenPwd)); err != nil {
		return false
	}
	return true
}
