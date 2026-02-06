package handler

import (
	"database/sql"
	"errors"
	"bufio"
	"fmt"
	"os"
	"strconv" 
	"strings"
	"time"
	"github.com/andrewidianto/Paddle-Booking/entity"
)

func RegisterUser(db *sql.DB, fullname, password string, roleID int) (*entity.RegisterUser, error) {
	var user entity.RegisterUser

	result, err := db.Exec(`
		INSERT INTO users (full_name, password, role_id)
		VALUES (?, ?, ?)
	`, fullname, password, roleID)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.UserId = int(lastID)
	user.Fullname = fullname
	user.Password = password
	user.RoleID = roleID

	return &user, nil
}

func LoginUser(db *sql.DB, fullName, password string) (*entity.LoginUser, error) {
	var user entity.LoginUser

	err := db.QueryRow(`
		SELECT user_id, full_name, password, role_id
		FROM users
		WHERE full_name = ?
		LIMIT 1
	`, fullName).Scan(
		&user.UserId,
		&user.Fullname,
		&user.Password,
		&user.RoleID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func ViewCourts(db *sql.DB) {
	rows, err := db.Query(`
		SELECT court_id, court_name, location, price_per_hour, status
		FROM courts
		WHERE status = 'AVAILABLE'
	`)
	if err != nil {
		fmt.Println("Error fetching courts:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== AVAILABLE COURTS ===")
	for rows.Next() {
		var id int
		var name, location, status string
		var price float64

		if err := rows.Scan(&id, &name, &location, &price, &status); err != nil {
			fmt.Println("Scan error:", err)
			return
		}

		fmt.Printf("ID: %d | %s | %s | Rp%.2f/hour\n",
			id, name, location, price)
	}
}

func CreateBooking(db *sql.DB, user *entity.LoginUser) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Court ID: ")
	courtIDStr, _ := reader.ReadString('\n')

	fmt.Print("Time Slot ID: ")
	timeSlotIDStr, _ := reader.ReadString('\n')

	fmt.Print("Booking Date (YYYY-MM-DD): ")
	dateStr, _ := reader.ReadString('\n')

	courtIDStr = strings.TrimSpace(courtIDStr)
	timeSlotIDStr = strings.TrimSpace(timeSlotIDStr)
	dateStr = strings.TrimSpace(dateStr)
	courtID, err := strconv.Atoi(courtIDStr)
	if err != nil {
		fmt.Println("Court ID harus angka")
		return
	}

	timeSlotID, err := strconv.Atoi(timeSlotIDStr)
	if err != nil {
		fmt.Println("Time Slot ID harus angka")
		return
	}

	bookingDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Format tanggal salah")
		return
	}

	_, err = db.Exec(`
		INSERT INTO bookings (user_id, court_id, booking_date, time_slot_id)
		VALUES (?, ?, ?, ?)
	`, user.UserId, courtID, bookingDate, timeSlotID)

	if err != nil {
		fmt.Println("Gagal membuat booking:", err)
		return
	}

	fmt.Println("Booking berhasil dibuat")
}



func ViewMyBookings(db *sql.DB, user *entity.LoginUser) {
	rows, err := db.Query(`
		SELECT b.booking_id, c.court_name, b.booking_date, b.status
		FROM bookings b
		JOIN courts c ON b.court_id = c.court_id
		WHERE b.user_id = ?
	`, user.UserId)

	if err != nil {
		fmt.Println("Error fetching bookings:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== MY BOOKINGS ===")
	for rows.Next() {
		var bookingID int
		var courtName, status string
		var bookingDate time.Time

		if err := rows.Scan(&bookingID, &courtName, &bookingDate, &status); err != nil {
			fmt.Println("Scan error:", err)
			return
		}

		fmt.Printf(
			"ID: %d | Court: %s | Date: %s | Status: %s\n",
			bookingID,
			courtName,
			bookingDate.Format("2006-01-02"),
			status,
		)
	}
}
