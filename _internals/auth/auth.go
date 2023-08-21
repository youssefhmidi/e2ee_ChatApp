package auth

import "strings"

// the return 'Name' value is the type of the given token it can be barear, bot, or anything that you wan't to
// add to your app
//
// it also returns the JWT token
func GetAuthType(RequestedHeader string) (Name string, jwt string) {
	seperated := strings.Split(RequestedHeader, " ")
	Name, jwt = seperated[0], seperated[1]
	return
}

// the 'Name' parameter is the returned string from the GetAuthType() function
//
// it returns true if the type (i.e the string) is the same as "Bearer" otherwise it returns
func IsBearer(Name string) bool {
	return Name == "Bearer"
}
