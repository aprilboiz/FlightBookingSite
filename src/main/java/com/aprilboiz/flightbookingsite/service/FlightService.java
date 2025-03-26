package com.aprilboiz.flightbookingsite.service;

import com.aprilboiz.flightbookingsite.repository.FlightRepository;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;

@Service
public class FlightService {
    private final FlightRepository flightRepository;

    public FlightService(FlightRepository flightRepository) {
        this.flightRepository = flightRepository;
    }

    public String generateFlightCode(String airlineCode, int flightNo, LocalDate departureDate){
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("ddMMyyyy");
        String formattedDepartureDate = departureDate.format(formatter);

        return airlineCode + flightNo + "-" + formattedDepartureDate;
    }
}
