package com.aprilboiz.flightbookingsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name = "intermediate_stops")
@NoArgsConstructor
@AllArgsConstructor
@Data
@Builder
@Entity
public class IntermediateStop {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "airport_id", nullable = false)
    private Airport airport;

    @ManyToOne
    @JoinColumn(name = "flight_id", nullable = false)
    private Flight flight;

    @Column(nullable = false)
    private Integer stopDuration;
    private String note;
}
