package main 

import (
   "fmt"
   "os"
   "os/user"
   "github.com/sttk/stringcase"
)

func main() {

	//Retrieve currently logged in user name
	//The current func returns username if successful otherwise returns err handler
	//golang functions can return multiple values
	u, err := user.Current()

	//in case the current function successful retrieved then the err variable will be null(go lang nil)
	if err != nil {
	   fmt.Println("Cannot get current user: ", err )
	   os.Exit(1)
	}

	fmt.Printf("Hello %s, welcome !\n", stringcase.PascalCase(u.Username) )

}
