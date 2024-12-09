Product Management System with Asynchronous Image Processing
Project Overview
This project implements a RESTful API in Go for managing products, with a focus on asynchronous image processing, caching, logging, and high scalability. It leverages RabbitMQ for queuing image processing tasks, Redis for caching product data, and Zap for structured logging.

Features
API Endpoints:
POST /products: Create a new product.
GET /products/:id: Get product details by ID.
GET /products: Get all products, with optional filtering by price range and name.
Asynchronous Image Processing using RabbitMQ.
Caching using Redis for faster retrieval.
Logging with structured logs using Zap.
Technologies Used
Go: Backend development.
PostgreSQL: Database to store product and user data.
RabbitMQ: Message queue for asynchronous image processing.
Redis: Caching system.
Zap: Structured logging.
Setup Instructions
1. Clone the Repository
bash
Copy code
git clone https://github.com/daniyal1337/zocket-daniyal.git
cd zocket-daniyal
2. Install Dependencies
Install Go dependencies:

bash
Copy code
go mod tidy
Install the required packages:

bash
Copy code
go get github.com/gin-gonic/gin
go get github.com/go-redis/redis/v8
go get github.com/streadway/amqp
go get go.uber.org/zap
3. Set Up Services
PostgreSQL: Ensure you have PostgreSQL installed and create the necessary users and products tables.
RabbitMQ: Ensure RabbitMQ is installed and running on localhost:5672.
Redis: Ensure Redis is running on localhost:6379.
4. Run the Application
bash
Copy code
go run main.go cache.go logs.go
The application will run on http://localhost:8080.

5. API Endpoints
POST /products: Create a product with data such as user_id, product_name, product_description, product_images, and product_price.
GET /products/:id: Retrieve the product by ID.
GET /products: Retrieve all products, with optional filtering by user_id, price, and product_name.
Testing
Unit Tests: Write unit tests for API endpoints and core functions.
Integration Tests: Validate caching, database interaction, and message queue functionality.
Run tests using:

bash
Copy code
go test
Logging
All API requests, errors, and image processing events are logged using Zap for structured logging.

Contributing
Feel free to fork the repository, submit issues, or pull requests.

License
This project is licensed under the MIT License.

