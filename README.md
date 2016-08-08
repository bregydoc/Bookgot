# Bookgot
BookGot is a simple boot for farm free books from PACKPUB.COM

# Usage

## How use Bookgot:
First create a User for Login in PackPub

```go
user := CreateNewPBUser("username@email.com", "password")
ok := user.Login()
if ok && user.Logged{
	fmt.Println("User logged!")
}

```
