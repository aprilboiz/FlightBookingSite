package com.aprilboiz.flightbookingsite.repository;

import com.aprilboiz.flightbookingsite.entity.Flight;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FlightRepository extends JpaRepository<Flight, Long> {

}
