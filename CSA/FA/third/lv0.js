const person = {
    name: '佐科姐姐',
    age: 1000000,
    address: {
    city: 'ChongQing',
    area: 'NanShan'
    },
    title: ['student',{year:2022, title:'GoodStudent'}]
}

const {...rest}=person
const{name,age:year,address,title}=rest

console.log(name) // 佐科姐姐
console.log(year) // 1000000 这里没有写错哈，就是要输出 1000000，结合课件

const{city,area:mountain}=address

console.log(city) // ChongQing
console.log(mountain) // NanShan //这里没有写错，就是要输出 NanShan，结合课件

const[title1,titleA]=title

console.log(title1) // student

const{year:temp,title:title2}=titleA

console.log(title2) // GoodStudent