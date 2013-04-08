gonavmap
========

A simple library to help you traverse maps of maps

Example
-------

// This will retrieve the map[string]interface{} under testing -> further
f := Get(testData, "testing.further")

// This will retrieve the value of testing -> further -> Str3
v := Value(testData, "testing.further.Str3")

// This will set the value of testing -> further -> Lolo
Set(testData, "testing.further.Lolo", 5)

The library assumes that you're always providing a map[string]interface{} for traversing.
