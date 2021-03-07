# Final Handler

 This module enables to register handlers with name and priority.
 Each handler is executed by the `RunHandlers` command.

# Methods

## Add new Handler `f`.

```
func AddHandler(f func()) string 
```

The return value is a name of handler.
The name is a uuid string.
You can remove handlere by this name.
The priority of this handler is zero.

## Add new Handler `f` with Priority.

```
func AddHandlerWithPriority(priority int, f func()) string
```

The return value is a name of handler.
The name is a uuid string.
The handler with the higher priority number will be given priority.
If they have the same priority value, the first handler to register will have higher priority.


## Add new Handler `f` with Name and Priority.

```
func AddHandlerWithNameAndPriority(name string, priority int, f func()) {
```

This can specify the name of handler and priority.
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

# Testing

```
$ go test
```

