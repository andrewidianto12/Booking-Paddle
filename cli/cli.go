package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andrewidianto/Paddle-Booking/cli/admin"
	"github.com/andrewidianto/Paddle-Booking/cli/user"
	"github.com/andrewidianto/Paddle-Booking/config"
	"github.com/andrewidianto/Paddle-Booking/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	db := config.InitDatabase(os.Getenv("DB_DSN"))
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== PADEL BOOKING CLI ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "2":
			fmt.Print("Full Name: ")
			fullName, _ := reader.ReadString('\n')
			fullName = strings.TrimSpace(fullName)

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			u, err := handler.LoginUser(db, fullName, password)
			if err != nil {
				fmt.Println("Login gagal:", err)
				continue
			}

			if strings.ToLower(u.Fullname) == "admin" {
				admin.Menu(reader, u)
			} else {
				user.Menu(reader, u)
			}

		case "3":
			fmt.Println("Good bye")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
