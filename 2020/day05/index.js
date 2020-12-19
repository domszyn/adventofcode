const fs = require("fs");

const translateRow = (ticket) => {
  let min = 0,
    max = (1 << 7) - 1;
  for (let i = 0; i < 7; i++) {
    const location = ticket.charAt(i);
    if (location === "B") {
      min = (max + min + 1) >> 1;
    } else if (location === "F") {
      max = (max + min) >> 1;
    }
  }
  if (min !== max) throw Error("could not locate the row");
  return min;
};

const translateSeat = (ticket) => {
  let min = 0,
    max = (1 << 3) - 1;
  for (let i = 0; i < 3; i++) {
    const location = ticket.charAt(i + 7);
    if (location === "R") {
      min = (max + min + 1) >> 1;
    } else if (location === "L") {
      max = (max + min) >> 1;
    }
  }
  if (min !== max) throw Error("could not locate the row");
  return min;
};

const getSeatID = (row, seat) => (row << 3) + seat;

const translateTicket = (ticket) => {
  const row = translateRow(ticket);
  const seat = translateSeat(ticket);
  return {
    row,
    seat,
    seatID: getSeatID(row, seat),
  };
};

const tickets = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map(translateTicket);

const findAvailableSeat = () => {
  const seatIDs = tickets.map(({ seatID }) => seatID).sort((a, b) => a - b);
  const rows = tickets.map(({ row }) => row);
  for (let row = 1; row < Math.max(...rows); row++) {
    for (let seat = 0; seat < 1 << 3; seat++) {
      const seatID = getSeatID(row, seat);
      if (
        !seatIDs.includes(seatID) &&
        seatIDs.includes(seatID - 1) &&
        seatIDs.includes(seatID + 1)
      ) {
        return seatID;
      }
    }
  }
};

console.log("Part1: ", Math.max(...tickets.map(({ seatID }) => seatID)));
console.log("Part2: ", findAvailableSeat());
