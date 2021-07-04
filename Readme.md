# go-journey

> Within a package, code can refer to any identifier (name) defined within, while clients of the package may only reference the package's exported types, functions, constants, and variables.

> A package should always have a directory. If you have a file with package name as 'xyz'. That file should live inside a directory xyz. That could be a direct parent or an ancestor directory.

> p.d is a shorthand for (*p).d and if d is also a pointer p.d.c is a shorthand for (*p)(.d).

> You can not access properties on uninitialized properties but can invoke methods on non null stuff https://www.reddit.com/r/golang/comments/5eizbo/why_doesnt_go_have_a_nullsafe_operator/

Zero values for int is 0, string is ""
Zero value for func(), map is nil, slice is nil, channel is nil, interface is nil