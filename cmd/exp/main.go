package main

import (
	"html/template"
	"os"
	"time"
)

type User struct {
	Name    string
	Age     int
	Bio     string
	Meta    UserMeta
	Contact int
	DOB     time.Time
	Langs   []string
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Gopal Agrawal",
		Meta: UserMeta{
			Visits: 10,
		},
		Bio:     `<script>alert("Yo");</script>`,
		Contact: 128493492349,
		DOB:     time.Date(2000, 2, 30, 0, 0, 0, 0, time.Local),
		Langs:   []string{"Hindi", "English"},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
