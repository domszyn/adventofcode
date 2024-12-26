import { input } from "./input.js";

export const knotHash = (input, rounds, dense) => {
    let lengths;
    if (dense) {
        lengths = [
            ...input.split('').map(s => s.charCodeAt(0)),
            17, 31, 73, 47, 23
        ];
    } else {
        lengths = input.split(',').map(Number);
    }

    var numbers = new Array(256);
    for (let i = 0; i < numbers.length; i++) {
        numbers[i] = i;
    }
    var currentPosition = 0;
    var skipSize = 0;

    for (var round = 0; round < rounds; round++) {
        for (const length of lengths) {
            var elements = [];
            for (var i = 0; i < length; i++) {
                elements.push(numbers[(i + currentPosition) % numbers.length]);
            }
            elements.reverse();
            for (var i = 0; i < length; i++) {
                numbers[(i + currentPosition) % numbers.length] = elements[i];
            }
            currentPosition += length + skipSize++;
            currentPosition %= numbers.length;
        }
    }

    if (dense) {
        const denseHash = new Array(16);
        for (let i = 0; i < 256; i += 16) {
            const n = numbers.slice(i, i + 16);
            for (var j = 0; j < 16; j++) {
                denseHash[i / 16] ^= n[j];
            }
        }
        return denseHash.map(n => n.toString(16).padStart(2, '0')).join('');
    } else {
        return numbers;
    }
};

export const solve = () => {

    const [a, b] = knotHash(input, 1, false);
    console.log("Part 1", a * b);
    console.log("Part 2", knotHash(input, 64, true));
}