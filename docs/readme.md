# learnings

Check your Go PATH by `echo $GOPATH`.

Every immediate child directory inside `$GOPATH/src` is start of a library/project.

`$GOPATH/src/todo`
`$GOPATH/src/strutils`
`$GOPATH/src/uuid`
`$GOPATH/src/facebook`

Let's call this the project's root.

Every file inside of a project even deeply nested file needs to be part of a package.

A package  is defned using a single word.

`package userservice'
`package usercontroller'
`package httputils'

All files that belong to the same package can access all the global variables of each other.

The files that are not part of the same package can only use what is exported out of all the files that belong to that package.

To import a package you start with the project's root and follow the directory structure where the package is defined.
