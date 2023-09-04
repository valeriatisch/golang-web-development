# Structuring Go Projects 🏗️

This document outlines some best practices for structuring GoLang projects, focusing on APIs.
Remember, these are just some **suggestions**, not stone tablets from Mount Gopher, and may therefore not be suitable for all cases.
Adapt as you see fit.

## Principles ✨

- **Keep it simple**: Avoid over-engineering. Creating unnecessary folders can lead to circular dependencies. It's easier to scale folders up over time than to scale them down.
- **Be consistent:** Stick to coding standards and conventions.
- **Write readable and simple code:** Code is more often read than written. Don't make it cryptic. Making your code clear and maintainable should be a priority over making it super "efficient," especially because Go is already inherently fast.

## Project Structure 📁

### Nr. 1 - The Cozy Structure 🏡

Here's a proposed structure for a simple GoLang project at the beginning of its lifecycle, that has just what you need:

```
.
├── api
│   ├── server.go
│   ├── server_test.go
│   └── handlers.go (optional)
├── storage
│   ├── storage.go
│   ├── mongodb.go
│   └── memory.go
├── types
│   ├── user.go
│   └── book.go
├── util
│   └── util.go
├── main.go
└── go.mod, go.sum
```

#### main.go 🚪

This is the entry point of your application, like the door to your house!
Here you can define flags, create instances of your server and storage, and start your server. 🚀

For example, something like this:

```go
var listenAddr = flag.String("address", "localhost:8080", "server listen address")
var storage = storage.NewMemoryStorage()

func main() {
    flag.Parse()
    server := api.NewServer(*listenAddr, storage)
    log.Fatal(server.Start())
}
```

#### api 🌐

This folder is where all your HTTP handlers live. Think of it as a kitchen, that's where all the magic happens. 🍳✨

##### server.go 🍽️

Here we define our server type and methods associated with it. It might be like a recipe. An instance of the server is then created in `main.go`.

```go
type Server struct {
    listenAddr string
    store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
    return &Server{listenAddr, store}
}

func (s *Server) Start() error {
    http.HandleFunc("/user", s.handleGetUserByID)
    return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) { }
```

Handlers can either be defined in the `server.go` file, or in a separate `handlers.go` file. If you have a lot of handlers, it's better to separate each of them into a separate file, for example `user_handler.go` and `book_handler.go`.

#### storage 🗄️

Welcome to the pantry! This folder is for your data layer. Always define an interface for your storage, and then implement this interface for each type of storage you have (e.g., MongoDB, in-memory for testing, etc.)

###### storage.go 📦

Define a `Storage` interface that all storage types must implement.

```go
type Storage interface {
    Get(id int) types.User
}
```

##### mongodb.go / memory.go 🗂️

Exemplary implementations of the `Storage` interface.

```go
type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{}
}

func (ms *MemoryStorage) Get(id int) types.User {
    return types.User{ID: 1, Name: "Foo"}
}
```

#### types 📝

Define all the types that you will use across your project here. ✍️

##### user.go / book.go

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Book struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
}
```

#### util 🛠️

The toolbox! This folder is for utility functions that you might use across your project. This can include functions for logging, error handling, etc.

#### Testing 🧪

Create test files in the same package as the code you're testing. This keeps the tests close to the code they're testing and follows common Go idioms.

### Nr. 2 - Extended Structure 🏰

If you're working on a more complex project, you might want to consider an extended and advanced structure. This is especially true if you're working on a project with multiple apps, or if the project will be used by multiple teams.

Here's a proposed structure for such a project:

```
├── cmd
│   ├── app1
│   │   └── main.go
│   └── app2
│       └── main.go
├── api
│   ├── handler
│   │   ├── user.go
│   │   └── book.go
│   └── middleware
│       └── auth.go
├── pkg
│   ├── models
│   │   ├── user.go
│   │   └── book.go
│   ├── storage
│   │   ├── db.go
│   │   ├── mongo.go
│   │   └── memory.go
│   └── util
│       └── util.go
├── tests
└── go.mod, go.sum
```

#### cmd 🏠

This folder houses your applications. Each app should have its own subdirectory, and each should contain a `main.go` file.

For example, you could have a user app, an admin app, a CLI app, etc.

#### api 🌐

Similar to the simple structure, this is where your HTTP or API handlers and middleware reside. But, now they have their own dedicated folders.

##### handler/user.go, handler/book.go 🖐️

This is where you implement the HTTP handlers, each in its own file.

##### middleware/auth.go 🛡️

Here, you write your middleware, each having its own file. This can include authorization, logging, or any other cross-cutting concerns.

#### pkg 📦

This is where the shared code lives. Anything that could be used by multiple applications belongs here.

##### models/user.go, models/book.go 📚

Just as in the `types` folder in the simple structure, define your types here.
Unlike the `models` folder in some other languages or frameworks, where data fetching, manipulation, and storage usually occur, in Go, it mainly is for describing the structure of our data (structs and their methods).

##### storage/db.go, storage/mongo.go, storage/memory.go 🗄️

These are your storage interfaces and implementations.

#### util 🛠️

This folder is for utility functions that could be used across your project.

#### tests 🧪

Unlike the simple structure where tests were alongside the code, here, tests have their own directory. This makes it easier to run all tests at once.

## Makefile 🔧

In a larger Go project, it's useful to introduce automation to handle common tasks, like building your application, running tests, or setting up your development environment.
That's where `Makefile` comes in.
A `Makefile` is a special file, residing in the root directory, that contains shell commands for managing your project.
You can think of it as an equivalent to the scripts section of a package.json file.

Here's a simple example:

```make
.PHONY: build test clean

build:
    go build -o bin/myapp cmd/myapp/main.go

test:
    go test -v ./...

clean:
    rm -rf bin/
```

- `build` command compiles your Go files and puts the output in the `bin/` directory.
- `test` command runs all your tests.
- `clean` command removes the `bin/` directory. 🧹

The `bin/` directory, usually not committed into source control (add it to your `.gitignore`), is where you store the compiled binary files.

To run a command, type `make <command>`. For example, to run the `build` command, type `make build`.

## Inspiration: Masterpieces in Go 🌟

It's always a good idea to study the greats! Take a look at these projects that have implemented best Go practices.

1. [WTF](https://github.com/benbjohnson/wtf)
2. [Kubernetes](https://github.com/kubernetes/kubernetes)
3. [InfluxDB](https://github.com/influxdata/influxdb)
