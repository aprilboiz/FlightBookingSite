package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User role constants
const (
	RoleUser  = "USER"  // Regular user
	RoleStaff = "STAFF" // Staff member
	RoleAdmin = "ADMIN" // Administrator
)

// Ticket status constants
const (
	TicketStatusActive    = "ACTIVE"    // Ticket is active and can be used
	TicketStatusCancelled = "CANCELLED" // Ticket has been cancelled
	TicketStatusUsed      = "USED"      // Ticket has been used for the flight
	TicketStatusExpired   = "EXPIRED"   // Ticket has expired (for place orders)
	TicketStatusRefunded  = "REFUNDED"  // Ticket has been refunded
)

// Booking type constants
const (
	BookingTypeTicket     = "TICKET"      // Regular confirmed ticket
	BookingTypePlaceOrder = "PLACE_ORDER" // Temporary place order
)

type Plane struct {
	gorm.Model
	PlaneCode string `gorm:"unique;not null"`
	PlaneName string `gorm:"not null"`

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
	AirportName string `gorm:"not null"`
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

	DepartureAirport  Airport `gorm:"foreignKey:DepartureAirportID;references:ID"`
	ArrivalAirport    Airport `gorm:"foreignKey:ArrivalAirportID;references:ID"`
	Plane             Plane   `gorm:"foreignKey:PlaneID;references:ID"`
	IntermediateStops []IntermediateStop
	Tickets           []Ticket
}

type IntermediateStop struct {
	FlightID     uint   `gorm:"primaryKey"`
	AirportID    uint   `gorm:"primaryKey"`
	StopOrder    int    `gorm:"not null"`
	StopDuration int    `gorm:"not null"`
	Note         string `gorm:"nullable"`

	Flight  Flight  `gorm:"foreignKey:FlightID;references:ID"`
	Airport Airport `gorm:"foreignKey:AirportID;references:ID"`
}

type Ticket struct {
	ID           uint    `gorm:"primaryKey"`
	FlightID     uint    `gorm:"primaryKey"`
	SeatID       uint    `gorm:"primaryKey"`
	Price        float64 `gorm:"not null"`
	FullName     string  `gorm:"not null"`
	IDCard       string  `gorm:"not null"`
	PhoneNumber  string  `gorm:"not null"`
	Email        string  `gorm:"not null"`
	TicketStatus string  `gorm:"not null;default:'ACTIVE'"` // Status of the ticket
	BookingType  string  `gorm:"not null;default:'TICKET'"` // Type of booking

	Flight Flight `gorm:"foreignKey:FlightID;references:ID"`
	Seat   Seat   `gorm:"foreignKey:SeatID;references:ID"`
}

type Parameter struct {
	gorm.Model                  `json:"-"`
	NumberOfAirports            int    `gorm:"not null" json:"number_of_airports"`
	MinFlightDuration           int    `gorm:"not null" json:"min_flight_duration"`
	MaxIntermediateStops        int    `gorm:"not null" json:"max_intermediate_stops"`
	MinIntermediateStopDuration int    `gorm:"not null" json:"min_intermediate_stop_duration"`
	MaxIntermediateStopDuration int    `gorm:"not null" json:"max_intermediate_stop_duration"`
	MaxTicketClasses            int    `gorm:"not null" json:"max_ticket_classes"`
	LatestTicketPurchaseTime    int    `gorm:"not null" json:"latest_ticket_purchase_time"`
	TicketCancellationTime      int    `gorm:"not null" json:"ticket_cancellation_time"`
	lock                        string `gorm:"type:char(1);unique;not null;default:'X';check:lock='X'"`
}

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Role      string         `gorm:"not null;default:'user'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
