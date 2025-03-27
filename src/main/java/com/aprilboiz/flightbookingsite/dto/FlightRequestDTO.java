package com.aprilboiz.flightbookingsite.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.Size;

import java.util.List;

public record FlightRequestDTO(
        @JsonProperty("flight_number")
        String flightNumber,

        @JsonProperty("base_price")
        Double basePrice,

        @JsonProperty("departure_airport")
        String departureAirport,

        @JsonProperty("arrival_airport")
        String arrivalAirport,

        @JsonProperty("departure_time")
        String departureTime,

        @JsonProperty("flight_duration")
        @Min(value = 30, message = "Flight duration must be at least 30 minutes!")
        Integer flightDuration,

        @JsonProperty("seat_classes")
        List<SeatClassDTO> seatClasses,

        @JsonProperty("intermediate_stops")
        @Size(max = 2, message = "Intermediate stops must be less than 2 places!")
        List<IntermediateStopRequestDTO> intermediateStops
) {
}
