import '../array.js';
import { input } from './input.js';
import { parseInput } from '../utils.js';

const equations = parseInput(input, l => {
    const [left, rest] = l.split(': ');

    return {
        left: +left,
        numbers: rest.split(' ').map(n => +n)
    }
});

const solve = ({ left, numbers }, base) => {
    for (let operators = 0; operators < Math.pow(base, numbers.length - 1); operators++) {
        const operatorString = operators.toString(base).padStart(numbers.length - 1, '0');

        if (left === reduce(operatorString, numbers)) {
            return true;
        }
    }

    return false;
}

const reduce = (mask, numbers) => {
    let result = numbers[0];
    for (let i = 0; i < mask.length; i++) {
        switch (mask.charAt(i)) {
            case '0':
                result += numbers[i + 1];
                break;
            case '1':
                result *= numbers[i + 1];
                break;
            case '2':
                result = +(String(result) + String(numbers[i + 1]));
                break;
        }
    }
    return result;
}

console.log('Part 1', equations.filter(e => solve(e, 2)).map(({ left }) => left).sum());
console.log('Part 2', equations.filter(e => solve(e, 3)).map(({ left }) => left).sum())