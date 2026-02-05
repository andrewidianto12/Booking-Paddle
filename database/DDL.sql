CREATE TABLE roles (
    role_id INT AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_users_role
        FOREIGN KEY (role_id)
        REFERENCES roles(role_id)
);

CREATE TABLE courts (
    court_id INT AUTO_INCREMENT PRIMARY KEY,
    court_name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    price_per_hour DECIMAL(10,2) NOT NULL,
    status ENUM('AVAILABLE', 'MAINTENANCE') DEFAULT 'AVAILABLE'
);

CREATE TABLE time_slots (
    time_slot_id INT AUTO_INCREMENT PRIMARY KEY,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
);

CREATE TABLE bookings (
    booking_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    court_id INT NOT NULL,
    booking_date DATE NOT NULL,
    time_slot_id INT NOT NULL,
    total_price DECIMAL(10,2),
    status ENUM('BOOKED', 'CANCELLED', 'COMPLETED') DEFAULT 'BOOKED',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_booking_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id),

    CONSTRAINT fk_booking_court
        FOREIGN KEY (court_id)
        REFERENCES courts(court_id),

    CONSTRAINT fk_booking_time
        FOREIGN KEY (time_slot_id)
        REFERENCES time_slots(time_slot_id),

    CONSTRAINT uq_booking UNIQUE (court_id, booking_date, time_slot_id)
);

CREATE TABLE reports (
    report_id INT AUTO_INCREMENT PRIMARY KEY,
    report_type ENUM('DAILY', 'MONTHLY'),
    total_booking INT,
    total_revenue DECIMAL(12,2),
    generated_by INT,
    generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_report_admin
        FOREIGN KEY (generated_by)
        REFERENCES users(user_id)
);