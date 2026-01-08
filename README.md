# AI Concurrent Gateway

This project is a high-performance Go-based microservice that functions as a concurrent, asynchronous gateway to Large Language Model (LLM) providers like OpenAI. It is designed to handle multiple simultaneous requests efficiently using a worker pool pattern.

## Features

- **Concurrent Processing**: Manages a pool of workers to handle multiple requests in parallel.
- **Asynchronous Job Queue**: Requests are queued and processed asynchronously to prevent blocking.
- **Graceful Shutdown**: Ensures that the server shuts down cleanly, without interrupting active requests.
- **Centralized Configuration**: Environment variables are used for easy configuration management.
- **Structured Logging**: Provides detailed logs to trace the flow of each request.

## Prerequisites

- Go 1.22+ or higher
- An AI API Key (OpenAI or Gemini)

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ivan-salazar14/send-promt-ai.git
   cd send-promt-ai
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Configure environment variables:**
   Copy the example `.env` file and update it with your credentials:
   ```sh
   cp .env.example .env
   ```
   Edit the `.env` file to add your `AI_API_KEY` and other settings.

## Configuration

The application is configured via the following environment variables in the `.env` file:

- `PORT`: The port on which the server will listen (default: `8080`).
- `AI_API_KEY`: Your secret key for the OpenAI or Gemini API.
- `INTERNAL_AUTH_TOKEN`: A secret token for securing the API endpoint.
- `MAX_WORKERS`: The number of concurrent workers in the pool (default: `5`).
- `QUEUE_SIZE`: The capacity of the job queue (default: `100`).

## Running the Service

To start the server, run the following command:

```sh
go run ./cmd/main.go
```

The server will start on the configured port, and you should see logs indicating that it is ready to accept requests.

## API Usage

The service exposes a single `POST` endpoint for processing text prompts.

- **Endpoint**: `POST /process`
- **Header**: `Authorization: Bearer <your_internal_token>`
- **Body**: `{ "prompt": "Your prompt here" }`

### Example Request

You can use the provided test script to send multiple concurrent requests to the API:

```sh
./test_api.sh
```

This script will send five simultaneous requests to the `/process` endpoint and display the responses from the server.

## Design Decisions

- **The Dispatcher Location**: Placed in the application layer because the strategy for handling jobs is a business-logic requirement for performance.
- **Context Propagation**: All requests use context.Context to ensure that if a client cancels a request, the AI call is terminated immediately, saving costs and resources.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GPL-3.0 License.
