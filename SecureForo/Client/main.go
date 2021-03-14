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

var templateRegister string = `
<html>
	<head>
		<title>
			Register
		</title>
	</head>
	<body>
		<h1>Register Page</h1>
		<form>
			<input type="text" name="name" id="name" placeholder="name"/>
			<input type="password" name="password" id="password" placeholder="password"/>
			<input type="submit" onclick="registerFunc()" />
		</form>
		<button onclick="goToLoginFunc()">Login</button>
	</body>
</html>	
`

var templateLogin string = `
<html>
	<head>
		<title>
			Login
		</title>
	</head>
	<body>
		<h1>Login Page</h1>
		<form>
			<input type="text" name="name" id="name" placeholder="name"/>
			<input type="password" name="password" id="password" placeholder="password"/>
			<input type="submit" onclick="loginFunc()" />
		</form>
		<button onclick="goToRegisterFunc()">Register</button>
	</body>
</html>	
`

func RegisterPage(ui lorca.UI) {
	user := User{}

	ui.Load("data:text/html," + url.PathEscape(templateRegister))

	ui.Bind("registerFunc", func() {
		user.name = ui.Eval(`document.getElementById('name').value`).String()
		user.password = ui.Eval(`document.getElementById('password').value`).String()

		fmt.Println("Register: ", user.name, user.password)
	})

	ui.Bind("goToLoginFunc", func() { LoginPage(ui) })
}

func LoginPage(ui lorca.UI) {
	user := User{}

	ui.Load("data:text/html," + url.PathEscape(templateLogin))

	ui.Bind("loginFunc", func() {
		user.name = ui.Eval(`document.getElementById('name').value`).String()
		user.password = ui.Eval(`document.getElementById('password').value`).String()

		fmt.Println("Login: ", user.name, user.password)
	})

	ui.Bind("goToRegisterFunc", func() { RegisterPage(ui) })
}

func main() {
	ui, err := lorca.New("", "", 500, 400)

	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()

	ui.Bind("start", func() { log.Print("UI is ready") })

	LoginPage(ui)

	<-ui.Done()
}
