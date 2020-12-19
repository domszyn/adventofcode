const fs = require("fs");

const terrain = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((s) => Array.from(s));

const traverse = (right, down) => {
  let width = terrain[0].length;
  let collisionCount = 0;
  let x = 0;
  for (let y = down; y < terrain.length - down; y += down) {
    x = (x + right) % width;
    switch (terrain[y][x]) {
      case ".":
        terrain[y][x] = "O";
        break;
      case "#":
        terrain[y][x] = "X";
      case "X":
        collisionCount++;
        break;
    }
  }

  return collisionCount;
};

console.log("Part1: ", traverse(3, 1));

const multiply = (a, b) => a * b;

console.log(
  "Part2: ",
  [
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2],
  ]
    .map(([right, down]) => traverse(right, down))
    .reduce(multiply, 1)
);
