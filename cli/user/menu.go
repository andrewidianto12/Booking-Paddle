package user

import (
	"bufio"
	"database/sql"
	"fmt"
	"strings"

	"github.com/andrewidianto/Paddle-Booking/entity"
	"github.com/andrewidianto/Paddle-Booking/handler"
)

func Menu(reader *bufio.Reader, db *sql.DB, user *entity.LoginUser) {

	for {
		fmt.Println("\n=== USER MENU ===")
		fmt.Println("1. View Courts")
		fmt.Println("2. Book Court")
		fmt.Println("3. My Bookings")
		fmt.Println("4. Logout")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("View Courts ")
			handler.ViewCourts(db)

		case "2":
			fmt.Println("Book Court ")
			handler.CreateBooking(db, user)

		case "3":
			fmt.Println("My Bookings ")
			handler.ViewMyBookings(db, user)

		case "4":
			fmt.Println("Logout berhasil")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
