## [Go](https://go.dev/) - The language used in this tutorial
  - ### Getting started:
    1. take [A Tour of Go](https://go.dev/tour/welcome/1)
    2. [Installation Instructions](https://go.dev/doc/install) vary by machine
  - ### Starting your first go project
  - Open a terminal in your root directory (keep it open), run: 
    ```
    go mod init <your-project-name>
    ```
    This creates a space for Go to store all project dependencies
  - Create a new file called "bot.go":
    ```
    touch bot.go
    ```
  - Say hello world
    ```go
    package main

    import (
      "fmt"
    )

    func main(){
      fmt.Println("Hello World")
    }
    ```
Now that Go has been successfully installed, and you have taken [A Tour of Go](https://go.dev/tour/welcome/1), we are ready to install our first package!