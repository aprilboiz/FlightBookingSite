package com.aprilboiz.flightbookingsite.repository;

import com.aprilboiz.flightbookingsite.entity.SeatClass;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SeatClassRepository extends JpaRepository<SeatClass, Long> {
}
