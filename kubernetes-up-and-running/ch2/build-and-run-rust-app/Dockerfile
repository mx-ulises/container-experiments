# Use a lightweight Alpine image
FROM alpine:latest

# Install the Rust compiler and Cargo
RUN apk add --no-cache rust cargo

# Copy the project files to the container
COPY . .

# Build the project
RUN cargo build --release

# Expose the port used by the server
EXPOSE 8080

# Start the server
CMD ["./target/release/hello_world"]
