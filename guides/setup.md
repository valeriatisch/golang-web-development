# Setup
## Step 1: Download Go
Go to  [golang.org/dl/](https://golang.org/dl/)  and download the installer for your operating system. Open the downloaded file and follow the instructions. <br>On **Linux** extract the downloaded `.tar.gz` file.
```
tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```
## Step 2: Go Path
Set the `PATH` environment variable. <br>
To setup the `PATH` on **Windows** please follow this [tutorial](https://medium.com/@bhanotvardana/setting-up-golang-environment-on-windows-3d50c2dbffe7). <br>
On **macOS** and **Linux**:
1. Navigate to your shell configuration file in your home directory:
    - For Bash: `vim ~/.bash_profile` or `vim ~/.bashrc`
    - For Zsh: `vim ~/.zshrc`
2. Add the following line at the end of the file: 
   ```
   export PATH=$PATH:/usr/local/go/bin
   ```
3. If you followed the workspace setup instructions, also add the `$GOPATH/bin` directory to the `PATH` by adding the following line: 
   ```
   export PATH=$PATH:$GOPATH/bin
   ```
4. Save the file and reload the configuration.
   ```
   source ~/.your_shell_config
   ```
## Step 3: Verification
Open a new terminal session and verify the installation.
```
go version
```
## Step 4: IDE / VSCode
Download [VSCode](https://code.visualstudio.com/download) or another appropriate IDE like [GoLand](https://www.jetbrains.com/go/).
In **VSCode**, go to the Extensions Marketplace, search for the Go plugin and install it.
![[go-plugin.png|600]]
Also install the necessary Go tools that comes with it.
Open the Command Palette by pressing `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (macOS), type *"Go: Install/Update Tools"* and press `Enter`.
## Step 7: PostgreSQL
1. Go to [postgresql.org/download/](https://www.postgresql.org/download/) & follow the installation instructions for your operating system.
   - On Windows, you will need to set the port. It’s normally 5432.
   - You might need to setup PgAdmin.
   - On Linux, you can verify the installation with: `psql --version`
2. Connect to PostgreSQL through an interface or the terminal:
   - On Ubuntu: `sudo -u postgres psql`
   - On Mac: `psql postgres`
3. Create a database: `CREATE DATABASE dbname;`
4. Create a user: `CREATE USER yourusername WITH ENCRYPTED PASSWORD 'yourpassword';`
5. Grant privileges to the user: `GRANT ALL PRIVILEGES ON DATABASE dbname TO yourusername;`
## Step 6: Test
To test if everything is working correctly, create a project.
1. Create a new directory for your project. This directory is usually the workspace for your Go code, dependencies, and other project-related files.
   ```
   mkdir your_go_project
   cd your_go_project
   ```
2. Go uses modules to manage dependencies. Initialise a new Go module with your desired module path (a unique identifier for your Go module)
   ```
   go mod init module_path/your_go_project
   ```
   A `go.mod` file will be created in your project directory, which will be used to manage your project's dependencies.
3. Create a new file named `main.go` in your project directory and write a simple Hello World program.
   ```go
   package main
   
   import "fmt"
   
   func main() {
	   fmt.Println("Hello, world!")
   }
   ```
4. Compile your program.
   ```
   go build
   ```
   This will create an executable binary named `main` (or `main.exe` on Windows).
   To run your program, simply execute the binary:
   ```
   ./main
   ```
   Alternatively, you can use compile and run your program in a single step with:
   ```
   go run main.go
   ```
5. Test your db connection with the following code.
   First, install the pq driver: `go get github.com/lib/pq`.
   Adjust the connection string according to your parameters (port, user, dbname) in the code.
   ```go
   package main
   
   import (
	   "database/sql"
	   "fmt"
	   "log"
	   _ "github.com/lib/pq"
    )

    func main() {

        // Connection string
        psqlconn := "host=localhost port=5432 user=yourusername dbname=dbname sslmode=disable"
        
        // Open a connection to the database
        db, err := sql.Open("postgres", psqlconn)
        if err != nil {
            log.Fatal("Failed to open a DB connection: ", err)
        }
        defer db.Close()
        
        // Ping the database to verify the connection
        err = db.Ping()
        if err != nil {
            log.Fatal("Failed to connect to the DB: ", err)
        }
        
        fmt.Println("Successfully connected!")
    }
    ```
Run the code: `go run main.go`. It should print *“Successfully connected”*.