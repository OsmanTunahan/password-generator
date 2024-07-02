## Password Generator

A simple password generator. It generates a random password based on a specified length.

### Installation

1. Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
2. Clone the repository or download the source code files.
3. Navigate to the directory containing the source code.

### Usage

To build and run the password generator, follow these steps:

1. Open a terminal and navigate to the project directory.
2. Build the application:
   ```sh
   go build -o password-generator
   ```
3. Run the application:
   ```sh
   ./password-generator [length]
   ```
   - `[length]` (optional): The length of the password to be generated. If not provided, the default length of 12 characters will be used.

#### Examples

- Generate a password with the default length:
  ```sh
  ./password-generator
  ```
  Output: `aBc123dEfGhI`

- Generate a password with a specified length (e.g., 16 characters):
  ```sh
  ./password-generator 16
  ```
  Output: `1aBc2dEf3GhIjK4L`

### Configuration

You can modify the default length and character set used for password generation by editing the `Config` struct in the `main.go` file:
```go
config := Config{
	DefaultLength: 12,
	CharSet:       "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}
```

### License

This project is open source and available under the [MIT License](LICENSE).