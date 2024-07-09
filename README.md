# BooKing

Welcome to **BooKing**, a powerful and efficient booking system built with Golang. BooKing is designed to handle your booking needs with simplicity and speed, ensuring a seamless experience for both administrators and users.

## Features

- **Robust Booking Management**: Easily create, update, and manage bookings.
- **User-Friendly Interface**: Clean and intuitive API for smooth integration.
- **Scalable Performance**: Built with Golang for high performance and scalability.
- **Secure and Reliable**: Implements best practices to ensure data security and integrity.

## Getting Started

### Prerequisites

- [Golang](https://golang.org/doc/install) 1.18 or higher
- [MongoDB](https://www.mongodb.com/) database

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/dmarquinah/boo-king.git
    cd boo-king
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up your database:**

    Ensure your database is running and update the `.env` file with your database credentials.

    ```sh
    MONGO_URI=
    ```

4. **Run the application:**

    ```sh
    go run cmd/api/main.go
    ```

## Contributing

We welcome contributions! Please fork the repository and submit a pull request for any enhancements or bug fixes.

---

Feel free to reach out if you have any questions or need further assistance.

Happy booking!
