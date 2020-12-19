const fs = require("fs");

const numbers = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((s) => +s);

const multiply = (a, b) => a * b;

let answers = [];
for (let i = 0; i < numbers.length; i++) {
  for (let j = i + 1; j < numbers.length; j++) {
    const a = numbers[i],
      b = numbers[j];
    if (a + b === 2020) {
      answers[0] = [a, b];
      if (answers[1]) break;
    }

    for (let k = j + 1; k < numbers.length; k++) {
      const c = numbers[k];
      if (a + b + c === 2020) answers[1] = [a, b, c];
    }
  }
}

console.log("Part1: ", answers[0].reduce(multiply, 1));
console.log("Part2: ", answers[1].reduce(multiply, 1));
