class Person {
    constructor(name, age) {
      this.name = name
      this.age = age
    }
  
    speak() {
      console.log(`姓名：${this.name}，年龄：${this.age}`)
    }
  }
  
  class Student extends Person {
    constructor(name, age, scores) {
      super(name, age)
      this.scores = scores
    }
  
    speak() {
      super.speak()
      console.log(`成绩：${this.scores}`)
    }
  }
  
  const person = new Person('张三', 18)
  person.speak()
  
  const student = new Student('李四', 19, [90, 80, 70])
  student.speak()
  