const { readFileSync } = require("fs");
const { env } = require("process");

const readTiles = (tileList) => {
  let tiles = [];
  let position = 0;

  do {
    if (tileList.startsWith("se", position)) {
      tiles.push("se");
      position += 2;
    } else if (tileList.startsWith("sw", position)) {
      tiles.push("sw");
      position += 2;
    } else if (tileList.startsWith("ne", position)) {
      tiles.push("ne");
      position += 2;
    } else if (tileList.startsWith("nw", position)) {
      tiles.push("nw");
      position += 2;
    } else if (tileList.startsWith("e", position)) {
      tiles.push("e");
      position++;
    } else if (tileList.startsWith("w", position)) {
      tiles.push("w");
      position++;
    }
  } while (position < tileList.length);

  return tiles;
};

const tileList = readFileSync(`${__dirname}/input.txt`, "utf-8")
  .split("\n")
  .map(readTiles);

const getTileKey = ({ row, column }) => `${row}#${column}`;
const getLocation = (tileKey) => {
  let [row, column] = tileKey.split("#");
  return { row: +row, column: +column };
};

const getAdjacentTiles = (floor, { row, column }) => {
  return [
    floor.get(getTileKey({ row, column: column + 1 })) || "white",
    floor.get(getTileKey({ row, column: column - 1 })) || "white",
    floor.get(getTileKey({ row: row + 1, column: column + 1 })) || "white",
    floor.get(getTileKey({ row: row + 1, column: column })) || "white",
    floor.get(getTileKey({ row: row - 1, column: column })) || "white",
    floor.get(getTileKey({ row: row - 1, column: column - 1 })) || "white",
  ];
};

let floor = new Map();
floor.set(getTileKey({ row: 0, column: 0 }), "white");

const flipTile = (location) => {
  const tileKey = getTileKey(location);
  floor.set(tileKey, floor.get(tileKey) === "white" ? "black" : "white");
};

for (let i = 0; i < tileList.length; i++) {
  let currentTile = { row: 0, column: 0 };
  for (let j = 0; j < tileList[i].length; j++) {
    let direction = tileList[i][j];
    let nextTile;
    switch (direction) {
      case "e":
        nextTile = { row: currentTile.row, column: currentTile.column + 1 };
        break;
      case "w":
        nextTile = { row: currentTile.row, column: currentTile.column - 1 };
        break;
      case "se":
        nextTile = {
          row: currentTile.row + 1,
          column: currentTile.column + 1,
        };
        break;
      case "sw":
        nextTile = { row: currentTile.row + 1, column: currentTile.column };
        break;
      case "ne":
        nextTile = { row: currentTile.row - 1, column: currentTile.column };
        break;
      case "nw":
        nextTile = {
          row: currentTile.row - 1,
          column: currentTile.column - 1,
        };
        break;
    }
    const nextTileKey = getTileKey(nextTile);
    if (!floor.has(nextTileKey)) {
      floor.set(nextTileKey, "white");
    }
    currentTile = nextTile;
  }
  flipTile(currentTile);
}

const countBlackTiles = (floor) =>
  [...floor.values()].filter((tile) => tile === "black").length;

console.log("Part 1:", countBlackTiles(floor));

const flipLivingTile = (prevFloor, floor, location) => {
  const tileKey = getTileKey(location);
  const color = prevFloor.get(tileKey) || "white";
  const adjacentBlack = getAdjacentTiles(prevFloor, location).filter(
    (_) => _ === "black"
  ).length;
  if (color === "white" && adjacentBlack === 2) {
    floor.set(tileKey, "black");
  } else if (color === "black" && (adjacentBlack === 0 || adjacentBlack > 2)) {
    floor.set(tileKey, "white");
  }
};

for (let day = 1; day <= 100; day++) {
  let nextFloor = new Map(floor.entries());

  for (const [tileKey] of floor) {
    const { row, column } = getLocation(tileKey);

    // Flip current tile
    flipLivingTile(floor, nextFloor, { row, column });

    // Flip all adjacent tiles to cover the case of 2 black tiles
    // on the border of the current floor
    flipLivingTile(floor, nextFloor, { row, column: column + 1 });
    flipLivingTile(floor, nextFloor, { row, column: column - 1 });
    flipLivingTile(floor, nextFloor, { row: row + 1, column: column + 1 });
    flipLivingTile(floor, nextFloor, { row: row + 1, column });
    flipLivingTile(floor, nextFloor, { row: row - 1, column });
    flipLivingTile(floor, nextFloor, { row: row - 1, column: column - 1 });
  }

  floor = nextFloor;

  if (env.DEBUG) {
    console.log(`Day ${day}`, countBlackTiles(floor));
  }
}

console.log("Part 2:", countBlackTiles(floor));
