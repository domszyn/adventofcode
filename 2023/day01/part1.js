var input = ``;

var strings = input.split('\n');
var numbers = Array(strings.length).fill(0);

for (let i = 0; i < strings.length; i++) {
    var message = strings[i];
    for (let j = 0; j < message.length; j++) {
        const element = message[j];
        if (element >= '0' && element <= '9') {
            numbers[i] += parseInt(element) * 10;
            break
        }
    }

    for (let j = message.length - 1; j >= 0; j--) {
        const element = message[j];
        if (element >= '0' && element <= '9') {
            numbers[i] += parseInt(element);
            break
        }
    }
}

var sum = numbers.reduce((a, b) => a + b, 0);
console.log(sum);