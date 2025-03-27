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

    @Column(nullable = false)
    private String flightNumber;

    @ManyToOne
    @JoinColumn(nullable = false, name = "plane_id")
    private Plane plane;

    @Column(nullable = false)
    private Double basePrice;

    @ManyToOne
    @JoinColumn(nullable = false, name = "departure_airport_id")
    private Airport departureAirport;

    @ManyToOne
    @JoinColumn(nullable = false, name = "arrival_airport_id")
    private Airport arrivalAirport;

    @Column(nullable = false)
    private Double flightTime;

    @Column(nullable = false)
    private LocalDateTime departureTime;

    @OneToMany
    private List<Seat> seats;

    @OneToMany
    private List<IntermediateStop> intermediateStops;
}
