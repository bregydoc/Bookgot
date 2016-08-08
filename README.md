# Bookgot
BookGot is a simple boot for farm free books from PACKTPUB.COM

#Install
```
go get -u github.com/bregydoc/Bookgot
```

# Usage

First import BookGot
```go
//...
import "github.com/bregydoc/Bookgot"
```

Second create a User for Login in PackPub

```go
user := Bookgot.CreateNewPBUser("example@email.com", "password")
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

If you want to know the list of books in your account, you can use the GetNamesOfBooks() function

```go
//...
user := Bookgot.CreateNewPBUser("example@email.com","password")
books := user.GetNamesOfBooks()
for i, book := range books{
	fmt.Println("Book ", i , " : ", book)
}

```

# Other functions

Return the name of last free book in PackPub.com

```go
//...
name := Bookgot.GetNameOfCurrentFreeBook()
fmt.Println(name)

```
Return the time left for update the free book

```go
//...
timeLeft, _ := Bookgot.GetTimeForNewFreeBook()
fmt.Println(timeLeft.String())

```


#Example for create a complete bot

```go
package main

import(
	"github.com/bregydoc/Bookgot"
	"fmt"
)

func main() {
	userExample := Bookgot.CreateNewPBUser("example@domain.com", "password")
	if userExample.Login() {
		fmt.Println("Logged!")
	}
	for {
		haveNewBook := userExample.VerifyIfIHaveLastFreeBook()

		if haveNewBook{
			haveNewBook = userExample.VerifyIfIHaveLastFreeBook()
			timeAfter, _ := Bookgot.GetTimeForNewFreeBook()
			fmt.Println("Not new free book, time for next book: ", timeAfter.String())
		}else {
			petition :=  userExample.PullNewFreeBook()
			if petition{
				fmt.Println("New free book added at your library")
			}else{
				fmt.Println("Error, verify your email and password")
			}
		}
		
	}

}
```

