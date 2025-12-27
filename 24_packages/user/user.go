package user

// this will be used outside of the package
type User struct {
	Email string // if we want to access these fields along with User then we need to make it start with Capital letter because lowercase will make it limited to package access only
	Name  string
	phoneNumber int
}
