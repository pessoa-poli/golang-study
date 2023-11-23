# Go Concurrency Patterns

https://go.dev/blog/pipelines

https://go.dev/talks/2012/concurrency.slide#1

https://vimeo.com/49718712

Concurrency is a Model For Software Construction - Rob Pike  
Concurrency is the composition of independently executing computations. 
Concurrency is about dealing with a lot of things at once, and parallelism is about doing a lot of things at once.
Concurrency is a way of structuring things so that mayber with parallelism you could do a better job, but parallelism is not the goal of concurrency.
Concurrency's goal is a good structure.
### Go enables concurrent designs by haveing 4 essential:
features: Goroutines, Channels, the Select structure and Closures.  
A goroutine is a lightweight thread of execution.  
	
# Read Communicating Sequential Processes
PDF Included in this repository.


# What is a Closure?
https://stackoverflow.com/questions/36636/what-is-a-closure

## A closure is a persistent local variable scope

A closure is a persistent scope which holds on to local variables even after the code execution has moved out of that block. Languages which support closure (such as JavaScript, Swift, and Ruby) will allow you to keep a reference to a scope (including its parent scopes), even after the block in which those variables were declared has finished executing, provided you keep a reference to that block or function somewhere.  

The scope object and all its local variables are tied to the function and will persist as long as that function persists.  

This gives us function portability. We can expect any variables that were in scope when the function was first defined to still be in scope when we later call the function, even if we call the function in a completely different context.  
# Declaring Variadic Functions
https://gobyexample.com/variadic-functions

# Algorithms And Data Structures
https://algs4.cs.princeton.edu/cheatsheet/

# WHAT TO DO
- Cracking the Coding Interview
- Design Patterns
- Coursera Algorithms and Data Structure
- Clean Code
- Clean Architecture
- Refactoring

# Good Sources of Go Knowledge
https://gobyexample.com/switch
https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1046s
https://go.dev/doc/effective_go


# Named Return Values:
https://stackoverflow.com/questions/45239409/empty-return-in-func-with-return-value-in-golang

# Strings, Runes And Characters
## The paradigm in Go
A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.  
> https://go.dev/blog/strings

# About interfaces:
https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

# About generics:
https://gobyexample.com/generics

# Go's Proverb
 Go's proverb is: "Do not communicate by sharing memory; instead, share memory by communicating."

 https://stackoverflow.com/questions/52863273/how-to-make-a-variable-thread-safe

# Closures Running as GoRoutines
https://go.dev/doc/faq#closures_and_goroutines

# Internet and Networking
https://www.youtube.com/watch?v=S7MNX_UD7vY&list=PLIhvC56v63IJVXv0GJcl9vO5Z6znCVb1P
