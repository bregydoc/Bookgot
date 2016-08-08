# Bookgot
BookGot is a simple boot for farm free books from PACKPUB.COM
```go
	func (user *PackPubUser) VerifyIfIHaveLastFreeBook() bool {
		myBooks := user.GetNamesOfBooks()
		actualFreeBook := GetNameOfCurrentFreeBook()
		for _, book := range myBooks {

			comp := strings.Contains(book, actualFreeBook)
			if comp {
				fmt.Println(book, " is ", actualFreeBook)
				return true
			}

		}
		return false
	}

```
# Usage
