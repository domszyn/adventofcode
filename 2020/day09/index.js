const fs = require("fs");

const numbers = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((_) => +_);

const validateNumberAt = (idx) => {
  if (idx < 25) throw new Error();
  const preamble = numbers.slice(idx - 25, idx);

  for (let i = 0; i < 25; i++) {
    for (let j = i + 1; j < 25; j++) {
      if (preamble[i] + preamble[j] === numbers[idx]) {
        return true;
      }
    }
  }

  return false;
};

const firstInvalidNumber = () => {
  for (let i = 25; i < numbers.length; i++) {
    if (!validateNumberAt(i)) return [numbers[i], i];
  }

  return [];
};

const addNumbers = (startIdx, endIdx) =>
  numbers.slice(startIdx, endIdx).reduce((a, b) => a + b, 0);

const findWeakness = ([number, idx]) => {
  for (let sliceLength = 2; sliceLength < idx + 1; sliceLength++) {
    for (let startIdx = 0; startIdx < idx - sliceLength + 1; startIdx++) {
      const sum = addNumbers(startIdx, startIdx + sliceLength);
      if (sum === number) {
        const range = numbers.slice(startIdx, startIdx + sliceLength);
        return Math.min(...range) + Math.max(...range);
      }
    }
  }

  return [];
};

console.log("Part1: ", firstInvalidNumber()[0])

console.log("Part2: ", findWeakness(firstInvalidNumber()));
