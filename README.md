# Mandelbrot Set Generator

This project generates an ASCII representation of the Mandelbrot Set using Python, Go, and Rust. It calculates the escape time for each point in the complex plane and visualizes the fractal in the terminal.

## Features

- Uses complex numbers to compute Mandelbrot iterations.
- Adjustable resolution and iteration limits.
- ASCII visualization with different symbols for iteration depths.
- Optimized memory allocation for performance.

## Configuration

You can adjust the parameters in the `main` file of each language's implementation. For example, in Rust:

```rust
let mandelbrot = calculate_mandelbrot(1000, -2.0, 1.0, -1.0, 1.0, 100, 24);
```

### Parameters:

- `1000`: Maximum iterations (higher values increase detail).
- `-2.0, 1.0, -1.0, 1.0`: Boundaries of the complex plane.
- `100, 24`: Width and height of the output.

## Code Usage

To run the Mandelbrot Set generator in different languages:

### Golang:

```sh
cd golang && go run main.go
```

### Python:

```sh
cd python && python3 main.py
```

### Rust:

```sh
cd rust && cargo run --release
```

## Example Output

```
                 .........................................••••*••**•................
              ...........................................•••••***••••..................
           ..........................................••••**+*+%%+***••....................
         .......................................••••••••••*%%%%%%%+*•••••••.................
       .....................................••***%*•••***+***%%%%********••••••*•.............
     ....................................••••••**%%$*%%%%%%%%%%%%%%%%%%%*+*+x***••..............
    ..................................•••••••****$%%%%%%%%%%%%%%%%%%%%%%%%%%%%*•••...............
   ....................•*•••••••••••••••••••*+%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*•••...............
  ...................••••****•••+*••••••••**%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%**•...............
 ....................•••••**%%*+#%%x%******+%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%**••................
 .................••••••*+$%%%%%%%%%%%%%+*x%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*••.................
 ........•••••••••••******%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*••..................
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%**••••..................
 ........•••••••••••******%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*••..................
 .................••••••*+$%%%%%%%%%%%%%+*x%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*••.................
 ....................•••••**%%*+#%%x%******+%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%**••................
  ...................••••****•••+*••••••••**%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%**•...............
   ....................•*•••••••••••••••••••*+%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%*•••...............
    ..................................•••••••****$%%%%%%%%%%%%%%%%%%%%%%%%%%%%*•••...............
     ....................................••••••**%%$*%%%%%%%%%%%%%%%%%%%*+*+x***••..............
       .....................................••***%*•••***+***%%%%********••••••*•.............
         .......................................••••••••••*%%%%%%%+*•••••••.................
           ..........................................••••**+*+%%+***••....................
              ...........................................•••••***••••..................
```

## License

This project is open-source and available under the MIT License.
