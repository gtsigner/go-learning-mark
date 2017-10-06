package main;

/*
* var 定义变量，支持类型推断，有助于跨平台应用开发。 编译器确保变量总是被初始化为零值。
*/
func main() {
    var x int32;
    println(x, 0);
    x = 100;
    var s = "hello world";
    println(x, s);
    println();
    y:=100
    print(y)
    variable2();
}

/*
* 函数内部，省略var关键词，更简单的定义模式
*/
func variable2() {
    x := 100;
    println("X:=100 print:", x);
}
