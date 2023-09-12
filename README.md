# dailyQuote

`dailyQuote` is a simple and effective quote generator that provides you with a daily dose of inspiration and motivation. Built in Go, it currently prints out quotes via the `main.go` file.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Upcoming Features](#upcoming-features)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/zajicekn/dailyQuote.git
```

2. Navigate to the project directory:

```bash
cd dailyQuote
```

3. Install the necessary dependencies:
require github.com/joho/godotenv v1.5.1

go get ./...
```bash
go get github.com/joho/godotenv
```

## Usage

To run the quote generator:
```bash
go run main.go
```

Upon executing the above command, either a successful message or error will be printed to the screen.

## Messaging System

I have integrated with Telegram to provide the messaging functionality. After setting up a bot, the program can now relay messages to the Telegram app. This feature is currently in beta and only supports Telegram, but I'm looking into expanding this capability for other messaging platforms in the future. 

Plans are in place to host the application on a cloud server to ensure uninterrupted service. We're also keen on adding a feedback system where users can indicate whether they liked the provided quote or not.

## Upcoming Features

- **Scalability**: I'm researching ways to scale the Telegram bot feature so that other users can easily use the bot and receive messages.
- **Cloud Hosting**: The plan is to host the application on a cloud server to ensure it runs continuously.
- **User Feedback**: I'm considering introducing a feature that allows users to give feedback on the quotes they receive.
- **Additional Customizations**: I'm exploring avenues to let users personalize their daily quotes experience further.

## Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-YourFeature`
3. Commit your changes: `git commit -am 'Added some feature'`
4. Push to the branch: `git push origin feature-YourFeature`
5. Submit a Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE.md file in the repository for more details.
