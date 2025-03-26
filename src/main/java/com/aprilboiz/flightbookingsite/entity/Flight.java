package com.aprilboiz.flightbookingsite.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.List;

@Entity
@Table(name = "flights")
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class Flight {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @Column(name = "flight_number", nullable = false)
    private String flightNumber;

    @Column(name = "base_price", nullable = false)
    private Double basePrice;

    @ManyToOne
    @JoinColumn(name = "departure_airport_id", nullable = false)
    private Airport departureAirport;

    @ManyToOne
    @JoinColumn(name = "arrival_airport_id", nullable = false)
    private Airport arrivalAirport;

    @Column(name = "duration", nullable = false)
    private Double flightTime;

    @Column(name = "departure_time", nullable = false)
    private LocalDateTime departureTime;

    @OneToMany(mappedBy = "flight")
    private List<SeatClass> seatClasses;

    @OneToMany(mappedBy = "flight")
    private List<IntermediateStops> intermediateStops;
}
