const fs = require("fs");

const instructions = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n");

const applyMask = (value, mask, skipChar) => {
  const bits = Array.from(value.toString(2).padStart(mask.length, "0"));
  for (let i = 0; i < mask.length; i++) {
    if (mask[i] === skipChar) continue;
    bits[i] = mask[i];
  }

  return bits.join("");
};

let mem = [0];
let mask;
instructions.forEach((line) => {
  if (line.startsWith("mask")) {
    mask = Array.from(line.substring(line.indexOf("=") + 2));
    return;
  }

  let [, address] = line.match(/\[(\d+)\]/);
  let [value] = line.match(/\d+$/);
  mem[+address] = parseInt(applyMask(+value, mask, "X"), 2);
});

console.log(
  "Part1: ",
  mem.filter(Boolean).reduce((a, b) => a + b, 0)
);

const replaceCharAt = (str, idx, char) =>
  str.substring(0, idx) + char + str.substring(idx + 1);

const getReplacements = (len) => {
  let replacements = [];
  for (let j = 0; j < 1 << len; j++) {
    replacements.push(j.toString(2).padStart(len, "0"));
  }
  return replacements;
};

const replaceFloatingPoints = (str) => {
  const xAt = [...str]
    .map((c, i) => (c === "X" ? i : undefined))
    .filter((c) => c !== undefined);
  let replacements = getReplacements(xAt.length);
  let replaced = [];

  for (let a = 0; a < replacements.length; a++) {
    let r = str;
    for (let b = 0; b < xAt.length; b++) {
      r = replaceCharAt(r, xAt[b], replacements[a].charAt(b));
    }
    replaced.push(r);
  }

  return replaced;
};

mem = new Map();
instructions.forEach((line) => {
  if (line.startsWith("mask")) {
    mask = Array.from(line.substring(line.indexOf("=") + 2));
    return;
  }
  let [, address] = line.match(/\[(\d+)\]/);
  let [value] = line.match(/\d+$/);
  const writeTo = replaceFloatingPoints(applyMask(+address, mask, "0"));
  for (let ra = 0; ra < writeTo.length; ra++) {
    mem.set(parseInt(writeTo[ra], 2), +value);
  }
});

console.log(
  "Part2: ",
  [...mem.values()].reduce((a, b) => a + b, 0)
);
