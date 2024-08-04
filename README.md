# ASCII-Art-Web

This project consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of my previous project, ASCII-Art.

The technologies used to build the project include Golang for the web server operations, HTML, CSS and JavaScript for the frontend and containerized with Docker.

## Project Preview

![](./img/preview)

## Getting Started
To get started with the project, download and install Go programming language using the link: https://go.dev/doc/install

## Installation
1. Clone the repository from its remote location to your local environment with the command below:
```bash
git clone https://github.com/johneliud/ASCII-Art-Web.git
```

2. Navigate to the project's path
```bash
cd Ascii-Art-Web
```

## Usage
Once in the correct project path, the user needs to specify the port on which the web server will run on via the terminal incase they don't want to run on the server's default port `8080`.

NOTE: The specified port to run the server should be a number ranging between `1024 to 65535`. Any other number not within this range has been reserved and won't work for this project.

Run the program with either of the commands below:
```bash
go run .
```

or

```bash
ro run . [SPECIFY-PORT]
```

From this point, the user can interact with the program in a variety of ways including:

- Providing input using the textarea

- Selecting their preferred banner file

- Submitting their input for processing

## Implementation
The project's implementation uses the same implementation of the original ascii-art project. This therefore affects how input passed by the user is displayed back to them in the mentioned cases below:

Occurrences of "\\a", "\\b", "\\v", "\\f", "\\r" or any other character absent in the range of 32 to 126 in the [ASCII manual](https://man.archlinux.org/man/core/man-pages/ascii.7.en) present in the input will result to a bad request error displayed to the user. This also applies for emoji characters.

### HTTP Endpoints
The project uses an HTTP method GET to retrieve users input and uses HTTP method POST to display information from the web server back to the user using the ascii-art URL. Incase of any endpoint failure, an HTTP status code will be displayed to the user. Below are some error codes and what may result them:

**400**: Signifies the input being processed is invalid.

**404**: Signifies the path being accessed is invalid.

**405**: Signifies the HTTP method being accessed is not allowed.

**500**: Signifies the error being experienced by the user results from an internal error of the program.

## Contact
Feel free to reach out incase of any enquiry in regards to this project via the email address johneliud4@gmail.com
