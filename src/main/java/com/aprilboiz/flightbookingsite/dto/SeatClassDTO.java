package com.aprilboiz.flightbookingsite.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public record SeatClassDTO(
        @JsonProperty("seat_class")
        String seatClass,

        @JsonProperty("number_of_seats")
        Integer numberOfSeats
) {
}
