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

-- Insert bookings for next 3 months
INSERT INTO bookings (user_id, court_id, booking_date, time_slot_id, total_price, status, created_at)
VALUES
-- February 2026
(1, 1, '2026-02-15', 1, 200000.00, 'COMPLETED', NOW()),
(2, 2, '2026-02-18', 2, 180000.00, 'BOOKED', NOW()),
(3, 1, '2026-02-20', 3, 220000.00, 'COMPLETED', NOW()),

-- March 2026
(1, 2, '2026-03-05', 1, 200000.00, 'BOOKED', NOW()),
(2, 1, '2026-03-10', 2, 180000.00, 'COMPLETED', NOW()),
(3, 2, '2026-03-15', 3, 220000.00, 'BOOKED', NOW()),

-- April 2026
(1, 1, '2026-04-02', 1, 200000.00, 'COMPLETED', NOW()),
(2, 2, '2026-04-08', 2, 180000.00, 'BOOKED', NOW()),
(3, 1, '2026-04-12', 3, 220000.00, 'COMPLETED', NOW());

INSERT INTO reports (
    report_type,
    total_booking,
    total_revenue,
    generated_by
) VALUES
('DAILY', 2, 380000, 1),
('MONTHLY', 25, 4200000, 1);
