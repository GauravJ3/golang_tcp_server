This is a simple HTTP server written in Go that listens for incoming TCP connections on port 1729 and responds with "Hello, World". The server uses the `net` package from the Go standard library.

## Usage

To run the server, execute the following command:

```shell
go run main.go
```
Once the server is running, you can hit it using curl from your terminal:

```
curl http://localhost:1729
Make sure to use the -UseBasicParsing flag when running curl.
```
## Explanation
The server implementation consists of the following components:

### The do function: This function handles an incoming connection and sends back the "Hello, World" response.
### The main function: This is the entry point of the program. It listens for incoming TCP connections on port 1729 and spawns a separate goroutine to handle each connection.

## Contributing
If you find any issues or have improvements to suggest, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
