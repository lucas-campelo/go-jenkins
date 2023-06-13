
# Go Jenkins: An Easy Setup Tool for Jenkins

An easy setup tool for Jenkins.

## Introduction

Go Jenkins is a user-friendly program designed to simplify the process of setting up Jenkins on your Linux machine. This documentation provides step-by-step instructions to help you install and configure Jenkins using the Jenkins Docker installation method.

## Installation

To install Go Jenkins, follow these steps:

1. Build the application by executing the following command in your terminal:

```sh
go build -o go-jenkins
```

This command compiles the Go code and generates the executable file for Go Jenkins.

2. Then sudo-move it to your local binaries folder, so it can be avaiable anywhere on the system:

```sh
sudo mv go-jenkins /usr/local/bin
```

## References

For more details on the installation process, refer to the [Jenkins Docker installation](https://www.jenkins.io/doc/book/installing/docker/).