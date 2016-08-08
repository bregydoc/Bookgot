# Bookgot
BookGot is a simple boot for farm free books from PACKTPUB.COM

# Usage

First create a User for Login in PackPub

```go
user := CreateNewPBUser("username@email.com", "password")
ok := user.Login()
if ok && user.Logged{
	fmt.Println("User logged!")
}

```

For verify if the user have the last free book from PacktPub.com

```go
//...
if ok := user.VerifyIfIHaveLastFreeBook(); ok{
	fmt.Println("You are up to day")
}
```

If VerifyIfIHaveLastFreeBook() return false is because PackPub have a new free book and you not claim this.
You can know the name of the current free book

```go
//...
if ok := user.PullNewFreeBook(); ok{
	fmt.Println("Ok, you have the last free book")
}

```
