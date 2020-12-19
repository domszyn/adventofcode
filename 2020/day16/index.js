const { readFileSync } = require("fs");

const [ruleBlock, myTicketBlock, nearbyTicketsBlock] = readFileSync(
  `${__dirname}/input.txt`,
  "utf-8"
).split("\n\n");

const getRuleKey = (ruleName) =>
  ruleName
    .split(" ")
    .map((w, i) => (i === 0 ? w : w.charAt(0).toUpperCase() + w.substring(1)))
    .join("");

const parseRule = (rule) => {
  const [ruleName, rest] = rule.split(": ");
  const validValues = rest
    .split(" or ")
    .map((range) => range.split("-").map((v) => +v));
  return { key: getRuleKey(ruleName), validValues };
};

const rules = ruleBlock.split("\n").map(parseRule);
const nearbyTickets = nearbyTicketsBlock
  .split("\n")
  .filter((_, idx) => idx > 0)
  .map((s) => s.split(",").map((s) => +s));
const myTicket = myTicketBlock
  .split("\n")[1]
  .split(",")
  .map((s) => +s);

const isValid = (val, { validValues }) =>
  validValues.some(([from, to]) => val >= from && val <= to);

const errorRate = nearbyTickets
  .map((ticket) =>
    ticket
      .filter((value) => rules.findIndex((rule) => isValid(value, rule)) === -1)
      .reduce((a, b) => a + b, 0)
  )
  .reduce((a, b) => a + b, 0);
console.log("Part1:", errorRate);

const validTickets = [
  myTicket,
  ...nearbyTickets
    .map((ticket) =>
      ticket.filter(
        (value) => rules.findIndex((rule) => isValid(value, rule)) === -1
      ).length === 0
        ? ticket
        : undefined
    )
    .filter(Boolean),
];

const fields = [];

for (let i = 0; i < myTicket.length; i++) {
  const ticketValues = validTickets.map((ticket) => ticket[i]);
  const matchingRulesIdx = rules
    .map((rule, idx) =>
      ticketValues.every((val) => isValid(val, rule)) ? idx : undefined
    )
    .filter((idx) => idx !== undefined);
  fields[i] = matchingRulesIdx;
}

let eliminated = new Set();

while (fields.some((f) => f.length > 1)) {
  fields
    .map((f, idx) => (f.length === 1 ? idx : undefined))
    .filter((idx) => idx !== undefined)
    .forEach((idx) => {
      const [removeValue] = fields[idx];
      if (!eliminated.has(removeValue)) {
        for (let j = 0; j < fields.length; j++) {
          if (j === idx) continue;

          fields[j] = fields[j].filter((_) => _ !== removeValue);
        }
        eliminated.add(removeValue);
      }
    });
}

const answer = fields
  .map(([val], idx) => [rules[val], idx])
  .filter(([rule]) => rule.key.startsWith("departure"))
  .map(([, idx]) => myTicket[idx])
  .reduce((a, b) => a * b, 1);
console.log("Part2:", answer);
