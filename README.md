## GopherPy

![logo](./img/logo.png)

Gopher is a transpiler that transfers basic codes written in Python to basic codes written in Go. If you want to know more about how the transpiler works, read the <a href="./docs/documentation.md">documentation</a>.

### Usage

```go run main.go <python-file>```

### Example

* Input code:

```Python
def soma(a, b):
    return a + b

if __name__ == "__main__":
    print(soma(1, 2))
```

* Output code:

```Golang
func soma(a,() {a + b 
}
```