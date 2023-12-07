import '../array.js';
import { isDigit, makeArray, parseInput } from '../utils.js';
import { input } from './input.js';

var strings = parseInput(input);
var numbers = makeArray(strings.length, 0);

for (let i = 0; i < strings.length; i++) {
    var message = strings[i];
    for (let j = 0; j < message.length; j++) {
        const element = message[j];
        if (isDigit(element)) {
            numbers[i] += parseInt(element) * 10;
            break
        }
    }

    for (let j = message.length - 1; j >= 0; j--) {
        const element = message[j];
        if (isDigit(element)) {
            numbers[i] += parseInt(element);
            break
        }
    }
}

console.log(numbers.sum());