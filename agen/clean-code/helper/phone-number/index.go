package phonenumber

import "github.com/ttacon/libphonenumber"

// Parse returns PhoneNumber && Country code if input string is valid PhoneNumber
func Parse(phoneNumber string) (num *libphonenumber.PhoneNumber, err error) {
	return libphonenumber.Parse(phoneNumber, "")
}
