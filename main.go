package main

import (
	"crypto/md5"
	"fmt"
	"koda-b8-Golang5/data"
	"os"
)

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", hash)
}
type Auth interface {
	Register()
	Login()
	ForgotPassword()
	Exit()
}

type authService struct {
	Users []data.User
}

func (auth *authService) Register() {
	defer fmt.Println("Leaving Register")

	var firstName string
	var lastName string
	var email string
	var password string
	var confirmPassword string
	var confirm string

	fmt.Println("---- Register ----")

	fmt.Print("What is your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email : ")
	fmt.Scan(&email)

	for _, user := range auth.Users {
		if user.Email == email {
			fmt.Println("Email already registered.")
			return
		}
	}

	for {
		fmt.Print("Enter a strong password : ")
		fmt.Scan(&password)

		fmt.Print("Confirm your password : ")
		fmt.Scan(&confirmPassword)

		if password == confirmPassword {
			break
		}

		fmt.Println("Password doesn't match, please try again.")
	}

	fmt.Println()

	fmt.Println("First Name :", firstName)
	fmt.Println("Last Name  :", lastName)
	fmt.Println("Email      :", email)

	fmt.Print("Are you sure? (Y/n) : ")
	fmt.Scan(&confirm)

	if confirm == "Y" || confirm == "y" {

		user := data.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  hashPassword(password),
		}

		auth.Users = append(auth.Users, user)

		fmt.Println("Registration Success")
		return
	}

	fmt.Println("Registration canceled")
}

func (auth *authService) Login() {

	var email string
	var password string

	fmt.Println("---- Login ----")
	fmt.Println()

	fmt.Print("Enter your email : ")
	fmt.Scan(&email)

	fmt.Print("Enter your password : ")
	fmt.Scan(&password)

	fmt.Println()

	for _, user := range auth.Users {

		if user.Email == email {

			if user.Password == hashPassword(password) {

				fmt.Println("Login Successful!")
				auth.Menu(user)
				return
			}

			fmt.Println("Wrong password")
			return
		}
	}

	fmt.Println("Email not found")
}

func (auth *authService) ForgotPassword() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	defer fmt.Println("Leaving Forgot Password")

	var email string

	fmt.Println("---- Forgot Password ----")
	fmt.Println()

	fmt.Print("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println()

	if email == "" {
		panic("Email cannot be empty")
	}

	for i, user := range auth.Users {

		if user.Email == email {

			var newPassword string
			var confirmPassword string

			for {

				fmt.Print("Enter new password : ")
				fmt.Scan(&newPassword)

				fmt.Print("Confirm new password : ")
				fmt.Scan(&confirmPassword)

				if newPassword == confirmPassword {
					break
				}

				fmt.Println("Password doesn't match!")
			}

			auth.Users[i].Password = hashPassword(newPassword)

			fmt.Println("Password updated successfully!")
			return
		}
	}

	fmt.Println("Email not found!")
}

func (auth *authService) ListUsers() {

	fmt.Println("--- List All Users ---")

	for i, user := range auth.Users {

		fmt.Printf("\nUser %d\n", i+1)
		fmt.Println("Full Name :", user.FirstName, user.LastName)
		fmt.Println("Email     :", user.Email)
		fmt.Println("Password  :", user.Password)
	}

	fmt.Println()
	fmt.Print("Press Enter to continue...")

	fmt.Scanln()
	fmt.Scanln()
}

func (auth *authService) Logout() {
	fmt.Println("Logout Success")
}

func (auth *authService) Menu(user data.User) {

	for {

		fmt.Println("--- Welcome to System ---")
		fmt.Println()

		fmt.Printf("Hello %s\n\n", user.Email)

		fmt.Println("1. List All Users")
		fmt.Println("2. Logout")
		fmt.Println()
		fmt.Println("0. Exit")

		var selectOpt int

		fmt.Print("Choose a menu : ")

		_, err := fmt.Scan(&selectOpt)

		if err != nil {

			fmt.Println("The input must be a number!")
			fmt.Scanln()
			continue
		}

		switch selectOpt {

		case 1:
			auth.ListUsers()

		case 2:
			auth.Logout()
			return

		case 0:
			auth.Exit()

		default:
			fmt.Println("Menu not found")
		}
	}
}

func (auth *authService) Exit() {
	fmt.Println("GoodBye")
	os.Exit(0)
}

func main() {

	service := &authService{}

	var auth Auth = service

	for {

		fmt.Println("------ Welcome to System ------")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Forgot Password")
		fmt.Println()
		fmt.Println("0. Exit")

		var menu int

		fmt.Print("Choose Menu : ")
		_, err := fmt.Scan(&menu)
		if err != nil {
			fmt.Println("Input must be number")
			fmt.Scanln()
			continue
		}

		switch menu {

		case 1:
			auth.Register()

		case 2:
			auth.Login()

		case 3:
			auth.ForgotPassword()

		case 0:
			auth.Exit()

		default:
			fmt.Println("Menu not found.")
		}

		fmt.Println()
	}
}