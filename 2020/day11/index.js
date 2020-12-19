const fs = require("fs");

const grid = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((s) => Array.from(s));

const changeSeats = (seats, takeWith = 0, vacateWith = 4, farSight = false) => {
  const nextSeats = [[]];
  let changes = 0;

  const checkSeatUp = (x, y) => {
    if (y === 0) return 0;
    switch (seats[y - 1][x]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatUp(x, y - 1) : 0;
    }
  };

  const checkSeatDown = (x, y) => {
    if (y === seats.length - 1) return 0;
    switch (seats[y + 1][x]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatDown(x, y + 1) : 0;
    }
  };

  const checkSeatLeft = (x, y) => {
    if (x === 0) return 0;
    switch (seats[y][x - 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatLeft(x - 1, y) : 0;
    }
  };

  const checkSeatRight = (x, y) => {
    if (x === seats[0].length - 1) return 0;
    switch (seats[y][x + 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatRight(x + 1, y) : 0;
    }
  };

  const checkSeatUpLeft = (x, y) => {
    if (y === 0 || x === 0) return 0;
    switch (seats[y - 1][x - 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatUpLeft(x - 1, y - 1) : 0;
    }
  };

  const checkSeatUpRight = (x, y) => {
    if (y === 0 || x === seats[0].length - 1) return 0;
    switch (seats[y - 1][x + 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatUpRight(x + 1, y - 1) : 0;
    }
  };

  const checkSeatDownLeft = (x, y) => {
    if (y === seats.length - 1 || x === 0) return 0;
    switch (seats[y + 1][x - 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatDownLeft(x - 1, y + 1) : 0;
    }
  };

  const checkSeatDownRight = (x, y) => {
    if (y === seats.length - 1 || x === seats[0].length - 1) return 0;
    switch (seats[y + 1][x + 1]) {
      case "#":
        return 1;
      case "L":
        return 0;
      case ".":
      default:
        return farSight ? checkSeatDownRight(x + 1, y + 1) : 0;
    }
  };

  const getAdjacentOccupied = (x, y) =>
    checkSeatUp(x, y) +
    checkSeatUpRight(x, y) +
    checkSeatRight(x, y) +
    checkSeatDownRight(x, y) +
    checkSeatDown(x, y) +
    checkSeatDownLeft(x, y) +
    checkSeatLeft(x, y) +
    checkSeatUpLeft(x, y);

  const takeEmptySeat = (x, y) => {
    if (getAdjacentOccupied(x, y) > takeWith) return;
    nextSeats[y][x] = "#";
    changes++;
  };

  const vacateOccupiedSeat = (x, y) => {
    if (getAdjacentOccupied(x, y) < vacateWith) return;
    nextSeats[y][x] = "L";
    changes++;
  };

  for (let y = 0; y < seats.length; y++) {
    nextSeats[y] = [];
    for (let x = 0; x < seats[0].length; x++) {
      nextSeats[y][x] = seats[y][x];
      switch (seats[y][x]) {
        case "L":
          takeEmptySeat(x, y);
          break;
        case "#":
          vacateOccupiedSeat(x, y);
          break;
        case ".":
        default:
          break;
      }
    }
  }

  return [nextSeats, changes > 0];
};

let changed = true;
let nextGrid = grid;
do {
  [nextGrid, changed] = changeSeats(nextGrid);
} while (changed);

console.log(
  "Part1: ",
  nextGrid
    .map((row) => row.filter((seat) => seat === "#").length)
    .reduce((a, b) => a + b, 0)
);

changed = true;
nextGrid = grid;
do {
  [nextGrid, changed] = changeSeats(nextGrid, 0, 5, true);
} while (changed);

console.log(
  "Part2: ",
  nextGrid
    .map((row) => row.filter((seat) => seat === "#").length)
    .reduce((a, b) => a + b, 0)
);
