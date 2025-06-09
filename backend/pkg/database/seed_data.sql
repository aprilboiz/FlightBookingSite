-- Inserting data into the Airport table
INSERT INTO airports (airport_code, airport_name, city_name, country_name, created_at, updated_at) VALUES
                                                                                                       ('SGN', 'Tan Son Nhat International Airport', 'Ho Chi Minh City', 'Vietnam', NOW(), NOW()),
                                                                                                       ('HAN', 'Noi Bai International Airport', 'Hanoi', 'Vietnam', NOW(), NOW()),
                                                                                                       ('DAD', 'Da Nang International Airport', 'Da Nang', 'Vietnam', NOW(), NOW()),
                                                                                                       ('CXR', 'Cam Ranh International Airport', 'Nha Trang', 'Vietnam', NOW(), NOW()),
                                                                                                       ('PQC', 'Phu Quoc International Airport', 'Phu Quoc', 'Vietnam', NOW(), NOW()),
                                                                                                       ('HPH', 'Cat Bi International Airport', 'Hai Phong', 'Vietnam', NOW(), NOW()),
                                                                                                       ('HUI', 'Phu Bai International Airport', 'Hue', 'Vietnam', NOW(), NOW()),
                                                                                                       ('VDO', 'Van Don International Airport', 'Quang Ninh', 'Vietnam', NOW(), NOW()),
                                                                                                       ('VCA', 'Can Tho International Airport', 'Can Tho', 'Vietnam', NOW(), NOW()),
                                                                                                       ('DLI', 'Lien Khuong International Airport', 'Da Lat', 'Vietnam', NOW(), NOW());


-- Inserting data into the Plane table
INSERT INTO planes (plane_code, plane_name, created_at, updated_at) VALUES
                                                                        ('RUA321', 'Airbus A321', NOW(), NOW()),
                                                                        ('RUA359', 'Airbus A350-900', NOW(), NOW()),
                                                                        ('RUA186', 'Airbus A321', NOW(), NOW()),
                                                                        ('RUA205', 'Boeing 787-9', NOW(), NOW()),
                                                                        ('RUA787', 'Boeing 787-10', NOW(), NOW());

-- Inserting data into the TicketClass table
INSERT INTO ticket_classes (ticket_class_name, price_percentage, created_at, updated_at) VALUES
                                                                                             ('Economy', 1.0, NOW(), NOW()),
                                                                                             ('Business', 1.05, NOW(), NOW());

-- Inserting data into the Seat table (assuming each plane has a mix of Economy and Business class seats)
-- For RUA321 (e.g., 150 Economy, 16 Business)
DO $$
    BEGIN
        FOR i IN 1..150 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA321'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..16 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA321'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For RUA359 (e.g., 250 Economy, 30 Business)
        FOR i IN 1..250 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA359'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..30 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA359'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For RUA186 (e.g., 180 Economy, 8 Business)
        FOR i IN 1..180 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA186'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..8 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA186'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For RUA205 (e.g., 200 Economy, 20 Business)
        FOR i IN 1..200 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA205'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..20 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA205'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For RUA787 (e.g., 280 Economy, 40 Business)
        FOR i IN 1..280 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA787'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..40 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'RUA787'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;
    END$$;

-- Inserting data into the Configuration table
INSERT INTO parameters (number_of_airports, min_flight_duration, max_intermediate_stops, min_intermediate_stop_duration, max_intermediate_stop_duration, max_ticket_classes, latest_ticket_purchase_time, ticket_cancellation_time, created_at, updated_at) VALUES
    (10, 30, 2, 10, 20, 2, 1, 0, NOW(), NOW());


-- Insert admin user
INSERT INTO users (email, username, password, role, created_at, updated_at)
VALUES (
    'admin@example.com',
    'admin',
    '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', -- password: admin123
    'ADMIN',
    NOW(),
    NOW()
);

-- Insert staff users
INSERT INTO users (email, username, password, role, created_at, updated_at)
VALUES 
    ('staff1@example.com', 'staff1', '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', 'STAFF', NOW(), NOW()),
    ('staff2@example.com', 'staff2', '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', 'STAFF', NOW(), NOW()),
    ('staff3@example.com', 'staff3', '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', 'STAFF', NOW(), NOW());

-- Insert super admin users
INSERT INTO users (email, username, password, role, created_at, updated_at)
VALUES 
    ('super@example.com', 'super', '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', 'SUPER_ADMIN', NOW(), NOW()),
    ('super2@example.com', 'super2', '$2a$10$DEpzULpFAvtanyctXdGrFOz1NfETI1LlJoP/Y.Zo0gzOFnIJWhFRe', 'SUPER_ADMIN', NOW(), NOW());