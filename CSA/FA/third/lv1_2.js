//循环
const num=prompt("输入大于0的整数n")

function factorial(n){
    if (n<0){
        alert("非法输入")
        return
    }else if(n==0){
        alert("阶乘为1")
        return
    }

    var sum=1
    for(var i=1;i<=n;i++){
        sum=i*sum
    }

    alert(`阶乘为 ${sum}`)
}

factorial(num)