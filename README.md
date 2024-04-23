# CS4390 Project README

## Overview

This README provides instructions on how to build and run the `CS4390_Project`, which consists of a client and server application. The client and server are written in Go, and the eval is a Python script.

## Prerequisites

Before you begin, ensure you have the following installed on your system:
- Go (version 1.x.x or later)
- Python (version 3.x.x or later)

## Getting Started

To get started with the `CS4390_Project`, clone the repository to your local machine using the following command:

```bash
git clone https://github.com/your-username/CS4390_Project.git
cd CS4390_Project
```

## Building the Project
You can build both the client and server applications using the provided Makefile.

To build the client and server, run:

```bash
make
```

## Running the Applications

To run the server application, use the following command:

```bash
make run-server
```
The server will start, and you should see output indicating that it is listening for connections or processing data.

To run the client application, you need to pass the name argument to the make run-client command. For example:

```bash
make ARGS="your_name" run-client
```

## Cleaning up
To clean up the binaries and any other files generated during the build, run:

```bash
make clean
```
