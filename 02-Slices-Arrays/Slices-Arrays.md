## Arrays & Slices (01-Arrays.go, 02-Slices.go)

1. Array - **Using `[...]int{1, 2, 3, 4, 5}`:**
    ```go
    var numbers = [...]int{1, 2, 3, 4, 5}
    ```
    - This line declares and initializes an array with a fixed size. The size is determined by the number of elements provided in the array literal.
    - In this case, the array will have the type `[5]int`, and its size is fixed at compile-time.
    - The `...` syntax indicates that the size is automatically determined based on the number of elements.

2. Slice - **Using `[]int{1, 2, 3, 4, 5}`:**
    ```go
    var numbers = []int{1, 2, 3, 4, 5}
    ```
    - This line declares and initializes a slice. Unlike an array, a slice is a dynamically-sized, flexible view into an underlying array.
    - Slices are more commonly used in Go for their flexibility and ease of use compared to fixed-size arrays.
    - The type of `numbers` is `[]int`, representing a slice of integers.

**Key Differences:**
- The first approach with `[...]int{1, 2, 3, 4, 5}` initializes an array with a fixed size (in this case, `[5]int`).
- The second approach with `[]int{1, 2, 3, 4, 5}` initializes a slice, which is more flexible and does not have a fixed size.

In many cases, slices are preferred in Go for their dynamic nature, and arrays are used less frequently due to their fixed size and more rigid nature. Slices allow for more convenient manipulation and passing of subsets of data.