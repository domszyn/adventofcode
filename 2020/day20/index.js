const { readFileSync } = require("fs");

const tiles = new Map(
  readFileSync(`${__dirname}/input.txt`, "utf-8")
    .split("\n\n")
    .map((block) => {
      const [firstLine, ...rest] = block.split("\n");
      const id = +firstLine.match(/\d+/);
      return [id, { id, image: rest.map((s) => Array.from(s)) }];
    })
);

const printTile = ({ image }) => {
  console.log(image.map((l) => l.join("")).join("\n"));
};

const flipTile = ({ id, image }) => {
  const flipped = [];
  const height = image.length;
  const width = (image[0] && image[0].length) || 0;

  for (let y = 0; y < height; y++) {
    flipped[y] = [];
    for (let x = 1; x <= width; x++) {
      flipped[y][width - x] = image[y][x - 1];
    }
  }

  return { id, image: flipped };
};

const emptyTile = (size) => {
  const tile = [];
  for (let y = 0; y < size; y++) {
    tile[y] = [];
    for (let x = 0; x < size; x++) {
      tile[y][x] = "";
    }
  }
  return tile;
};

const rotateTile = ({ id, image }, degree) => {
  if (degree % 360 === 0) return { id, image };
  const tileSize = image.length;
  const rotated = emptyTile(tileSize);

  for (let y = 0; y < tileSize; y++) {
    for (let x = 0; x < tileSize; x++) {
      let rx, ry;
      switch (degree % 360) {
        case 90:
          ry = x;
          rx = tileSize - y - 1;
          break;
        case 180:
          ry = tileSize - y - 1;
          rx = tileSize - x - 1;
          break;
        case 270:
          ry = tileSize - x - 1;
          rx = y;
          break;
      }
      rotated[ry][rx] = image[y][x];
    }
  }

  return { id, image: rotated };
};

const arrayEqual = (a, b) => {
  if (a.length !== b.length) return false;

  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) return false;
  }

  return true;
};

const tileLinesUpWith = ({ image }, { image: adjacentImage }, direction) => {
  if (image === undefined) return true;
  switch (direction) {
    case "up":
      return arrayEqual(image[0], adjacentImage[adjacentImage.length - 1]);
    case "down":
      return arrayEqual(image[image.length - 1], adjacentImage[0]);
    case "left":
      return arrayEqual(
        image.map((row) => row[0]),
        adjacentImage.map((row) => row[row.length - 1])
      );
    case "right":
      return arrayEqual(
        image.map((row) => row[row.length - 1]),
        adjacentImage.map((row) => row[0])
      );
    default:
      return false;
  }
};

const tileIDs = Array.from(tiles.keys()).sort((a, b) => a - b);

const indexToRowAndColumn = (idx) => [idx % 12, Math.floor(idx / 12)];
const rowAndColumnToIndex = (x, y) => y * 12 + x;

const assemblePuzzle = (puzzle) => {
  if (puzzle.length === 12 * 12) {
    return puzzle;
  }

  const [x, y] = indexToRowAndColumn(puzzle.length);

  const usedTiles = puzzle.map(({ id }) => id);

  const tileLinesUp = (tile) => {
    let linedUp = true;

    if (y > 0) {
      linedUp &&= tileLinesUpWith(
        tile,
        puzzle[rowAndColumnToIndex(x, y - 1)],
        "up"
      );
    }

    if (x > 0) {
      linedUp &&= tileLinesUpWith(
        tile,
        puzzle[rowAndColumnToIndex(x - 1, y)],
        "left"
      );
    }

    return linedUp;
  };

  const findMatches = () => {
    const matches = [];
    for (let i = 0; i < tileIDs.length; i++) {
      const id = tileIDs[i];
      if (usedTiles.includes(id)) continue;
      const tile = tiles.get(id);
      tileOrientations = [
        tile,
        flipTile(tile),
        rotateTile(tile, 90),
        rotateTile(tile, 180),
        rotateTile(tile, 270),
        rotateTile(flipTile(tile), 90),
        rotateTile(flipTile(tile), 180),
        rotateTile(flipTile(tile), 270),
      ];

      matches.push(...tileOrientations.filter(tileLinesUp));
    }

    return matches;
  };

  const matches = findMatches();

  if (matches.length === 0) {
    return [];
  }

  const combinations = matches.map((m) => assemblePuzzle([...puzzle, m]));
  let longestCombination = -1;
  for (let m = 0; m < combinations.length; m++) {
    if (combinations[m].length > longestCombination) longestCombination = m;
  }
  return combinations[longestCombination];
};

const cropBorder = ({ image }) => {
  const imageSize = image.length;
  return image
    .slice(1, imageSize - 1)
    .map((row) => row.slice(1, imageSize - 1));
};

const result = assemblePuzzle([]);
const arrangedTiles = [];
for (let y = 0; y < 12; y++) {
  arrangedTiles[y] = [];
  for (let x = 0; x < 12; x++) {
    const tile = result[rowAndColumnToIndex(y, x)];
    arrangedTiles[y][x] = cropBorder(flipTile(rotateTile(tile, 90)));
  }
}

const cornerTiles = [
  result[rowAndColumnToIndex(0, 0)].id,
  result[rowAndColumnToIndex(0, 11)].id,
  result[rowAndColumnToIndex(11, 0)].id,
  result[rowAndColumnToIndex(11, 11)].id,
];
console.log(
  "Part1:",
  cornerTiles.reduce((a, b) => a * b, 1)
);

let image = [];

for (let y = 0; y < 12 * 8; y++) {
  image[y] = [];
  for (let x = 0; x < 12 * 8; x++) {
    const tile = arrangedTiles[Math.floor(y / 8)][Math.floor(x / 8)];
    image[y][x] = tile[y % 8][x % 8];
  }
}
image = rotateTile({ image }, 180).image;
image = image.map((row) => row.join(""));

const monsterRegex = [
  /..................#./,
  /#....##....##....###/,
  /.#..#..#..#..#..#.../,
];

let monsterCount = 0;
for (let y = 0; y < image.length - 2; y++) {
  for (let x = 0; x < image.length - 20; x++) {
    if (
      monsterRegex.every((r, i) => r.test(image[y + i].substring(x, x + 20)))
    ) {
      monsterCount++;
    }
  }
}
console.log(
  "Part2: ",
  image
    .map((s) => Array.from(s).filter((s) => s === "#").length)
    .reduce((a, b) => a + b, 0) -
    15 * monsterCount
);
