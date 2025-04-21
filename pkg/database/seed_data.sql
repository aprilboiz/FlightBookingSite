-- Inserting data into the Airport table
INSERT INTO airports (airport_code, airport_name, city_name, country_name, created_at, updated_at) VALUES
                                                                                                       ('HAN', 'Noi Bai International Airport', 'Hanoi', 'Vietnam', NOW(), NOW()),
                                                                                                       ('SGN', 'Tan Son Nhat International Airport', 'Ho Chi Minh City', 'Vietnam', NOW(), NOW()),
                                                                                                       ('DAD', 'Da Nang International Airport', 'Da Nang', 'Vietnam', NOW(), NOW()),
                                                                                                       ('CXR', 'Cam Ranh International Airport', 'Nha Trang', 'Vietnam', NOW(), NOW()),
                                                                                                       ('HUI', 'Phu Bai International Airport', 'Hue', 'Vietnam', NOW(), NOW()),
                                                                                                       ('VDO', 'Van Don International Airport', 'Quang Ninh', 'Vietnam', NOW(), NOW()),
                                                                                                       ('PQC', 'Phu Quoc International Airport', 'Phu Quoc', 'Vietnam', NOW(), NOW()),
                                                                                                       ('UIH', 'Phu Cat Airport', 'Quy Nhon', 'Vietnam', NOW(), NOW()),
                                                                                                       ('VII', 'Vinh International Airport', 'Vinh', 'Vietnam', NOW(), NOW()),
                                                                                                       ('BMV', 'Ca Mau Airport', 'Ca Mau', 'Vietnam', NOW(), NOW());

-- Inserting data into the Plane table
INSERT INTO planes (plane_code, plane_name, created_at, updated_at) VALUES
                                                                        ('VNA321', 'Airbus A321', NOW(), NOW()),
                                                                        ('VNA359', 'Airbus A350-900', NOW(), NOW()),
                                                                        ('VJC186', 'Airbus A321', NOW(), NOW()),
                                                                        ('BAM205', 'Boeing 787-9', NOW(), NOW()),
                                                                        ('VNA787', 'Boeing 787-10', NOW(), NOW());

-- Inserting data into the TicketClass table
INSERT INTO ticket_classes (ticket_class_name, price_percentage, created_at, updated_at) VALUES
                                                                                             ('Economy', 1.0, NOW(), NOW()),
                                                                                             ('Business', 1.05, NOW(), NOW());

-- Inserting data into the Seat table (assuming each plane has a mix of Economy and Business class seats)
-- For VNA321 (e.g., 150 Economy, 16 Business)
DO $$
    BEGIN
        FOR i IN 1..150 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA321'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..16 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA321'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For VNA359 (e.g., 250 Economy, 30 Business)
        FOR i IN 1..250 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA359'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..30 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA359'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For VJC186 (e.g., 180 Economy, 8 Business)
        FOR i IN 1..180 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VJC186'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..8 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VJC186'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For BAM205 (e.g., 200 Economy, 20 Business)
        FOR i IN 1..200 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'BAM205'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..20 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'BAM205'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;

        -- For VNA787 (e.g., 280 Economy, 40 Business)
        FOR i IN 1..280 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA787'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Economy'), 'E' || i, NOW(), NOW());
            END LOOP;
        FOR i IN 1..40 LOOP
                INSERT INTO seats (plane_id, ticket_class_id, seat_number, created_at, updated_at)
                VALUES ((SELECT id FROM planes WHERE plane_code = 'VNA787'), (SELECT id FROM ticket_classes WHERE ticket_class_name = 'Business'), 'B' || i, NOW(), NOW());
            END LOOP;
    END$$;

-- Inserting data into the Configuration table
INSERT INTO configurations (number_of_airports, min_flight_duration, max_intermediate_stops, min_intermediate_stop_duration, max_intermediate_stop_duration, max_ticket_classes, latest_ticket_purchase_time, ticket_cancellation_time, created_at, updated_at) VALUES
    (10, 30, 2, 10, 20, 2, 1440, 0, NOW(), NOW());