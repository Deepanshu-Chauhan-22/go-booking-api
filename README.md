

```markdown
# Go Booking API

A concurrent REST API built with Go for managing conference ticket bookings, with thread-safe inventory management, CRUD operations, automated tests, and load testing.

## 🚀 Features

- RESTful API built using Go's `net/http`
- Create, retrieve, update, and cancel bookings
- Thread-safe booking and ticket inventory management using `sync.Mutex`
- Duplicate email detection
- Ticket availability validation
- JSON request/response handling
- Concurrent request handling
- Unit tests using Go's standard `testing` package
- Race detection using Go's race detector
- Load testing with Vegeta

## 🛠️ Tech Stack

- **Language:** Go
- **HTTP:** `net/http`
- **Concurrency:** Goroutines, `sync.Mutex`
- **Testing:** Go `testing`, Race Detector
- **Load Testing:** Vegeta
- **Data Storage:** In-memory data structures

## 📁 Project Structure

```text
booking-app-go/
│
├── booking/
│   ├── book.go
│   └── book_test.go
│
├── collectData/
│   └── collect.go
│
├── models/
│   └── user.go
│
├── sendEmail/
│   └── send.go
│
├── validation/
│   └── validate.go
│
├── api.go
├── main.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

## 🔌 API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/health` | Check API health |
| GET | `/tickets` | Get remaining tickets |
| POST | `/bookings` | Create a booking |
| GET | `/bookings` | Get all bookings |
| GET | `/bookings/{email}` | Get booking by email |
| PUT | `/bookings/{email}` | Update ticket count |
| DELETE | `/bookings/{email}` | Cancel a booking |

### Create Booking

```http
POST /bookings
Content-Type: application/json
```

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "tickets": 2
}
```

### Update Booking

```http
PUT /bookings/john@example.com
Content-Type: application/json
```

```json
5
```

## 🔒 Concurrency & Thread Safety

The application maintains shared booking state in memory:

- Remaining ticket inventory
- Booking list
- Email-to-booking mapping

Access to this shared state is protected using `sync.Mutex` to prevent concurrent requests from causing inconsistent state or overselling tickets.

The project was also tested using Go's race detector:

```bash
go test -race ./booking
```

Result:

```text
ok    bookingapp/booking
```

## 🧪 Testing

The booking business logic currently has **12 automated tests** covering:

- Successful bookings
- Duplicate bookings
- Insufficient ticket inventory
- Retrieving bookings
- User lookup
- Updating ticket counts
- Increasing/decreasing ticket counts
- Invalid updates
- Cancelling bookings
- Cancelling non-existent bookings

Run the tests:

```bash
go test ./...
```

Run with verbose output:

```bash
go test -v ./...
```

Run with the race detector:

```bash
go test -race ./booking
```

## ⚡ Load Testing

The API was load-tested locally using **Vegeta**.

### GET `/tickets`

| Offered RPS | Throughput | P95 Latency | Success |
|---:|---:|---:|---:|
| 50 | 50.06/s | 1.414 ms | 100% |
| 200 | 199.94/s | 1.354 ms | 100% |
| 500 | 499.67/s | 1.523 ms | 100% |
| 1,000 | 999.30/s | 1.856 ms | 100% |

### POST `/bookings`

| Offered RPS | Throughput | P95 Latency | Success |
|---:|---:|---:|---:|
| 50 | 50.06/s | 2.231 ms | 100% |
| 500 | 499.64/s | 1.701 ms | 100% |
| 1,000 | 999.47/s | 1.961 ms | 100% |

### GET `/bookings` — ~15,000 bookings

| Offered RPS | Throughput | P95 Latency | Success |
|---:|---:|---:|---:|
| 100 | 99.96/s | 559 ms | 100% |
| 200 | 195.53/s | 901 ms | 100% |
| 300 | 191.12/s | 7.664 s | 95.93% |
| 400 | 184.32/s | 4.087 s | 60.32% |
| 1,000 | 185.57/s | 7.896 s | 28.37% |

> Benchmarks were performed locally using Vegeta. Results are environment-dependent and should not be interpreted as production capacity.

## ▶️ Running Locally

Clone the repository:

```bash
git clone <your-repository-url>
cd booking-app-go
```

Run the application:

```bash
go run .
```

The API will start on:

```text
http://localhost:8080
```

Run the test suite:

```bash
go test ./...
```

## 📌 Current Limitations

- Booking data is stored in memory and is lost when the application restarts.
- No persistent database is currently used.
- `GET /bookings` returns the complete booking list rather than using pagination.
- Authentication and authorization are not implemented.

## 🔮 Future Improvements

- PostgreSQL persistence
- Pagination for booking retrieval
- Better request validation and error handling
- HTTP handler tests using `httptest`
- Structured logging
- Authentication and authorization
- Dockerization
- API documentation with Swagger/OpenAPI
- Improved application architecture and dependency management

## 👨‍💻 Author

**Deepanshu Chauhan**

Built as a backend engineering project to explore Go, REST APIs, concurrency, testing, and performance engineering.
```
