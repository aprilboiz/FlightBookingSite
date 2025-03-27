package com.aprilboiz.flightbookingsite.service;

import com.aprilboiz.flightbookingsite.entity.Airport;
import com.aprilboiz.flightbookingsite.repository.AirportRepository;
import org.springframework.stereotype.Service;

@Service
public class AirportService {
    private final AirportRepository airportRepository;

    public AirportService(AirportRepository airportRepository) {
        this.airportRepository = airportRepository;
    }

    public Airport getAirportByName(String airportName) {
        return airportRepository.findAirportByName(airportName).orElse(null);
    }
}
