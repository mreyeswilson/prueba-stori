# Stori Test Project

## Overview

This README provides a comprehensive guide to setting up and running the application, including installation instructions for AWS SAM, AWS CLI, Go, and Mockery. It also explains the commands available in the Makefile for managing the project.

## Prerequisites

Before you can run the application, you need to install several tools:

- **AWS SAM CLI**: A command-line tool that simplifies the process of managing serverless applications.
- **AWS CLI**: A command-line tool to interact with AWS services.
- **Go**: The Go programming language.
- **Mockery**: A mock code generation tool for Go.

## Installation

### 1. Installing AWS SAM CLI

AWS SAM (Serverless Application Model) CLI is a command-line tool that simplifies the development of serverless applications. To install it, follow these steps:

#### For macOS

1. **Install Homebrew** if you haven't already:

    ```bash
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

2. **Install AWS SAM CLI**:

    ```bash
    brew tap aws/tap
    brew install aws-sam-cli
    ```

#### For Windows

1. **Install Chocolatey** if you haven't already:

    Open an administrative PowerShell window and run:

    ```powershell
    Set-ExecutionPolicy Bypass -Scope Process -Force
    iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
    ```

2. **Install AWS SAM CLI**:

    ```powershell
    choco install aws-sam-cli
    ```

#### For Linux

1. **Download the AWS SAM CLI binary**:

    ```bash
    curl "https://d1uj6qtbmh3dt5.cloudfront.net/aws-sam-cli-linux-x86_64.zip" -o "sam-cli-linux-x86_64.zip"
    ```

2. **Unzip the package and install**:

    ```bash
    unzip sam-cli-linux-x86_64.zip
    sudo ./install
    ```

3. **Verify the installation**:

    ```bash
    sam --version
    ```

### 2. Installing AWS CLI

AWS CLI is a command-line tool that allows you to interact with AWS services.

#### For macOS

1. **Install Homebrew** if you haven't already:

    ```bash
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    ```

2. **Install AWS CLI**:

    ```bash
    brew install awscli
    ```

#### For Windows

1. **Download the AWS CLI installer** from the [official AWS website](https://aws.amazon.com/cli/).

2. **Run the installer** and follow the instructions.

#### For Linux

1. **Install using the package manager**:

    ```bash
    sudo apt-get update
    sudo apt-get install awscli
    ```

2. **Verify the installation**:

    ```bash
    aws --version
    ```

### 3. Configuring AWS CLI Profile

To interact with AWS services and deploy your application, you need to configure your AWS CLI profile. Follow these steps to set it up:

1. **Run the AWS Configure Command**:

    ```bash
    aws configure
    ```

2. **Provide Your AWS Credentials**:

    You will be prompted to enter the following information:

    - **AWS Access Key ID**: Your AWS access key.
    - **AWS Secret Access Key**: Your AWS secret key.
    - **Default region name**: The AWS region you want to use (e.g., `us-east-1`).
    - **Default output format**: The format for output (e.g., `json`).

   These credentials are saved in the `~/.aws/credentials` and `~/.aws/config` files.

### 4. Installing Go

Go is a statically typed, compiled programming language designed at Google.

#### For macOS and Linux

1. **Download the Go binary**:

    ```bash
    wget https://golang.org/dl/go1.20.3.darwin-amd64.tar.gz
    ```

2. **Extract and install Go**:

    ```bash
    sudo tar -C /usr/local -xzf go1.20.3.darwin-amd64.tar.gz
    ```

3. **Add Go to your PATH**:

    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

4. **Verify the installation**:

    ```bash
    go version
    ```

#### For Windows

1. **Download the Go installer** from the [official Go website](https://golang.org/dl/).

2. **Run the installer** and follow the instructions.

3. **Verify the installation**:

    ```powershell
    go version
    ```

### 5. Installing Mockery

Mockery is a tool for generating mock objects in Go.

1. **Install Mockery**:

    ```bash
    go install github.com/vektra/mockery/v2@latest
    ```

2. **Add the Go binary path to your PATH**:

    ```bash
    export PATH=$PATH:$(go env GOPATH)/bin
    ```

3. **Verify the installation**:

    ```bash
    mockery --version
    ```

## Makefile Commands

The `Makefile` provides several commands for managing the project. Below are the commands and their descriptions:

- **build**: Builds the project using the AWS SAM CLI. This command packages your Lambda functions and prepares them for deployment.

- **start**: Builds the project (using the `build` command) and starts a local API gateway for testing. The local API runs on port 5000.

- **coverage**: Runs tests and generates a code coverage report. It also generates an HTML report for viewing test coverage.

- **gen-mocks**: Generates mock implementations for interfaces in the specified directory using Mockery. The mocks are saved in the `mocks` directory.

- **gofmt**: Formats all Go source files in the project using `gofmt`, ensuring consistent code formatting.

- **deploy**: Formats Go files (using the `gofmt` command), builds the project (using the `build` command), and deploys it to AWS using SAM. It specifies the stack name, AWS region, S3 bucket, and other deployment options.

## Running the Application

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/your-repo/project.git
    cd project
    ```

2. **Install Dependencies**:

    Navigate to your project directory and run:

    ```bash
    go mod tidy
    ```

3. **Build the Project**:

    ```bash
    go build -o your-app
    ```

4. **Deploy the Application Using SAM**:

    ```bash
    sam build
    sam deploy --guided
    ```

5. **Run Unit Tests**:

    To run tests for your Go application, use:

    ```bash
    go test ./...
    ```

6. **Generate Mocks**:

    Use Mockery to generate mocks for your interfaces:

    ```bash
    mockery --all
    ```

## Generating SES Template

To create the required SES (Simple Email Service) template, you need to use the AWS CLI. Follow these steps:

1. **Ensure you have the `template.json` file**: This file, located at the root of the project directory, contains the template definition for SES.

2. **Run the following AWS CLI command** to create the SES template:

    ```bash
    aws ses create-template --cli-input-json file://template.json
    ```

   This command uses the `template.json` file to define the email template in SES.

## Uploading CSV to S3 Bucket

For the application to function correctly, you need to upload your CSV file to the specified S3 bucket. Follow these steps:

1. **Upload the CSV file** to the following public S3 bucket with read and write permissions:

    ```text
    https://prueba-stori-bucket.s3.amazonaws.com
    ```

2. **Ensure the file is accessible**: The bucket is public and allows both reading and writing, so make sure your CSV file is correctly uploaded.

   You can use the AWS CLI or the AWS Management Console to upload the file. For example, using the AWS CLI:

    ```bash
    aws s3 cp your-file.csv s3://prueba-stori-bucket/your-file.csv
    ```

## Additional Information

- **AWS SAM Documentation**: [AWS SAM CLI Documentation](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli.html)
- **AWS CLI Documentation**: [AWS CLI Documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)
- **Mockery Documentation**: [Mockery GitHub Repository](https://github.com/vektra/mockery)
- **Go Documentation**: [Go Documentation](https://golang.org/doc/)

Feel free to reach out if you have any questions or run into any issues.
