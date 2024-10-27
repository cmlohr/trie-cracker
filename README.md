# Trie-Cracker - WIP
### !! **Work in Progress**: This project is in its early stages and not yet fully functional !!


### Overview

`Trie-Cracker` is an experimental implementation of a password-cracking tool focused on optimizing data structure and algorithm efficiency. It explores the use of a **trie (prefix tree)** structure to speed up password searches when certain constraints (such as known starting characters) are available. The primary goal is to demonstrate how custom data structures can refine traditional brute-force and dictionary-based approaches for faster, targeted cracking. Written in **Go** for initial prototyping, the project balances performance with flexibility.

### Key Features

-   **Trie-Based Structure**: Efficiently stores and retrieves potential password candidates by character, enabling rapid path pruning based on known constraints.
-   **Configurable Constraints**: Experiment with starting character constraints or known patterns within the password to minimize unnecessary search paths.
-   **Modular Wordlist Loading**: Easily import different wordlists (e.g., `.txt` files) to test cracking speed across various dictionaries.
-   **MD5 Hash Matching**: Initial hashing is based on MD5 for testing purposes, but future extensions may include support for SHA, bcrypt, etc.

### Future Plans

-   **Advanced Constraints**: Develop targeted rules for character positions (e.g., digits in specific positions), allowing further pruning of the trie.
-   **Parallel Processing**: Take advantage of Go's concurrency model to distribute cracking tasks across multiple cores.
-   **Zig Optimization**: Once stable, consider rewriting in Zig to leverage lower-level control for memory optimization and raw speed improvements.

### Getting Started

#### Prerequisites

-   **Go**: Version 1.16 or higher

#### Installation

Clone the repository:

bash

Copy code

`git clone https://github.com/cmlohr/trie-cracker.git
cd trie-cracker` 

Install dependencies:

bash

Copy code

`go mod tidy` 

#### Usage

1.  **Add Wordlist**: Place your password wordlist in the `data/` directory as a `.txt` file. A default wordlist of 1,000 entries is recommended for initial testing.
2.  **Compile the Program**:
    
    bash
    
    Copy code
    
    `go build -o trie-cracker main.go` 
    
3.  **Run the Program**:
    
    bash
    
    Copy code
    
    `./trie-cracker --wordlist data/wordlist.txt --hash <target-hash>` 
    
    Replace `<target-hash>` with the MD5 hash you want to crack.

### Directory Structure

-   `data/`: Contains wordlists used for testing.
-   `src/`: Holds source code, including trie structure, hashing functions, and search algorithms.
-   `docs/`: Additional documentation, including references to Go libraries or Zig if later incorporated.

### Contributing

This repo is open for contributions and feedback. Feel free to fork, experiment, and suggest new methods for trie optimization or hashing techniques.

### License

MIT License
