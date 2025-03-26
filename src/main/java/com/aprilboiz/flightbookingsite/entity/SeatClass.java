package com.aprilboiz.flightbookingsite.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.validator.constraints.Range;

@Table(name = "seats")
@NoArgsConstructor
@AllArgsConstructor
@Data
@Builder
@Entity
public class SeatClass {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "flight_id")
    private Flight flight;

    @Column(name = "seat_class", nullable = false)
    private String seatClass;

    @Column(name = "number_of_seats", nullable = false)
    private Integer numberOfSeats;
}
