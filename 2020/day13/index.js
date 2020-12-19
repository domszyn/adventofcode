const fs = require("fs");

const [earliest, busLines] = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n");

const getLineNumbers = (includeBroken = false) => {
  let lineNumbers = busLines.split(",").map((s) => +s);

  if (includeBroken) return lineNumbers;

  return lineNumbers.filter((ln) => !isNaN(ln)).sort((a, b) => a - b);
};

const activeLines = getLineNumbers();
for (let departure = +earliest; ; departure++) {
  let [firstAvailable] = activeLines.filter((x) => departure % x === 0);
  if (firstAvailable) {
    console.log("Part1: ", firstAvailable * (departure - earliest));
    break;
  }
}

const allLines = getLineNumbers(true);
let t = 0;
let inc = allLines[0];
for (let i = 1; i < allLines.length; i++) {
  if (isNaN(allLines[i])) continue;

  for (let j = t; ; j += inc) {
    if ((j + i) % allLines[i] === 0) {
      t = j;
      break;
    }
  }
  inc *= allLines[i];
}
console.log("Part2: ", t);

