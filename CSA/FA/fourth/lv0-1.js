const funcs = []
for (let i = 0; i < 10; i++) {
  funcs.push(() => {
    console.log(i)
  })
}

funcs.forEach(func => func())
