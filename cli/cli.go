package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

)

func Start(db *sql.DB) {
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
		case "1":
			//handler.Register(db)
		case "2":
			//handler.Login(db)
		case "3":
			fmt.Println("Good bye")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}