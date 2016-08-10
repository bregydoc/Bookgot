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
