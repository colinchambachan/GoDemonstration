func fib(n int) int {
    // "base case" where if n <=1 then fib(n) = n
    if (n <= 1){
        return n
    }

    // create two variables to track going onward 
    f1, f2 := 0, 1


    // iteratively add previous terms until n-th term of fibonacci sequence
    for i:= 2; i <= n; i++ {

        f1, f2 = f2, f1 + f2
    } 

    // return f2 which is the n-th term of the fib sequence
    return f2


}