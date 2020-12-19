const { readFileSync } = require("fs");

const cubes = [
  readFileSync(`${__dirname}/input.txt`, "utf-8")
    .split("\n")
    .map((s) => Array.from(s)),
];

const changeCubes = (cubes) => {
  const nextCubes = [[[[]]]];

  const getActiveNeighbors = (x, y, z, w) => {
    let active = 0;

    for (let dw = w - 1; dw <= w + 1; dw++) {
      for (let dz = z - 1; dz <= z + 1; dz++) {
        for (let dy = y - 1; dy <= y + 1; dy++) {
          for (let dx = x - 1; dx <= x + 1; dx++) {
            if (dx === x && dy === y && dz === z && dw === w) continue;
            if (getState(dx, dy, dz, dw) === "#") active++;
          }
        }
      }
    }

    return active;
  };

  const getState = (x, y, z, w) => {
    if (
      cubes[w] === undefined ||
      cubes[w][z] === undefined ||
      cubes[w][z][y] === undefined ||
      cubes[w][z][y][x] === undefined
    ) {
      return ".";
    }
    return cubes[w][z][y][x];
  };

  const changeState = (x, y, z, w) => {
    const activeNeighbors = getActiveNeighbors(x, y, z, w);
    const currentState = getState(x, y, z, w);
    if (currentState === "#") {
      return activeNeighbors === 2 || activeNeighbors === 3 ? "#" : ".";
    }
    if (currentState === ".") {
      return activeNeighbors === 3 ? "#" : ".";
    }
  };

  for (let w = 0; w < cubes.length + 2; w++) {
    nextCubes[w] = [];
    for (let z = 0; z < cubes[0].length + 2; z++) {
      nextCubes[w][z] = [];
      for (let y = 0; y < cubes[0][0].length + 2; y++) {
        nextCubes[w][z][y] = [];
        for (let x = 0; x < cubes[0][0][0].length + 2; x++) {
          nextCubes[w][z][y][x] = changeState(x - 1, y - 1, z - 1, w - 1);
        }
      }
    }
  }

  return nextCubes;
};

let nextCubes = cubes;
for (let i = 1; i < 7; i++) {
  nextCubes = changeCubes(nextCubes, i);
}

let active4D = 0;
for (let w = 0; w < nextCubes.length; w++) {
  let wCubes = nextCubes[w];
  for (let z = 0; z < wCubes.length; z++) {
    active4D += wCubes[z]
      .map((row) => row.filter((seat) => seat === "#").length)
      .reduce((a, b) => a + b, 0);
  }
}

module.exports = { active4D };
