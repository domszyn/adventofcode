const fs = require("fs");

const answers = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n\n");

const answerGroups = answers.map((_) => new Set(_.replaceAll("\n", "")));

const yesAnswers = answerGroups
  .map((_) => _.size)
  .reduce((a, b) => a + b, 0);

const separateAnswers = answers
  .map((_) => _.split("\n").map((s) => Array.from(s)))
  .map((arr) => arr.reduce((a, b) => a.filter((_) => b.includes(_)), arr[0]));

const commonYesAnswers = separateAnswers
  .map((arr) => arr.length)
  .reduce((a, b) => a + b, 0);

console.log("Part1: ", yesAnswers);
console.log("Part2: ", commonYesAnswers);
