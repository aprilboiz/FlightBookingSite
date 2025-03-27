package com.aprilboiz.flightbookingsite.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.hibernate.validator.constraints.Range;

public record IntermediateStopRequestDTO(
        @JsonProperty("airport_code")
        String airportCode,

        @JsonProperty("duration")
        @Range(min = 10, max = 20, message = "The stop time should be between 10 and 20 minutes!")
        Integer duration,

        @JsonProperty("note")
        String note
) {
}
