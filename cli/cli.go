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
	godotenv.Load(".env")

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
		// Register
		case "1":
			fmt.Print("Full Name: ")
			fullName, _ := reader.ReadString('\n')
			fullName = strings.TrimSpace(fullName)

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			fmt.Println("Role:")
			fmt.Println("1. Admin")
			fmt.Println("2. User")
			fmt.Print("Choose role (1/2): ")

			roleStr, _ := reader.ReadString('\n')
			roleStr = strings.TrimSpace(roleStr)

			roleID := 2
			if roleStr == "1" {
				roleID = 1
			}

			u, err := handler.RegisterUser(db, fullName, password, roleID)
			if err != nil {
				fmt.Println("Register gagal:", err)
				continue
			}

			fmt.Println("\nRegister berhasil!")
			fmt.Println("User ID:", u.UserId)

		// Login
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

			fmt.Println("\nLogin berhasil. Welcome,", u.Fullname)
			switch u.RoleID {
			case 1:
				admin.Menu(reader, db, u)
			case 2:
				user.Menu(reader, db, u)

			default:
				fmt.Println("Role tidak dikenal")
			}

		case "3":
			fmt.Println("Good bye")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
