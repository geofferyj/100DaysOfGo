# Mr. J’s 100DaysOfCode Challenge - Go Edition

# Curriculum

Day 1: Introduction to Programming, Golang, and Setting Up Your Go Environment

* Introduction to programming concepts
* Introduction to Golang 
    * Resource:[ A Tour of Go](https://tour.golang.org/welcome/1)
* Setting up your Go environment
    * Resource:[ Getting Started](https://golang.org/doc/install)	
* Exercise: Set up your Go environment and run a simple "Hello, World!" program.

Day 2 (28 Apr 2023): Hello World, Basic Syntax, Variables & Constants



* Hello World program
* Basic syntax
    * Resource:[ A Tour of Go - Basics](https://tour.golang.org/basics/1)
* Variables and constants
    * Resource:[ Go by Example - Variables](https://gobyexample.com/variables)
* Exercise: Write a program that reads your name from the command line and greets you with "Hello, [Your Name]!".

Day 3: Basic Data Types, Operators and Type Conversion



* Basic data types
    * Resource:[ Go by Example - Basic Types](https://gobyexample.com/basic-types)
* Operators and type conversion
    * Resource:[ Go by Example - Operators](https://gobyexample.com/operators)
* Exercise: Write a program that reads two integers variables and prints their sum, difference, product, and quotient

Day 4: Strings



* Resource: [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
* Resource: [strings - Go Packages](https://pkg.go.dev/strings)
* Exercise: Write a program that converts a string to uppercase and replaces a specified substring with another substring.

Day 5: Output Formatting



* Resource: [fmt - Go Packages](https://pkg.go.dev/fmt)
* Exercise: Write a program that reads two integers and prints their sum, difference, product, and quotient with a formatted output.

Day 6: Inputs 



* Resource: 
    * [fmt - Go Packages](https://pkg.go.dev/fmt)
    * [bufio - Go Packages](https://pkg.go.dev/bufio@go1.20.3)
* Exercise: Write a program that reads a line of text from the user and prints the number of words in the input string.

	

Day 7: Control Structures: If and Else



* Resource:[ Go by Example - If/Else](https://gobyexample.com/if-else)
* Exercise: Write a program that finds the largest of three numbers using if-else statements.

Day 8: Control Structures: Switch



* Resource:[ Go by Example - Switch](https://gobyexample.com/switch)
* Exercise: Write a program that converts a numerical grade to a letter grade using switch statements.

Day 9: Control Structures: For Loops



* Resource:[ Go by Example - For](https://gobyexample.com/for)
* Exercise: Write a program that prints the first n Fibonacci numbers.

Day 10: Control Structures: Range



* Resource:[ Go by Example - Range](https://gobyexample.com/range)
* Exercise: Write a program that calculates the sum of elements in an integer slice.

Day 11: Functions



* Resource:[ Go by Example - Functions](https://gobyexample.com/functions)
* Exercise: Implement a function to calculate the factorial of a given number.

Day 12: Multiple Return Values



* Resource:[ Go by Example - Multiple Return Values](https://gobyexample.com/multiple-return-values)
* Exercise: Write a function that returns the minimum and maximum elements of an integer slice.

Day 13: Variadic Functions



* Resource:[ Go by Example - Variadic Functions](https://gobyexample.com/variadic-functions)
* Exercise: Implement a variadic function that returns the average of a series of integers.

Day 14: Closures



* Resource:[ Go by Example - Closures](https://gobyexample.com/closures)
* Exercise: Write a closure to create a sequence of numbers.

Day 15: Recursion



* Resource:[ Go by Example - Recursion](https://gobyexample.com/recursion)
* Exercise: Implement a recursive function to find the nth Fibonacci number.

Day 16: Methods



* Resource:[ Go by Example - Methods](https://gobyexample.com/methods)
* Exercise: Implement a method for a "Circle" struct that calculates the area.

Day 17: Interfaces



* Resource:[ Go by Example - Interfaces](https://gobyexample.com/interfaces)
* Exercise: Create an interface "Shape" with a method "Area()" and implement it for different structs (e.g., "Circle" and "Rectangle").

Day 18: Errors



* Resource:[ Go by Example - Errors](https://gobyexample.com/errors)
* Exercise: Implement a custom error type to handle division by zero in a division function.

Day 19: Error Handling



* Resource:[ Error Handling in Go](https://blog.golang.org/error-handling-and-go)
* Exercise: Refactor the division function to handle errors using multiple return values.

Day 20: Custom Error Types



* Resource:[ Creating Custom Error Types in Go](https://dave.cheney.net/2014/12/24/inspecting-errors)
* Exercise: Create a custom error type to handle file-not-found errors in a file reading function.

Day 21: Structs



* Resource:[ Go by Example - Structs](https://gobyexample.com/structs)
* Exercise: Create a "Student" struct with methods for adding and retrieving grades.

Day 22: Pointers



* Resource:[ Go by Example - Pointers](https://gobyexample.com/pointers)
* Exercise: Implement a function that swaps the values of two integer variables using pointers.

Day 23: Arrays



* Resource:[ Go by Example - Arrays](https://gobyexample.com/arrays)
* Exercise: Write a program that reads an array of integers and reverses it.

Day 24: Slices



* Resource:[ Go by Example - Slices](https://gobyexample.com/slices)
* Exercise: Implement a function that removes a specific element from a slice.

Day 25: Maps



* Resource:[ Go by Example - Maps](https://gobyexample.com/maps)
* Exercise: Write a program that counts the frequency of words in a string using a map.

Day 26: Composition & Embedding (continued)



* Resource:[ Composition in Go](https://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html)
* Exercise: Implement a "Person" and an "Employee" struct using composition.

Day 27: Goroutines



* Resource:[ Go by Example - Goroutines](https://gobyexample.com/goroutines)
* Exercise: Write a program that calculates the factorial of a number concurrently using goroutines.

Day 28: Channels



* Resource:[ Go by Example - Channels](https://gobyexample.com/channels)
* Exercise: Modify the concurrent factorial program to use channels for communication.

Day 29: Channel Synchronization



* Resource:[ Go by Example - Channel Synchronization](https://gobyexample.com/channel-synchronization)
* Exercise: Write a program that uses channels to synchronize the completion of multiple goroutines.

Day 30: Channel Buffering



* Resource:[ Go by Example - Channel Buffering](https://gobyexample.com/channel-buffering)
* Exercise: Implement a buffered channel to simulate a producer-consumer scenario.

Day 31: Channel Select



* Resource:[ Go by Example - Channel Select](https://gobyexample.com/select)
* Exercise: Write a program that uses the select statement to read from multiple channels.

Day 32: Timeouts and Channels



* Resource:[ Go by Example - Timeouts](https://gobyexample.com/timeouts)
* Exercise: Implement a timeout for a slow-running goroutine using a time.After channel.

Day 33: Non-blocking Channel Operations



* Resource:[ Go by Example - Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations)
* Exercise: Modify a channel-based program to use non-blocking operations with the select statement.

Day 34: Closing Channels



* Resource:[ Go by Example - Closing Channels](https://gobyexample.com/closing-channels)
* Exercise: Implement a program that uses a closed channel to signal the end of a sequence of Fibonacci numbers.

Day 35: Range and Channels



* Resource:[ Go by Example - Range over Channels](https://gobyexample.com/range-over-channels)
* Exercise: Write a program that calculates a series of Fibonacci numbers and sends them over a channel. The receiver should use range to iterate over the numbers.

Day 36: Timers



* Resource:[ Go by Example - Timers](https://gobyexample.com/timers)
* Exercise: Write a program that schedules an event using a timer.

Day 37: Tickers



* Resource:[ Go by Example - Tickers](https://gobyexample.com/tickers)
* Exercise: Write a program that periodically sends a message to a channel using a ticker.

Day 38: Packages



* Resource:[ A Tour of Go - Packages](https://tour.golang.org/basics/1)
* Exercise: Create a simple package to handle arithmetic operations and import it into your main program.

Day 39: Modules and Dependency Management



* Resource:[ Go Modules Reference](https://golang.org/ref/mod)
* Exercise: Set up a new Go project with module support and add a third-party dependency.

Day 40: Testing



* Resource:[ Writing Tests in Go](https://gobyexample.com/testing)
* Exercise: Write tests for your arithmetic operations package.

Day 41: Benchmarking



* Resource:[ Go by Example - Benchmarking](https://gobyexample.com/benchmarking)
* Exercise: Write benchmarks for your arithmetic operations package.

Day 42: Debugging



* Resource:[ Debugging Go Code](https://golang.org/doc/gdb)
* Exercise: Debug a sample Go program to fix an issue.

Day 43: Defer, Panic, and Recover



* Resource:[ Go by Example - Defer, Panic, Recover](https://gobyexample.com/defer-panic-recover)
* Exercise: Implement a function that uses defer and recover to handle a division by zero panic.

Day 44: HTTP Clients



* Resource:[ Go by Example - HTTP Clients](https://gobyexample.com/http-clients)
* Exercise: Write a program that makes a simple HTTP GET request to a REST API.

Day 45: HTTP Servers



* Resource:[ Go by Example - HTTP Servers](https://gobyexample.com/http-servers)
* Exercise: Write a simple HTTP server that responds to requests with "Hello, World!".

Day 46: JSON



* Resource:[ Go by Example - JSON](https://gobyexample.com/json)
* Exercise: Write a program that reads a JSON file and decodes it into a struct.

Day 47: XML



* Resource:[ Go by Example - XML](https://gobyexample.com/xml)
* Exercise: Write a program that reads an XML file and decodes it into a struct.

Day 48: File I/O



* Resource:[ Go by Example - Reading Files](https://gobyexample.com/reading-files)
* Exercise: Write a program that reads a text file line by line and counts the number of lines.

Day 49: Writing Files



* Resource:[ Go by Example - Writing Files](https://gobyexample.com/writing-files)
* Exercise: Write a program that writes a series of strings to a text file.

Day 50: Line Filters



* Resource:[ Go by Example - Line Filters](https://gobyexample.com/line-filters)
* Exercise: Write a simple line filter that converts all input text to uppercase.

Day 51: Regular Expressions (continued)



* Resource:[ Go by Example - Regular Expressions](https://gobyexample.com/regular-expressions)
* Exercise: Write a program that uses regular expressions to validate email addresses.

Day 52: Sorting



* Resource:[ Go by Example - Sorting](https://gobyexample.com/sorting)
* Exercise: Write a program that sorts a slice of integers using the sort package.

Day 53: Custom Sorting



* Resource:[ Go by Example - Custom Sorting](https://gobyexample.com/custom-sorting)
* Exercise: Implement a custom sorting function to sort a slice of strings by length.

Day 54: Panic and Recover in Web Applications



* Resource:[ Graceful Handling of Panics in Go Web Applications](https://www.calhoun.io/using-panic-and-recover-to-create-graceful-web-applications/)
* Exercise: Add a panic recovery middleware to your HTTP server.

Day 55: Templates



* Resource:[ Go by Example - Templates](https://gobyexample.com/templates)
* Exercise: Write a program that uses HTML templates to render a dynamic web page.

Day 56: Middleware and Web Applications



* Resource:[ Creating Golang Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
* Exercise: Implement a custom middleware for logging HTTP requests in your server.

Day 57: Advanced Concurrency Patterns: sync



* Resource:[ Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* Exercise: Implement a pipeline pattern to process a series of integers concurrently.

Day 58: Advanced Concurrency Patterns: context



* Resource:[ Go Concurrency Patterns: Context](https://blog.golang.org/context)
* Exercise: Modify the pipeline pattern from Day 57 to support cancellation using the context package.

Day 59: Advanced Concurrency Patterns: atomic



* Resource:[ Go by Example - Atomic Counters](https://gobyexample.com/atomic-counters)
* Exercise: Write a concurrent counter using the sync/atomic package.

Day 60: Mutexes and Sync



* Resource:[ Go by Example - Mutexes](https://gobyexample.com/mutexes)
* Exercise: Implement a safe concurrent counter using sync.Mutex.

Day 61: Reflection



* Resource:[ The Laws of Reflection](https://blog.golang.org/laws-of-reflection)
* Exercise: Write a program that inspects the methods of a given struct using reflection.

Day 62: Unsafe Code



* Resource:[ An Introduction to the unsafe package](https://go101.org/article/unsafe.html)
* Exercise: Write a program that uses the unsafe package to directly manipulate memory.

Day 63: C Interoperability: cgo



* Resource:[ C? Go? Cgo!](https://blog.golang.org/c-go-cgo)
* Exercise: Write a simple Go program that calls a C function using cgo.

Day 64: Profiling Go Programs



* Resource:[ Profiling Go Programs](https://blog.golang.org/pprof)
* Exercise: Profile a Go program to identify performance bottlenecks.

Day 65: Optimization and Performance Tips



* Resource:[ High-performance Go Tips](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)
* Exercise: Optimize a sample Go program based on performance tips and profiling results.

Day 66-100: Project-based Learning & Practice**



* Apply your knowledge and skills by building projects and practicing coding problems on platforms like[ LeetCode](https://leetcode.com/) and[ Codewars](https://www.codewars.com/)

Suggested projects:



* Day 66-75: Develop a simple web API using Golang
* Day 76-85: Build a concurrent web scraper
* Day 86-95: Create a chat server using WebSockets
* Day 96-100: Implement a distributed system, such as a key-value store

Resources



1. Go by Example ([https://gobyexample.com/](https://gobyexample.com/)): A hands-on introduction to Go using annotated example programs.
2. GoDoc ([https://pkg.go.dev/](https://pkg.go.dev/)): The official Go documentation for all standard library packages, including fmt, bufio, strings, and regexp.
3. The Go Programming Language ([https://www.gopl.io/](https://www.gopl.io/)): A comprehensive book on learning Go, written by Alan A. A. Donovan and Brian W. Kernighan. (Paid resource)
4. Golang Tutorials ([https://golangbot.com/](https://golangbot.com/)): A collection of tutorials covering various aspects of Go programming, including the fmt and strings packages.
5. Go 101 ([https://go101.org/](https://go101.org/)): A comprehensive guide to the Go programming language, covering topics such as formatting, parsing, and data structures.
6. Effective Go ([https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html)): The official Go guide that provides tips and best practices for writing effective Go code.
7. YourBasic.org Go Tutorials ([https://yourbasic.org/golang/](https://yourbasic.org/golang/)): A collection of Go tutorials and articles, including a cheat sheet on string formatting and references on data structures and algorithms.
8. Calhoun.io ([https://www.calhoun.io/](https://www.calhoun.io/)): A blog by Jon Calhoun featuring articles and tutorials on Go programming, covering topics such as string manipulation and web development.
