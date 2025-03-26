package com.aprilboiz.flightbookingsite.repository;

import com.aprilboiz.flightbookingsite.entity.Airport;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AirportRepository extends JpaRepository<Airport, Long> {
}
