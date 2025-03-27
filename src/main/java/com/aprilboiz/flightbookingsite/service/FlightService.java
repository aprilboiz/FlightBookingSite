package com.aprilboiz.flightbookingsite.service;

import com.aprilboiz.flightbookingsite.dto.FlightRequestDTO;
import com.aprilboiz.flightbookingsite.dto.IntermediateStopRequestDTO;
import com.aprilboiz.flightbookingsite.dto.SeatClassDTO;
import com.aprilboiz.flightbookingsite.entity.Airport;
import com.aprilboiz.flightbookingsite.entity.SeatClass;
import com.aprilboiz.flightbookingsite.repository.FlightRepository;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@Service
public class FlightService {
    private final AirportService airportService;
    private final FlightRepository flightRepository;
    private final String airlineCode = "RuaAirplane";

    public FlightService(AirportService airportService, FlightRepository flightRepository) {
        this.airportService = airportService;
        this.flightRepository = flightRepository;
    }

    public String generateFlightCode(String airlineCode, int flightNo) {
        return this.airlineCode + flightNo;
    }

    public void addFlight(FlightRequestDTO flightRequest) {
        Long nextId = flightRepository.getNextSeriesId();
        String flightCode = generateFlightCode(airlineCode, nextId.intValue());
        Double basePrice = flightRequest.basePrice();
        Airport departureAirport = airportService.getAirportByName(flightRequest.departureAirport());
        Airport arrivalAirport = airportService.getAirportByName(flightRequest.arrivalAirport());
        LocalDateTime departureDateTime = LocalDateTime.parse(
                flightRequest.departureTime(),
                DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")
        );
        Integer flightDuration = flightRequest.flightDuration();
//        List<SeatClass> seatClasses = flightRequest.seatClasses().stream()
//                .map(seatClassDTO -> new SeatClass(
//                        seatClassDTO.seatClass(),
//                        seatClassDTO.seatPrice()
//                ))
//                .toList(); ;
        List<IntermediateStopRequestDTO> intermediateStops = flightRequest.intermediateStops();


    }
}
