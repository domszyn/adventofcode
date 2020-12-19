const { readFileSync } = require("fs");

const cubes = [
  readFileSync(`${__dirname}/input.txt`, "utf-8")
    .split("\n")
    .map((s) => Array.from(s)),
];

const changeCubes = (cubes) => {
  const nextCubes = [[[]]];

  const getActiveNeighbors = (x, y, z) => {
    let active = 0;

    for (let dz = z - 1; dz <= z + 1; dz++) {
      for (let dy = y - 1; dy <= y + 1; dy++) {
        for (let dx = x - 1; dx <= x + 1; dx++) {
          if (dy === y && dx === x && dz === z) continue;
          if (getState(dx, dy, dz) === "#") active++;
        }
      }
    }

    return active;
  };

  const getState = (x, y, z) => {
    if (
      cubes[z] === undefined ||
      cubes[z][y] === undefined ||
      cubes[z][y][x] === undefined
    ) {
      return ".";
    }
    return cubes[z][y][x];
  };

  const changeState = (x, y, z) => {
    const activeNeighbors = getActiveNeighbors(x, y, z);
    const currentState = getState(x, y, z);
    if (currentState === "#") {
      return activeNeighbors === 2 || activeNeighbors === 3 ? "#" : ".";
    }
    if (currentState === ".") {
      return activeNeighbors === 3 ? "#" : ".";
    }
  };

  for (let z = 0; z < cubes.length + 2; z++) {
    nextCubes[z] = [];
    for (let y = 0; y < cubes[0].length + 2; y++) {
      nextCubes[z][y] = [];
      for (let x = 0; x < cubes[0][0].length + 2; x++) {
        nextCubes[z][y][x] = getState(x - 1, y - 1, z - 1);
        nextCubes[z][y][x] = changeState(x - 1, y - 1, z - 1);
      }
    }
  }

  return nextCubes;
};

let nextCubes = cubes;
for (let i = 1; i < 7; i++) {
  nextCubes = changeCubes(nextCubes);
}

let active3D = 0;
for (let z = 0; z < nextCubes.length; z++) {
  active3D += nextCubes[z]
    .map((row) => row.filter((seat) => seat === "#").length)
    .reduce((a, b) => a + b, 0);
}

module.exports = { active3D };
