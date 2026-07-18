-- +migrate Up
-- ==========================================
-- 1. DATA DEFINITION LANGUAGE (Tabel Utama)
-- ==========================================

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    role VARCHAR(20) DEFAULT 'customer', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    brand VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    license_plate VARCHAR(20) UNIQUE NOT NULL,
    description TEXT,
    price_per_day DECIMAL(12, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE car_images (
    id SERIAL PRIMARY KEY,
    car_id INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    car_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_price DECIMAL(12, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE RESTRICT,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE RESTRICT
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL UNIQUE,
    amount DECIMAL(12, 2) NOT NULL,
    payment_method VARCHAR(50), 
    payment_status VARCHAR(20) DEFAULT 'pending', 
    transaction_id VARCHAR(100), 
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE RESTRICT
);

-- ==========================================
-- 2. INDEXING & KEAMANAN (PostgreSQL)
-- ==========================================

-- Index untuk mempercepat query pencarian jadwal
CREATE INDEX idx_bookings_car_dates ON bookings (car_id, start_date, end_date);

-- Ekstensi untuk range logic
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Constraint absolut untuk mencegah dua booking tumpang tindih di database level
ALTER TABLE bookings 
ADD CONSTRAINT prevent_double_booking 
EXCLUDE USING GIST (
    car_id WITH =,
    daterange(start_date, end_date) WITH &&
) WHERE (status IN ('pending', 'paid', 'active'));

-- +migrate Down
ALTER TABLE bookings DROP CONSTRAINT prevent_double_booking;
DROP INDEX idx_bookings_car_dates;
DROP TABLE payments;
DROP TABLE bookings;
DROP TABLE car_images;
DROP TABLE cars;
DROP TABLE users;
