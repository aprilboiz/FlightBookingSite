package com.aprilboiz.flightbookingsite.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Table(name = "airports")
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Data
@Entity
public class Airport {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    @Column(name = "airport_code", unique = true, nullable = false)
    private String airportCode;

    private String airportName;

    private String airportAddress;
}
