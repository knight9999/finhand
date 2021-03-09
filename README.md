# Final Handler

This module allows you to register a handler and execute it at a specific time, similar to the `atexit` in `c++` language.

# Methods

## Add new Handler `f`.

```
func AddHandler(f func()) string 
```

Register the new handler `f`.
The return value is a name of handler.
In this case the name of the handler is a uuid string.
The priority of this handler is zero.

## Add new Handler `f` with Priority.

```
func AddHandlerWithPriority(priority int, f func()) string
```

Register the new handler `f`.
The return value is a name of handler.
In this case the name of the handler is a uuid string.
The handler with the higher priority number will be given priority.
If they have the same priority value, the first handler to register will have higher priority.


## Add new Handler `f` with Name and Priority.

```
func AddHandlerWithNameAndPriority(name string, priority int, f func()) {
```

Register the new handler `f`.
This can specify the name of the handler and priority.
There is no return value.
This overrides the handler with the same name if it exists.

## Remove Handler by Name

```
func RemoveHandler(name string)
```

## Run all Handlers

```
func RunHandlers()
```

This clears list of the registered handler.

# Testing

```
$ go test
```

