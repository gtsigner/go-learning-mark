package main

func main() {

    for i := 0; i < 10000; i++ {
        print(log("hello world"))
        go log("hello")
    }
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

func log(msg string) (n int) {
    print(msg + "\n")
    return 1
}
