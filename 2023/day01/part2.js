import '../array.js';
import { isDigit, makeArray, parseInput } from "../utils.js";
import { input } from "./input.js";

var strings = parseInput(input);
var numbers = makeArray(strings.length, 0);

function wordToNumber(word) {
    switch (word) {
        case "one":
            return 1;
        case "two":
            return 2;
        case "three":
            return 3;
        case "four":
            return 4;
        case "five":
            return 5;
        case "six":
            return 6;
        case "seven":
            return 7;
        case "eight":
            return 8;
        case "nine":
            return 9;
        default:
            return 0;
    }
}

const spelledDigits = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

function findFirstSpelledDigit(message) {
    var digits = spelledDigits.map(x => ({ x: x, index: message.indexOf(x) }));
    var minIndex = message.length;
    var firstDigit = "";
    for (let j = 0; j < digits.length; j++) {
        if (digits[j].index < 0) {
            continue;
        }
        if (digits[j].index < minIndex) {
            minIndex = digits[j].index;
            firstDigit = digits[j].x;
        }
    }

    return [minIndex, firstDigit];
}

function findLastSpelledDigit(message) {
    var digits = spelledDigits.map(x => ({ x: x, index: message.lastIndexOf(x) }));
    var maxIndex = -1;
    var lastDigit = "";
    for (let j = 0; j < digits.length; j++) {
        if (digits[j].index < 0) {
            continue;
        }
        if (digits[j].index > maxIndex) {
            maxIndex = digits[j].index;
            lastDigit = digits[j].x;
        }
    }

    return [maxIndex, lastDigit];
}

for (let i = 0; i < strings.length; i++) {
    var message = strings[i];
    var digits = ["one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine"].map(x => ({ x: x, index: strings[i].indexOf(x) }));
    var [minIndex, firstDigit] = findFirstSpelledDigit(message);

    var found = false;
    for (let j = 0; j < minIndex; j++) {
        const element = message[j];
        if (isDigit(element)) {
            numbers[i] += parseInt(element) * 10;
            found = true;
            break
        }
    }
    if (!found) {
        numbers[i] += 10 * wordToNumber(firstDigit)
    }

    var [maxIndex, lastDigit] = findLastSpelledDigit(message);

    found = false
    for (let j = message.length - 1; j > maxIndex; j--) {
        const element = message[j];
        if (isDigit(element)) {
            numbers[i] += parseInt(element);
            found = true;
            break
        }
    }
    if (!found) {
        numbers[i] += wordToNumber(lastDigit);
    }

}

console.log(numbers.sum());