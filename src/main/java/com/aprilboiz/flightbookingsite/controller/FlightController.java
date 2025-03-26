package com.aprilboiz.flightbookingsite.controller;

import com.aprilboiz.flightbookingsite.dto.FlightRequestDTO;
import com.aprilboiz.flightbookingsite.entity.Flight;
import com.aprilboiz.flightbookingsite.service.FlightService;
import jakarta.validation.Valid;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/flights")
public class FlightController {
    private final FlightService flightService;

    public FlightController(FlightService flightService) {
        this.flightService = flightService;
    }

    @PostMapping
    public ResponseEntity<String> addFlight(@Valid @RequestBody FlightRequestDTO flightRequest) {
        return ResponseEntity.ok("Successfully added flight");
    }
}
