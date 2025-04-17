package models

import (
	"gorm.io/gorm"
	"time"
)

type Plane struct {
	gorm.Model
	PlaneCode string `gorm:"unique;not null"`

	Seats []Seat `gorm:"foreignKey:PlaneID;references:ID"`
}

type TicketClass struct {
	gorm.Model
	TicketClassName string  `gorm:"not null"`
	PricePercentage float64 `gorm:"not null"`

	Seats []Seat `gorm:"foreignKey:TicketClassID;references:ID"`
}

type Airport struct {
	gorm.Model
	AirportCode string `gorm:"not null"`
	CityName    string `gorm:"not null"`
	CountryName string `gorm:"not null"`

	Flights []Flight `gorm:"foreignKey:DepartureAirportID;references:ID"`
}

type Seat struct {
	gorm.Model
	SeatNumber    string `gorm:"not null"`
	PlaneID       uint   `gorm:"not null"`
	TicketClassID uint   `gorm:"not null"`

	TicketClass TicketClass `gorm:"foreignKey:TicketClassID;references:ID"`
	Plane       Plane       `gorm:"foreignKey:PlaneID;references:ID"`
	Tickets     []Ticket    `gorm:"foreignKey:SeatID;references:ID"`
}

type Flight struct {
	gorm.Model
	FlightCode         string    `gorm:"unique;not null"`
	PlaneID            uint      `gorm:"not null"`
	DepartureAirportID uint      `gorm:"not null"`
	ArrivalAirportID   uint      `gorm:"not null"`
	DepartureDateTime  time.Time `gorm:"not null"`
	FlightDuration     int       `gorm:"not null"`
	BasePrice          float64   `gorm:"not null"`

	DepartureAirport Airport `gorm:"foreignKey:DepartureAirportID;references:ID"`
	ArrivalAirport   Airport `gorm:"foreignKey:ArrivalAirportID;references:ID"`
	Plane            Plane   `gorm:"foreignKey:PlaneID;references:ID"`
}

type IntermediateStop struct {
	FlightID     uint   `gorm:"primaryKey"`
	AirportID    uint   `gorm:"primaryKey"`
	StopDuration int    `gorm:"not null"`
	Note         string `gorm:"nullable"`

	Flight  Flight  `gorm:"foreignKey:FlightID;references:ID"`
	Airport Airport `gorm:"foreignKey:AirportID;references:ID"`
}

type Ticket struct {
	gorm.Model
	FlightID     uint    `gorm:"not null"`
	SeatID       uint    `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	FullName     string  `gorm:"not null"`
	IDCard       string  `gorm:"not null"`
	PhoneNumber  string  `gorm:"not null"`
	Email        string  `gorm:"not null"`
	FlightStatus string  `gorm:"not null"`

	Flight Flight `gorm:"foreignKey:FlightID;references:ID"`
	Seat   Seat   `gorm:"foreignKey:SeatID;references:ID"`
}

type Configuration struct {
	gorm.Model
	NumberOfAirports            int `gorm:"not null"`
	MinFlightDuration           int `gorm:"not null"`
	MaxIntermediateStops        int `gorm:"not null"`
	MinIntermediateStopDuration int `gorm:"not null"`
	MaxIntermediateStopDuration int `gorm:"not null"`
	MaxTicketClasses            int `gorm:"not null"`
	LatestTicketPurchaseTime    int `gorm:"not null"`
	TicketCancellationTime      int `gorm:"not null"`
}
