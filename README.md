# mapstructure [![Godoc](https://godoc.org/github.com/eluv-io/mapstructure?status.svg)](https://godoc.org/github.com/eluv-io/mapstructure)

mapstructure is a Go library for decoding generic map values to structures
and vice versa, while providing helpful error handling.

This library is most useful when decoding values from some data stream (JSON,
Gob, etc.) where you don't _quite_ know the structure of the underlying data
until you read a part of it. You can therefore read a `map[string]interface{}`
and use this library to decode it into the proper underlying native Go
structure.

## Fork 

The original repository of this library has been archived as [planned by its author](https://gist.github.com/mitchellh/90029601268e59a29e64e55bab1c5bdc)
Therefore, the library was repackaged in this repository to include few fixes.

## Installation

Standard `go get`:

```
$ go get github.com/eluv-io/mapstructure
```

## Usage & Example

For usage and examples see the [Godoc](http://godoc.org/github.com/eluv-io/mapstructure).

The `Decode` function has examples associated with it there.

## But Why?!

Go offers fantastic standard libraries for decoding formats such as JSON.
The standard method is to have a struct pre-created, and populate that struct
from the bytes of the encoded format. This is great, but the problem is if
you have configuration or an encoding that changes slightly depending on
specific fields. For example, consider this JSON:

```json
{
  "type": "person",
  "name": "Mitchell"
}
```

Perhaps we can't populate a specific structure without first reading
the "type" field from the JSON. We could always do two passes over the
decoding of the JSON (reading the "type" first, and the rest later).
However, it is much simpler to just decode this into a `map[string]interface{}`
structure, read the "type" key, then use something like this library
to decode it into the proper structure.
