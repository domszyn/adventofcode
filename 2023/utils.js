export const parseInput = (input, parseFn) => {
    const lines = input.split('\n');
    return parseFn ? lines.map(parseFn) : lines;
}

export const isDigit = (char) => char && char >= '0' && char <= '9';

export const makeArray = (length, initValue) => new Array(length).fill(initValue);