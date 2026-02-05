INSERT INTO roles (role_name) VALUES
('ADMIN'),
('USER');


INSERT INTO users (full_name, password, role_id) VALUES
('Admin Padel', 'admin123', 1),
('Budi Santoso', 'user123', 2),
('Siti Aminah', 'user123', 2);

INSERT INTO courts (court_name, location, price_per_hour, status) VALUES
('Padel Court A', 'Jakarta Selatan', 200000, 'AVAILABLE'),
('Padel Court B', 'Jakarta Selatan', 180000, 'AVAILABLE'),
('Padel Court C', 'Jakarta Selatan', 150000, 'MAINTENANCE');

INSERT INTO time_slots (start_time, end_time) VALUES
('08:00', '09:00'),
('09:00', '10:00'),
('10:00', '11:00'),
('11:00', '12:00'),
('13:00', '14:00'),
('14:00', '15:00');

INSERT INTO bookings (
    user_id,
    court_id,
    booking_date,
    time_slot_id,
    total_price,
    status
) VALUES
(2, 1, '2026-02-10', 1, 200000, 'BOOKED'),
(3, 2, '2026-02-10', 2, 180000, 'COMPLETED');

INSERT INTO reports (
    report_type,
    total_booking,
    total_revenue,
    generated_by
) VALUES
('DAILY', 2, 380000, 1),
('MONTHLY', 25, 4200000, 1);
