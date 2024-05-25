# Flight Route Reconstruction

web application built using the Golang and the Echo framework. The application provides a single HTTP endpoint that accepts POST requests containing a JSON payload of flight tickets and reconstructs the complete itinerary from the sequence of ticket tuples provided.

## Features

- Accepts a POST request with a JSON payload of flight tickets.
- Reconstructs the complete itinerary from the provided tickets.
- Returns the reconstructed itinerary as a JSON array of strings.

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/baselrabia/echo-task.git
cd echo-task
```
2. **Install dependencies:**
```bash
go mod download
```
## Usage
1. **Run the application:**

```bash
go run main.go
```

2. **Send a POST request to the /route endpoint:**
```bash
curl -X POST http://localhost:8080/route \
-H "Content-Type: application/json" \
-d '[{"source":"LAX","destination":"DXB"},{"source":"JFK","destination":"LAX"},{"source":"SFO","destination":"SJC"},{"source":"DXB","destination":"SFO"}]'

```

### Example

For the given input:
```json
[
    {"source": "LAX", "destination": "DXB"},
    {"source": "JFK", "destination": "LAX"},
    {"source": "SFO", "destination": "SJC"},
    {"source": "DXB", "destination": "SFO"}
]

```
The expected output is:
```bash
["JFK", "LAX", "DXB", "SFO", "SJC"]

```

## Tests

Unit tests are provided to ensure the functionality of the itinerary reconstruction. To run the tests:

```bash 
go test ./...
```