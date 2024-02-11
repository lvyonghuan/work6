//递归
const num=prompt("输入大于0的整数n")

function factorial(n,sum=n){
    if (n<0){
        alert("非法输入")
        return
    }else if(n==0){
        alert("阶乘为1")
        return
    }

    if (--n==0){
        alert(`阶乘为 ${sum}`)
        return
    }

    sum=(n)*sum
    factorial(n,sum)
}

factorial(num)