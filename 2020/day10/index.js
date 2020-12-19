const fs = require("fs");

let adapters = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((_) => +_)
  .sort((a, b) => a - b);

adapters = [0, ...adapters, adapters[adapters.length - 1] + 3];

const diff = [, 0, , 0];

for (let i = 1; i < adapters.length; i++) {
  diff[adapters[i] - adapters[i - 1]] += 1;
}

console.log("Part1: ", diff[1] * diff[3]);

const ways = [1];
ways.length = adapters.length;
ways.fill(0, 1);

for (let i = 0; i < ways.length; i++) {
  for (let j = i - 3; j < i; j++) {
    if (adapters[i] <= adapters[j] + 3) {
      ways[i] += ways[j];
    }
  }
}

console.log("Part2: ", ways.reverse()[0]);
