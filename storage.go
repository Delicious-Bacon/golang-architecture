package architecture

//User is the mock to save into fake DB
type User struct {
	First string
}

//Accessor saves and retrieves users to fake DB
type Accessor interface {
	Save(n int, u User)
	Retrieve(n int) User
}

//Put saves a user into the fake DB
func Put(a Accessor, n int, u User) {
	a.Save(n, u)
}

//Get retrieves a user from the fake DB
func Get(a Accessor, n int) User {
	return a.Retrieve(n)
}
