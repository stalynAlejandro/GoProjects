package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

type User struct {
	name, password string
}

var loginPage string = `
<html>
	<head>
		<title>
			Login Page
		</title>
	</head>
	<body>
		<h1>Login Page</h1>
		<form>
			<input type="text" name="name" id="name"/>
			<input type="text" name="password" id="password"/>
			<input type="submit" onclick="loginFunc()" />
		</form>
	</body>
</html>	
`

func LoginPage() {
	fmt.Println("helo")
	ui, err := lorca.New("data:text/html, "+url.PathEscape(loginPage), "", 500, 400)

	user := User{}

	ui.Bind("loginFunc", func() {
		user.name = ui.Eval(`document.getElementById('name').value`).String()
		user.password = ui.Eval(`document.getElementById('password').value`).String()

		fmt.Println(user.name, user.password)
	})

	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()

	//Wait unit Page window is closed
	<-ui.Done()
}

func RegisterPage() {

}

func main() {
	LoginPage()
}
