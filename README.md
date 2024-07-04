# sum-calculator

This program calculates using parallelism sum of integers in array of objects in JSON file

### Usage

1. You can create JSON file that will be in such format

```json
[
    {
        "a": 1,
        "b": 3
    },
    {
        "a": 5,
        "b": -9
    },
    {
        "a": -2,
        "b": 4
    }
    // ...
]
```
  Or you can just run program in generate directory to create data.json file in root directory:
  ```bash
  go run generate/main.go
  ```
2. Then run main.go file in root directory and pass two arguments(name of the file, number of goroutines):
### Example
  ```bash
  go run main.go data.json 3
  ```
