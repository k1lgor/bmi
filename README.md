# BMI Calculator CLI

This is a command-line interface (CLI) tool for calculating body mass index (BMI) based on user input for weight and
height.

## Features

- Calculate BMI based on weight (in kilograms) and height (in centimeters).
- Display the calculated BMI value and corresponding weight category (underweight, normal, overweight, or obese).
- Allow for input of weight and height either through command-line arguments.
- Validate user input to ensure that weight and height are valid numeric values.

## Usage

To calculate your BMI using this CLI tool, run the following command:

```go
bmi <weight> <height>
```

Replace <weight-in-kg> with your weight in kilograms (e.g., 70) and <height-in-m> with your height in meters (e.g.,
175).

## Installation

To install the BMI Calculator CLI tool, simply download the executable file for your operating system from the releases
page on GitHub.

Alternatively, you can clone the repository and build the executable file yourself using Go:

```go
git clone git@github.com:k1lgor/bmi.git
cd bmi
go build
sudo mv bmi /usr/bin
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.