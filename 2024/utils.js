export const parseInput = (input, parseFn) => {
    const lines = input.split('\n');
    return parseFn ? lines.map(parseFn) : lines;
}

export const isDigit = (char) => char && char >= '0' && char <= '9';

export const makeArray = (length, initValue) => new Array(length).fill(initValue);

export const primeFactors = (n) => {
    let factors = [];

    while (n % 2 == 0) {
        factors.push(2);
        n /= 2;
    }

    for (let i = 3; i * i <= n; i = i + 2) {
        while (n % i == 0) {
            factors.push(i);
            n /= i;
        }
    }

    if (n > 2) {
        factors.push(n);
    }

    return factors;
}