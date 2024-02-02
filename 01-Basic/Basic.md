# Basic Types

```Go
bool

string

int  int8   int16   int32   int64
uint uint8  uint16  uint32  uint64

byte // alias for uint8

rune // alias for int32
     // represent a Unicode code point

float32 float64

complex64 complex128
```

- `string` : Read Only slice of bytes.
- `uint` : Unsigned integer


## Numeric types

### Signed and Unsigned integers

Go has several built-in integer types of varying sizes for storing signed and unsigned integers

The size of the generic int and uint types are platform-dependent. This means it is 32-bits wide on a 32-bit system and 64-bits wide on a 64-bit system.

```Go
var i int = 404                     // Platform dependent
var i8 int8 = 127                   // -128 to 127
var i16 int16 = 32767               // -2^15 to 2^15 - 1
var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
```
Similar to signed integers, we have unsigned integers.

```Go
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64
var uiptr uintptr                     // Integer representation of a memory address
```

If you noticed, there's also an unsigned integer pointer uintptr type, which is an integer representation of a memory address. It is not recommended to use this, so we don't have to worry about it.

### So which one should we use?

It is recommended that whenever we need an integer value, we should just use int unless we have a specific reason to use a sized or unsigned integer type.

### Byte and Rune

Golang has two additional integer types called byte and rune that are aliases for uint8 and int32 data types respectively.

```Go
type byte = uint8
type rune = int32
```
A rune represents a unicode code point.

```Go
var b byte = 'a'
var r rune = '🍕'
```

### Floating point

Next, we have floating point types which are used to store numbers with a decimal component.

Go has two floating point types float32 and float64. Both type follows the IEEE-754 standard.

The default type for floating point values is float64.

```Go
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit
```

## Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

## Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:

```Go 
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
```
Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

## Short Variable Declaration

> Note: Shorthand only works inside function bodies.

```Go
var empty string

// is the same as
empty := ""

numCars := 10 // Inferred to be an integer

var isFunny = true

```

## Same Line Declarations

```Go
mileage, company := 1029, "Toyota"

// is the same as

mileage := 1029
company := "Toyota"
```

# Constant

```Go
const a = 10
const b = a // ✅ Works

var a = 10
const b = a // ❌ a (variable of type int) is not constant (InvalidConstInit)
```

# String

```Go
var name string = "My name is Go"

var bio string = `I am statically typed.
									I was designed at Google.`
```