function sumAsciiCodes() {
    const inputString = prompt("输入一段字符串:");
    let asciiSum = 0;

    for (let i = 0; i < inputString.length; i++) {
        if (inputString[i] === ' ') {
            asciiSum += 0;
        } else {
            asciiSum += inputString.charCodeAt(i);
        }
    }

    alert(`字符串的ASCII码和为: ${asciiSum}`);
}

sumAsciiCodes();
