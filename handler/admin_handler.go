package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andrewidianto/Paddle-Booking/entity"
)

func AddCourt(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	var court entity.Courts
	fmt.Println("\n=== ADD COURT ===")

	// Read court name input from cli
	fmt.Print("Court Name: ")
	name, _ := reader.ReadString('\n')
	court.Name = strings.TrimSpace(name)

	// Read location input from cli
	fmt.Print("Location: ")
	location, _ := reader.ReadString('\n')
	court.Location = strings.TrimSpace(location)

	// Read Price input from cli
	fmt.Print("Price per hour: ")
	_, err := fmt.Scanf("%f\n", &court.PricePerHour)
	if err != nil {
		fmt.Println("Invalid price input")
		return
	}

	// Default Status is AVAILABLE
	court.Status = "AVAILABLE"

	query := `INSERT INTO courts (court_name, location, price_per_hour, status) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, court.Name, court.Location, court.PricePerHour, court.Status)
	if err != nil {
		fmt.Println("Failed to add court:", err)
		return
	}

	fmt.Println("Court added successfully with default status AVAILABLE!")
}

func AddTimeSlot(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	var timeslot entity.TimeSlots

	fmt.Println("\n--- ADD TIME SLOT ---")

	// Read Start time input from cli
	fmt.Print("Start Time (HH:MM:SS): ")
	start, _ := reader.ReadString('\n')
	timeslot.StartTime = strings.TrimSpace(start)

	// Read End time input from cli
	fmt.Print("End Time (HH:MM:SS): ")
	end, _ := reader.ReadString('\n')
	timeslot.EndTime = strings.TrimSpace(end)

	query := `INSERT INTO time_slots (start_time, end_time) VALUES (?, ?)`
	_, err := db.Exec(query, timeslot.StartTime, timeslot.EndTime)
	if err != nil {
		fmt.Println("Failed to add time slot:", err)
		return
	}

	fmt.Println("Time Slot added successfully!")
}

func UpdateCourt(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n=== UPDATE COURT ===")

	var court entity.Courts

	// Read Court ID input from cli
	fmt.Print("Input Court ID: ")
	courtIDStr, _ := reader.ReadString('\n')
	courtIDStr = strings.TrimSpace(courtIDStr)
	courtID, err := strconv.Atoi(courtIDStr)
	if err != nil {
		fmt.Println("Invalid Court ID")
		return
	}
	court.ID = courtID

	// Fetch current court data
	query := `SELECT court_name, location, price_per_hour, status FROM courts WHERE court_id = ?`
	row := db.QueryRow(query, court.ID)
	err = row.Scan(&court.Name, &court.Location, &court.PricePerHour, &court.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Court ID not found!")
		} else {
			fmt.Println("Error fetching court:", err)
		}
		return
	}

	// Display current values
	fmt.Println("Current Court Name:", court.Name)
	fmt.Println("Current Location:", court.Location)
	fmt.Println("Current Price per hour:", court.PricePerHour)
	fmt.Println("Current Status:", court.Status)

	// Input new court name
	fmt.Print("New Court Name (leave blank to keep current): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		court.Name = name
	}

	// Input new court location
	fmt.Print("New Location (leave blank to keep current): ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)
	if location != "" {
		court.Location = location
	}

	// Input new court price per hour
	fmt.Print("New Price per hour (leave blank to keep current): ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("Invalid price input")
			return
		}
		court.PricePerHour = price
	}

	// Input new court status
	fmt.Print("New Status (AVAILABLE/MAINTENANCE, leave blank to keep current): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)
	if status != "" {
		court.Status = status
	}

	// Update court in database
	updateQuery := `UPDATE courts SET court_name = ?, location = ?, price_per_hour = ?, status = ? WHERE court_id = ?`
	result, err := db.Exec(updateQuery, court.Name, court.Location, court.PricePerHour, court.Status, court.ID)
	if err != nil {
		fmt.Println("Failed to update court:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No changes made.")
	} else {
		fmt.Println("Court updated successfully!")
	}
}

func UpdateTimeSlot(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n=== UPDATE TIME SLOT ===")

	var timeslot entity.TimeSlots

	// Read Time Slot ID
	fmt.Print("Input Time Slot ID: ")
	slotIDStr, _ := reader.ReadString('\n')
	slotIDStr = strings.TrimSpace(slotIDStr)
	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		fmt.Println("Invalid Time Slot ID")
		return
	}
	timeslot.TimeSlotID = slotID

	// Fetch current data
	query := `SELECT start_time, end_time FROM time_slots WHERE time_slot_id = ?`
	row := db.QueryRow(query, timeslot.TimeSlotID)
	err = row.Scan(&timeslot.StartTime, &timeslot.EndTime)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Time Slot ID not found!")
		} else {
			fmt.Println("Error fetching time slot:", err)
		}
		return
	}

	// Display current values
	fmt.Println("Current Start Time:", timeslot.StartTime)
	fmt.Println("Current End Time:", timeslot.EndTime)

	// Input new values
	fmt.Print("New Start Time (HH:MM:SS, leave blank to keep current): ")
	start, _ := reader.ReadString('\n')
	start = strings.TrimSpace(start)
	if start != "" {
		timeslot.StartTime = start
	}

	fmt.Print("New End Time (HH:MM:SS, leave blank to keep current): ")
	end, _ := reader.ReadString('\n')
	end = strings.TrimSpace(end)
	if end != "" {
		timeslot.EndTime = end
	}

	// Update time slot in database
	updateQuery := `UPDATE time_slots SET start_time = ?, end_time = ? WHERE time_slot_id = ?`
	result, err := db.Exec(updateQuery, timeslot.StartTime, timeslot.EndTime, timeslot.TimeSlotID)
	if err != nil {
		fmt.Println("Failed to update time slot:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No changes made.")
	} else {
		fmt.Println("Time Slot updated successfully!")
	}
}

func AddUser(db *sql.DB) {
	var user entity.RegisterUser

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== ADD USER ===")

	// Read full name input from CLI
	fmt.Print("Full Name: ")
	fullName, _ := reader.ReadString('\n')
	user.Fullname = strings.TrimSpace(fullName)

	// Read password input from CLI
	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	user.Password = strings.TrimSpace(password)

	// Ask user to choose role
	fmt.Println("Role:")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Print("Choose role: ")

	roleStr, _ := reader.ReadString('\n')
	roleStr = strings.TrimSpace(roleStr)

	// Convert CLI role selection into role_id
	switch roleStr {
	case "1":
		user.RoleID = 1
	case "2":
		user.RoleID = 2
	default:
		fmt.Println("Invalid user role input")
		return
	}

	// SQL query to insert new user into database
	query := `
		INSERT INTO users (full_name, password, role_id)
		VALUES (?, ?, ?)
	`

	// Execute insert query
	_, err := db.Exec(query, user.Fullname, user.Password, user.RoleID)
	if err != nil {
		fmt.Println("Failed to add user:", err)
		return
	}

	fmt.Println("User successfully added!")
}

func UpdateUser(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n=== UPDATE USER ===")

	var user entity.User

	// Read User id input from cli
	fmt.Print("Input user ID: ")
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	UserId, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("Invalid User ID")
		return
	}
	user.ID = UserId

	// Check if the user exists in the database
	query := `
        SELECT u.full_name, u.role_id, r.role_name
        FROM users u
        LEFT JOIN roles r ON u.role_id = r.role_id
        WHERE u.user_id = ?`
	row := db.QueryRow(query, user.ID)

	err = row.Scan(&user.FullName, &user.RoleID, &user.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User ID not found!")
		} else {
			fmt.Println("Error fetching user:", err)
		}
		return
	}

	// Display current values
	fmt.Println("Current Full Name:", user.FullName)
	fmt.Println("Current Role Name:", user.RoleName)

	// Read full name input from CLI
	fmt.Print("Input new Full Name (leave blank to keep current): ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)
	if fullName != "" {
		user.FullName = fullName
	}

	// Fetch available roles from roles table
	fmt.Println("\nAvailable Roles:")
	rolesQuery := `SELECT role_id, role_name FROM roles`
	rows, err := db.Query(rolesQuery)
	if err != nil {
		fmt.Println("Error fetching roles:", err)
		return
	}
	defer rows.Close()

	// map from string input -> role_id
	roleMap := make(map[string]int)
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("%d. %s\n", id, name)
		roleMap[strconv.Itoa(id)] = id
	}

	// Let user select new role (optional)
	fmt.Print("Choose role (leave blank to keep current): ")
	roleChoice, _ := reader.ReadString('\n')
	roleChoice = strings.TrimSpace(roleChoice)

	if roleChoice != "" {
		selectedRoleID, ok := roleMap[roleChoice]
		if !ok {
			fmt.Println("Invalid role selection")
			return
		}
		user.RoleID = selectedRoleID
	}

	// Update user in database
	updateQuery := `UPDATE users SET full_name = ?, role_id = ? WHERE user_id = ?`
	result, err := db.Exec(updateQuery, user.FullName, user.RoleID, user.ID)
	if err != nil {
		fmt.Println("Failed to update user:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No changes made.")
	} else {
		fmt.Println("User updated successfully!")
	}
}

func UserReport(db *sql.DB) {
	// Query to get all users
	rows, err := db.Query(`
		SELECT 
			u.user_id,
			u.full_name,
			r.role_name,
			u.created_at
		FROM users u
		JOIN roles r ON r.role_id = u.role_id
		ORDER BY u.user_id
	`)

	if err != nil {
		fmt.Println("Failed to fetch users:", err)
		return
	}

	defer rows.Close()

	fmt.Println("\n=== USER LIST ===")
	// Iterate through result rows
	for rows.Next() {
		var users entity.User

		// Scan database row into entity user
		err := rows.Scan(&users.ID, &users.FullName, &users.RoleName, &users.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		// Print all users
		fmt.Printf(
			"ID: %d | Name: %s | Role: %s | Created: %s\n",
			users.ID,
			users.FullName,
			users.RoleName,
			users.CreatedAt.Format("2006-01-02 15:04"),
		)

	}
}

func BookingReport(db *sql.DB) {
	fmt.Println("\n=== BOOKING REPORT ===")

	// SQL query to get booking details with joined tables
	query := `
        SELECT b.booking_id, u.full_name, c.court_name, DATE(b.booking_date) AS booking_date, t.start_time, t.end_time, b.status, b.total_price
        FROM bookings b
        LEFT JOIN users u ON b.user_id = u.user_id
        LEFT JOIN courts c ON b.court_id = c.court_id
        LEFT JOIN time_slots t ON b.time_slot_id = t.time_slot_id
        ORDER BY b.booking_date, t.start_time
    `
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error fetching bookings:", err)
		return
	}
	defer rows.Close()
	fmt.Println("BookingID | User | Court | Date | Time | Status | Total Price")

	for rows.Next() { // iterate through result set
		var booking entity.Booking
		rows.Scan(&booking.ID, &booking.UserName, &booking.CourtName, &booking.Date, &booking.StartTime, &booking.EndTime, &booking.Status, &booking.TotalPrice)
		fmt.Printf("%d | %s | %s | %s | %s-%s | %s | %.2f\n",
			booking.ID, booking.UserName, booking.CourtName, booking.Date.Format("2006-01-02"), booking.StartTime, booking.EndTime, booking.Status, booking.TotalPrice) // print row
	}
}

func DailyRevenue(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date (YYYY-MM-DD): ") // prompt for date
	date, _ := reader.ReadString('\n')     // read date
	date = strings.TrimSpace(date)         // trim spaces

	query := `
            SELECT SUM(total_price) as total_revenue, COUNT(*) as total_bookings
            FROM bookings
            WHERE booking_date = ? AND status = 'COMPLETED'
        `

	var revenue float64
	var total int
	row := db.QueryRow(query, date) // execute query with parameter
	row.Scan(&revenue, &total)      // scan result
	fmt.Printf("Date: %s | Total Bookings: %d | Total Revenue: %.2f\n", date, total, revenue)
}

func MonthlyRevenue(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter month (YYYY-MM): ") // prompt for month
	month, _ := reader.ReadString('\n')  // read month
	month = strings.TrimSpace(month)     // trim spaces

	query := `
            SELECT SUM(total_price) as total_revenue, COUNT(*) as total_bookings
            FROM bookings
            WHERE DATE_FORMAT(booking_date, '%Y-%m') = ? AND status = 'COMPLETED'
        `

	var revenue float64
	var total int
	row := db.QueryRow(query, month)                                                            // execute query
	row.Scan(&revenue, &total)                                                                  // scan result
	fmt.Printf("Month: %s | Total Bookings: %d | Total Revenue: %.2f\n", month, total, revenue) // print result
}
