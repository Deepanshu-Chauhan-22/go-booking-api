package booking

import (
	"bookingapp/models"
	"testing"
)

func resetBookingState() {
	RemainingTickets = 20000
	members = nil
	UserData = make(map[string]models.User)
}

func TestBookingSuccess(t *testing.T) {
	resetBookingState()

	user, err := Booking("John Doe", "john@example.com", 2)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "John Doe" {
		t.Errorf("expected name John Doe, got %s", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", user.Email)
	}

	if user.TicketsBought != 2 {
		t.Errorf("expected 2 tickets, got %d", user.TicketsBought)
	}

	if RemainingTickets != 19998 {
		t.Errorf("expected 19998 remaining tickets, got %d", RemainingTickets)
	}
}

func TestBookingDuplicateEmail(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 2)
	if err != nil {
		t.Fatalf("first booking should succeed, got %v", err)
	}

	_, err = Booking("Jane Doe", "john@example.com", 1)

	if err == nil {
		t.Fatal("expected duplicate email error, got nil")
	}

	if err.Error() != "email already has a booking; update the existing booking instead" {
		t.Errorf("unexpected error: %v", err)
	}

	if RemainingTickets != 19998 {
		t.Errorf("expected 19998 remaining tickets, got %d", RemainingTickets)
	}
}

func TestBookingInsufficientTickets(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 20001)

	if err == nil {
		t.Fatal("expected insufficient tickets error, got nil")
	}

	if RemainingTickets != 20000 {
		t.Errorf("expected 20000 remaining tickets, got %d", RemainingTickets)
	}
}

func TestGetBookings(t *testing.T) {
	resetBookingState()

	Booking("John Doe", "john@example.com", 2)
	Booking("Jane Doe", "jane@example.com", 3)

	bookings := GetBookings()

	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings, got %d", len(bookings))
	}

	if bookings[0].Email != "john@example.com" {
		t.Errorf("expected first email john@example.com, got %s", bookings[0].Email)
	}

	if bookings[1].Email != "jane@example.com" {
		t.Errorf("expected second email jane@example.com, got %s", bookings[1].Email)
	}

	if bookings[0].TicketsBought != 2 {
		t.Errorf("expected first booking to have 2 tickets, got %d", bookings[0].TicketsBought)
	}

	if bookings[1].TicketsBought != 3 {
		t.Errorf("expected second booking to have 3 tickets, got %d", bookings[1].TicketsBought)
	}
}

func TestUserDatabyEmail(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 2)
	if err != nil {
		t.Fatalf("booking should succeed, got %v", err)
	}

	user, err := UserDatabyEmail("john@example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "John Doe" {
		t.Errorf("expected name John Doe, got %s", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", user.Email)
	}

	if user.TicketsBought != 2 {
		t.Errorf("expected 2 tickets, got %d", user.TicketsBought)
	}
}

func TestUserDatabyEmailNotFound(t *testing.T) {
	resetBookingState()

	_, err := UserDatabyEmail("doesnotexist@example.com")

	if err == nil {
		t.Fatal("expected error for non-existent user, got nil")
	}

	if err.Error() != "User doesn't exist" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateTickets(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 2)
	if err != nil {
		t.Fatalf("booking should succeed, got %v", err)
	}

	user, err := UpdateTickets("john@example.com", 5)

	if err != nil {
		t.Fatalf("expected update to succeed, got %v", err)
	}

	if user.TicketsBought != 5 {
		t.Errorf("expected 5 tickets, got %d", user.TicketsBought)
	}

	if RemainingTickets != 19995 {
		t.Errorf("expected 19995 remaining tickets, got %d", RemainingTickets)
	}
}

func TestUpdateTicketsDecrease(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 5)
	if err != nil {
		t.Fatalf("booking should succeed, got %v", err)
	}

	user, err := UpdateTickets("john@example.com", 2)

	if err != nil {
		t.Fatalf("expected update to succeed, got %v", err)
	}

	if user.TicketsBought != 2 {
		t.Errorf("expected 2 tickets, got %d", user.TicketsBought)
	}

	if RemainingTickets != 19998 {
		t.Errorf("expected 19998 remaining tickets, got %d", RemainingTickets)
	}
}

func TestUpdateTicketsNotFound(t *testing.T) {
	resetBookingState()

	_, err := UpdateTickets("unknown@example.com", 5)

	if err == nil {
		t.Fatal("expected error for non-existent booking, got nil")
	}

	if err.Error() != "Booking not found" {
		t.Errorf("unexpected error: %v", err)
	}

	if RemainingTickets != 20000 {
		t.Errorf("expected 20000 remaining tickets, got %d", RemainingTickets)
	}
}

func TestUpdateTicketsInsufficientTickets(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 2)
	if err != nil {
		t.Fatalf("booking should succeed, got %v", err)
	}

	_, err = UpdateTickets("john@example.com", 20001)

	if err == nil {
		t.Fatal("expected insufficient tickets error, got nil")
	}

	if err.Error() != "Not enough tickets remaining" {
		t.Errorf("unexpected error: %v", err)
	}

	if RemainingTickets != 19998 {
		t.Errorf("expected 19998 remaining tickets, got %d", RemainingTickets)
	}
}

func TestCancelBooking(t *testing.T) {
	resetBookingState()

	_, err := Booking("John Doe", "john@example.com", 5)
	if err != nil {
		t.Fatalf("booking should succeed, got %v", err)
	}

	user, err := CancelBooking("john@example.com")

	if err != nil {
		t.Fatalf("expected cancellation to succeed, got %v", err)
	}

	if user.Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", user.Email)
	}

	if user.TicketsBought != 5 {
		t.Errorf("expected 5 tickets, got %d", user.TicketsBought)
	}

	if RemainingTickets != 20000 {
		t.Errorf("expected 20000 remaining tickets, got %d", RemainingTickets)
	}

	if len(members) != 0 {
		t.Errorf("expected 0 members, got %d", len(members))
	}

	if len(UserData) != 0 {
		t.Errorf("expected UserData to be empty, got %d entries", len(UserData))
	}
}

func TestCancelBookingNotFound(t *testing.T) {
	resetBookingState()

	_, err := CancelBooking("unknown@example.com")

	if err == nil {
		t.Fatal("expected error for non-existent booking, got nil")
	}

	if err.Error() != "User doesn't Exist" {
		t.Errorf("unexpected error: %v", err)
	}

	if RemainingTickets != 20000 {
		t.Errorf("expected 20000 remaining tickets, got %d", RemainingTickets)
	}

	if len(members) != 0 {
		t.Errorf("expected 0 members, got %d", len(members))
	}

	if len(UserData) != 0 {
		t.Errorf("expected UserData to be empty, got %d entries", len(UserData))
	}
}