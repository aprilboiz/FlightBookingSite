package com.aprilboiz.flightbookingsite.dto;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.Size;

import java.time.LocalDateTime;
import java.util.List;

public record FlightRequestDTO(
        String flight_number,
        Double base_price,
        String departure_airport,
        String arrival_airport,
        LocalDateTime departure_time,
        @Min(value = 30, message = "Flight duration must be at least 30 minutes!")
        Integer duration,
        List<SeatClassDTO> seat_classes,
        @Size(max = 2, message = "Intermediate stops must be less than 2 places!")
        List<IntermediateStopRequestDTO> intermediate_stops
) {
}
