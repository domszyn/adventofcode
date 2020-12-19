const fs = require("fs");

const instructions = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map((s) => ({
    action: s.charAt(0),
    value: +s.substring(1),
  }));

const rotateWaypoint = (wpx, wpy, degrees, clockwise) => {
  const coordinates = [
    { value: Math.abs(wpx), bearing: wpx >= 0 ? 90 : 270 },
    { value: Math.abs(wpy), bearing: wpy >= 0 ? 180 : 0 },
  ];

  return coordinates
    .map(({ value, bearing }) => ({
      value,
      bearing: (clockwise ? bearing + degrees : bearing - degrees + 360) % 360,
    }))
    .reduce((wp, { value, bearing }) => {
      switch (bearing) {
        case 0:
          return { ...wp, wpy: -value };
        case 90:
          return { ...wp, wpx: value };
        case 180:
          return { ...wp, wpy: value };
        case 270:
          return { ...wp, wpx: -value };
        default:
          return wp;
      }
    }, {});
};

const moveShip = ({ x, y, bearing }, { action, value }) => {
  switch (action) {
    case "N":
      return { x, y: y - value, bearing };
    case "S":
      return { x, y: y + value, bearing };
    case "E":
      return { x: x + value, y, bearing };
    case "W":
      return { x: x - value, y, bearing };
    case "L":
      return { x, y, bearing: (bearing - value + 360) % 360 };
    case "R":
      return { x, y, bearing: (bearing + value) % 360 };
    case "F":
      switch (bearing) {
        case 0:
          return moveShip({ x, y, bearing }, { action: "N", value });
        case 90:
          return moveShip({ x, y, bearing }, { action: "E", value });
        case 180:
          return moveShip({ x, y, bearing }, { action: "S", value });
        case 270:
          return moveShip({ x, y, bearing }, { action: "W", value });
        default:
          return { x, y, bearing };
      }

    default:
      return { x, y, bearing };
  }
};

const moveShipUsingWaypoint = ({ x, y, wpx, wpy }, { action, value }) => {
  switch (action) {
    case "N":
      return { x, y, wpx, wpy: wpy - value };
    case "S":
      return { x, y, wpx, wpy: wpy + value };
    case "E":
      return { x, y, wpx: wpx + value, wpy };
    case "W":
      return { x, y, wpx: wpx - value, wpy };
    case "L":
      return { x, y, ...rotateWaypoint(wpx, wpy, value, false) };
    case "R":
      return { x, y, ...rotateWaypoint(wpx, wpy, value, true) };
    case "F":
      return { x: x + value * wpx, y: y + value * wpy, wpx, wpy };
    default:
      return { x, y, wpx, wpy };
  }
};

let ship = instructions.reduce(moveShip, { x: 0, y: 0, bearing: 90 });

console.log("Part1: ", Math.abs(ship.x) + Math.abs(ship.y));

ship = instructions.reduce(moveShipUsingWaypoint, {
  x: 0,
  y: 0,
  wpx: 10,
  wpy: -1,
});

console.log("Part2: ", Math.abs(ship.x) + Math.abs(ship.y));
