function User(QQ, Password, Name, Birthday, NickName, Habit) {
    this.QQ = QQ;
    this.Password = Password;
    this.Name = Name;
    this.Birthday = Birthday;
    this.NickName = NickName;
    this.Habit = Habit;
    this.Friends = [];
    this.State = "";

    // 方法：登录QQ
    this.login = function (enteredPassword) {
        if (enteredPassword === this.Password) {
            console.log("Login successful!");
        } else {
            console.log("Incorrect password. Please try again.");
        }
    };

    // 方法：更改密码
    this.changePassword = function (newPassword) {
        this.Password = newPassword;
        console.log("Password changed successfully!");
    };

    // 方法：更改昵称
    this.changeNickname = function (newNickname) {
        this.NickName = newNickname;
        console.log("Nickname changed successfully!");
    };

    // 方法：添加好友
    this.addFriend = function (friendQQ) {
        this.Friends.push(friendQQ);
        console.log("Friend added successfully!");
    };

    // 方法：删除好友
    this.removeFriend = function (friendQQ) {
        const index = this.Friends.indexOf(friendQQ);
        if (index !== -1) {
            this.Friends.splice(index, 1);
            console.log("Friend removed successfully!");
        } else {
            console.log("Friend not found in your list.");
        }
    };

    // 方法：设置生日
    this.setBirthday = function (newBirthday) {
        this.Birthday = newBirthday;
        console.log("Birthday set successfully!");
    };

    // 方法：设置兴趣
    this.setHabit = function (newHabit) {
        this.Habit = newHabit;
        console.log("Habit set successfully!");
    };

    // 方法：查找好友信息
    this.findFriendInfo = function (friendQQ) {
        if (this.Friends.includes(friendQQ)) {
            console.log(`Friend's information - QQ: ${friendQQ}, Nickname: ${this.NickName}`);
        } else {
            console.log("Friend not found in your list.");
        }
    };

    // 方法：设置状态
    this.setState = function (newState) {
        this.State = newState;
        console.log("State set successfully!");
    };
}


// let user1 = new User(114514, "李田所", {Year: "1919", Month: "8", Day: "10"}, "李天梭", "健身");