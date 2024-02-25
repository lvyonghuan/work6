setTimeout(() => {
    console.log('setTimeout start');
    new Promise((resolve) => {
    console.log('promise1 start');
    resolve();
    }).then(() => {
    console.log('promise1 end');
    })
    console.log('setTimeout end');
    }, 0);
    function promise2() {
    return new Promise((resolve) => {
    console.log('promise2');
    resolve();
    })
    }
    async function async1() {
    console.log('async1 start');
    await promise2();
    console.log('async1 end');
    }
    async1();
    console.log('script end');

    //1.async1 start 开始先执行全局
    //2.promise2 接着1打印
    //3.script end 2执行完之后等待状态改变。执行微任务。
    //4.async1 end 微任务队列的下一个
    //5.setTimeout start 微任务队列空，下一个宏任务
    //6.promise1 start 顺序接5
    //7. setTimeout end promise1状态变换，先打印
    //8.promise1 end resolve状态打印