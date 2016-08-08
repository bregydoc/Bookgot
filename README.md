# Bookgot
Simple script for free books farm

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
