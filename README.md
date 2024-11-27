# Receipt Processor Challenge

## Project Structure

```bash
receipt-processor/
├── api/
│   ├── handlers.go        # Contains the Gin handlers for endpoints
│   └── router.go          # Initializes Gin routes
├── models/
│   └── receipt.go         # Data models and validation logic
├── services/
│   └── points.go          # Business logic for calculating points
├── storage/
│   └── memory.go          # In-memory storage logic
├── main.go                # Application entry point
├── api.yml                # OpenAPI specification
├── Dockerfile             # Docker configuration
└── go.mod                 # Go module file

```

# Running the Application

## Using Go

1- Clone the repository and navigate to the project folder.

2- Run the following commands:

```bash
go mod tidy
go run main.go

```

3- The server will start at http://localhost:8080.

## Using Docker

1- Build the Docker image:

```bash
docker build -t receipt-processor .
```

2- Run the container:

```bash
docker run -p 8080:8080 receipt-processor
```

# Steps to Test

## 1\. Process Receipt

### Endpoint

**POST** `http://localhost:8080/receipts/process`

### Headers

- `Content-Type: application/json`

### Body

Use the **raw** format in Postman and provide a valid receipt JSON. Example:

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

### Send the Request

Click **Send** in Postman.

### Expected Response

A JSON object containing the `id` of the receipt:

```json
{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### 2\. Get Points for a Receipt

1. **Endpoint**:  
   **GET** `http://localhost:8080/receipts/{id}/points`
2. Example:  
   `http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points`
3. **Send the Request**:  
   Click **Send** in Postman.
4. **Expected Response**:  
   A JSON object containing the number of points awarded for the receipt:

```json
{
  "points": 28
}
```
