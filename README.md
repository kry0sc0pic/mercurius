```markdown
# YourProject

YourProject is a commandline tool built with Go. It allows you to run a long-running command and get a notification on Telegram, Whatsapp, or Discord when the command completes or errors out.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Docker (optional)

### Installing

1. Clone the repository

```bash
git clone https://github.com/yourusername/yourproject.git
```

2. Navigate to the project directory

```bash
cd yourproject
```

3. Install the dependencies

```bash
go mod download
```

4. Copy the `.env.example` file to `.env` and fill in your configuration details

```bash
cp .env.example .env
```

5. Build the project

```bash
go build -o yourproject
```

## Usage

To run a command and get a notification when it completes or errors out, use the `run` command followed by the command you want to run. For example:W

```bash
./mercuriys run "sleep 30"
```

This will run the `sleep 30` command and send a notification when it completes or if it errors out.


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](LICENSE.md) file for details
```
