package strings

import (
	"crypto/rand"
	"encoding/base64"
	"regexp"
	"strings"
)

// Generate random bytes
func RandomBytes(size int) (rb []byte, err error) {
	rb = make([]byte, size)
	_, err = rand.Read(rb)
	return
}

// Generate random string
func RandomString(length int) string {
	result := ""
	for len(result) < length {
		size := length - len(result)
		randBytes, _ := RandomBytes(size)
		encoded := base64.StdEncoding.EncodeToString(randBytes)
		encoded = strings.ReplaceAll(encoded, "/", "")
		encoded = strings.ReplaceAll(encoded, "+", "")
		encoded = strings.ReplaceAll(encoded, "=", "")
		encoded = encoded[0:size]
		result = result + encoded
	}

	return result
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// ToTitleNorm ...
func ToTitleNorm(input string) string {
	var output []byte
	var upperCount int
	for i, c := range input {
		switch {
		case c >= 'A' && c <= 'Z':
			if upperCount == 0 || nextIsLower(input, i) {
				output = append(output, byte(c))
			} else {
				output = append(output, byte(c-'A'+'a'))
			}
			upperCount++

		case c >= 'a' && c <= 'z':
			output = append(output, byte(c))
			upperCount = 0

		case c >= '0' && c <= '9':
			if i == 0 {
				panic("common/str: Identifier must start with a character: `" + input + "`")
			}
			output = append(output, byte(c))
			upperCount = 0
		}
	}
	return string(output)
}

// ToSnake ...
func ToSnake(input string) string {
	var output []byte
	var upperCount int
	for i, c := range input {
		switch {
		case c >= 'A' && c <= 'Z':
			if i > 0 && (upperCount == 0 || nextIsLower(input, i)) {
				output = append(output, '_')
			}
			output = append(output, byte(c-'A'+'a'))
			upperCount++

		case c >= 'a' && c <= 'z':
			output = append(output, byte(c))
			upperCount = 0

		case c >= '0' && c <= '9':
			if i == 0 {
				panic("common/str: Identifier must start with a character: `" + input + "`")
			}
			output = append(output, byte(c))
			// prevIsLower = true

		default:
			panic("common/str: Invalid identifier: `" + input + "`")
		}
	}
	return string(output)
}

// The next character is lower case, but not the last 's'.
//
//     HTMLFile -> html_file
//     URLs     -> urls
func nextIsLower(input string, i int) bool {
	i++
	if i >= len(input) {
		return false
	}
	c := input[i]
	if c == 's' && i == len(input)-1 {
		return false
	}
	return c >= 'a' && c <= 'z'
}

var uppercaseAcronym = map[string]bool{
	"ID": true,
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if uppercaseAcronym[s] {
		s = strings.ToLower(s)
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

func CensorString(str string) string {
	if len(str) <= 6 {
		return "***"
	}

	return str[:2] + "***" + str[len(str)-2:]
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func NewReplaceByArray(olds []string, new string, s string) string {
	oldReplace := make([]string, 0)
	for _, v := range olds {
		oldReplace = append(oldReplace, v, new)
	}
	return strings.NewReplacer(oldReplace...).Replace(s)
}
