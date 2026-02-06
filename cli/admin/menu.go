package admin

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/andrewidianto/Paddle-Booking/entity"
	"github.com/andrewidianto/Paddle-Booking/handler"
)

func Menu(reader *bufio.Reader, db *sql.DB, user *entity.LoginUser) {
	for {
		fmt.Println("\n==== ADMIN MENU ====")
		fmt.Println("1. Manage Schedule")
		fmt.Println("2. Manage User")
		fmt.Println("3. Reports")
		fmt.Println("4. Logout")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			ManageSchedule(db)
		case "2":
			ManageUser(db)
		case "3":
			reportMenu(db)
		case "4":
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func ManageSchedule(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== ADD SCHEDULE ===")

		fmt.Println("1. Add Court")
		fmt.Println("2. Update Court")
		fmt.Println("3. Add Time Slot")
		fmt.Println("4. Update Time Slot")
		fmt.Println("0. Back")
		fmt.Print("Pilih: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handler.AddCourt(db)
		case "2":
			handler.UpdateCourt(db)
		case "3":
			handler.AddTimeSlot(db)
		case "4":
			handler.UpdateTimeSlot(db)
		case "0":
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func ManageUser(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MANAGE USERS ===")
		fmt.Println("1. Add User")
		fmt.Println("2. Update User")
		fmt.Println("0. Back")
		fmt.Print("Choose: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handler.AddUser(db)
		case "2":
			handler.UpdateUser(db)
		case "0":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func reportMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n==== REPORT MENU ====")
		fmt.Println("1. User Report")
		fmt.Println("2. Booking Report")
		fmt.Println("3. Daily Revenue Report")
		fmt.Println("4. Monthly Revenue Report")
		fmt.Println("0. Back")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			handler.UserReport(db)
		case "2":
			handler.BookingReport(db)
		case "3":
			handler.DailyRevenue(db)
		case "4":
			handler.MonthlyRevenue(db)
		case "0":
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
