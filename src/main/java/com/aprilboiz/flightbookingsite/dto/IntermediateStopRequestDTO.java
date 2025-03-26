package com.aprilboiz.flightbookingsite.dto;

import org.hibernate.validator.constraints.Range;

public record IntermediateStopRequestDTO(
        String airport_code,
        @Range(min = 10, max = 20, message = "The stop time should be between 10 and 20 minutes!")
        Integer duration,
        String note
) {
}
