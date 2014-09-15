AppEngine DeMilitarized Zone
============================

Package aedmz is an AppEngine abstraction layer to run a service either locally
or on AppEngine.

This package contains code and interfaces to make it possible for a server to
run on either AppEngine or as a local process. This permits not being locked in
on AppEngine for a service coded using exclusively this package instead of using
the "appengine" package directly.

See the [![GoDoc](https://godoc.org/github.com/maruel/aedmz?status.svg)](https://godoc.org/github.com/maruel/aedmz).


Testing package
---------------

Package aedmz/aedmztest is an AppEngine abstraction layer for unit testing.

Coupled with the package aedmz, it enables unit testing that properly simulate
either a running (local or AppEngine) environment.

See the testing specific [![GoDoc](https://godoc.org/github.com/maruel/aedmz/aedmztest?status.svg)](https://godoc.org/github.com/maruel/aedmz/aedmztest).
