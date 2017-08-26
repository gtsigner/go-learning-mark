package main;



func main(){
    x:=100;

    // if else
    if x>0 {
        println("x>0");
    }else if x<0 {
        println("x<0");
    }else{
        println("x=0");
    }


    // switch
    x=-100;
    switch{
        case x>0:
            println("x>0");
        case x<0:
            println("x<0");
        default:
            println("x=0");
    }


    //for
    for i:=0;i<5;i++ {
        println(i);
    }

    //for 相当于while 一样
    for x<0 {
        x++;
        println(x);
    }


    //迭代
    y:=[]int{100,101,102};
    for i,n:=range y{
        println(i,":",n);
    }
}
