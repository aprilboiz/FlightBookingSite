package com.aprilboiz.flightbookingsite.repository;

import com.aprilboiz.flightbookingsite.entity.IntermediateStop;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface IntermediateStopRepository extends JpaRepository<IntermediateStop, Long> {
}
