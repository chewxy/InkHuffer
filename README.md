# package `inkhuffer`

package `inkhuffer` is a package that provides debugging-by-print-statement harnesses. It's essentially  `printf` statements on steroids. Annotate your code with print statements to trace execution. The extra features this package provides over normal print statements is scope tracking (manually managed of course). 

Think of this as a debugger where your step-throughs are logged to a file.

# Rational #
The intent is that this package should only be used during initial development and debug phases - this is useful for tracing execution of particular code paths without using a debugger. 

This package is NOT a logging solution. It's NOT a debugger solution either. 

The idea behind this is rather simple: There are "scopes" of execution in the flow of normal code. Take for example this snippet:

```go
func A() {
	B()
}

func B() {
	C()
	D()
}

func C() {
	// do work
}

func D() {
	// do work
}

```

The flow of execution can be visualized as such:

```
A()
   -> B()
         -> C()
         -> D()
```

Here we see `C()` and `D()` are pretty much on the same "scope" of execution scope.

Of course, scopes are not just limited to function calls. Scopes can be semantic in nature. This package offers only annotation capacity, and does not restrict on (ab)use.

# Poor Designs #

This package is intentionally poorly designed - the use of global variables is rife, and it's guaranteed to not play well with concurrency. 

The purpose as previously mentioned, is to aid in development and rapid understanding of the semantics of the program you're developing. You're not supposed to keep using this package.

# Usage #

If you import this package you  will note that most functions are NO-OPs. To use it properly you need to use the `tags=debug` build tag.