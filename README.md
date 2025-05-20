# Blockchain Implementation in Go

A foundational blockchain implementation written in Go, designed for educational purposes and to demonstrate core blockchain concepts such as block creation, proof-of-work, and wallet management.

## Features

- **Block Creation**: Ability to create and add new blocks to the blockchain.
- **Proof-of-Work**: Implements a basic proof-of-work consensus mechanism.
- **Wallet Management**: Includes wallet generation and storage functionality.

## Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system.

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/abhishek-tiwari-016/Blockchain.git
   cd Blockchain

2. **Build the application**:
   
   ```bash
   go build -o blockchain

4. **Run the application:**
   
   ```bash
   ./blockchain

### Project structure

Blockchain/
├── main.go          # Entry point of the application
├── go.mod           # Go module file
├── go.sum           # Go module checksums
├── wallet.dat       # Wallet data file
├── utils/           # Utility functions and helpers
└── README.md        # Project documentation
