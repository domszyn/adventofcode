const fs = require("fs");

const [ruleList, messageList] = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n\n");
const parsedRules = new Map();
const rules = new Map(
  ruleList.split("\n").map((ruleStr) => {
    const [index, ruleTxt] = ruleStr.split(": ");
    if (ruleTxt.startsWith('"')) {
      parsedRules.set(+index, ruleTxt.charAt(1));
      return [+index, ruleTxt.charAt(1)];
    } else {
      return [+index, ruleTxt];
    }
  })
);

const buildRule = (idx, n = 1) => {
  if (idx === 11 && n > 1) {
    let rule42 = buildRule(42);
    let rule31 = buildRule(31);

    return `${rule42}{${n}}${rule31}{${n}}`;
  }

  if (parsedRules.has(idx)) {
    let parsed = parsedRules.get(idx);
    if (idx === 8 && n > 1) parsed = `${parsed}{${n}}`;
    return parsed;
  }

  const rule = rules.get(idx);
  const rulePipes = rule.split(" | ");
  let parsed = rulePipes
    .map((rp) =>
      rp
        .split(" ")
        .map((rps) => buildRule(+rps))
        .join(" ")
    )
    .join(" | ")
    .replaceAll(" ", "");

  if (rulePipes.length > 1) parsed = `(${parsed})`;

  parsedRules.set(idx, parsed);
  return parsed;
};

let matchCount = 0;
for (let i = 1; i < 100; i++) {
  let prevMatchCount = matchCount;
  for (let j = 1; j < 100; j++) {
    const messageRegex = new RegExp(`^${buildRule(8, i)}${buildRule(11, j)}$`);
    const matches = messageList
      .split("\n")
      .map((m) => (messageRegex.test(m) ? 1 : 0))
      .reduce((a, b) => a + b, 0);

    if (i === 1 && j === 1) {
      console.log("Part1:", matches);
    }
    if (matches === 0) break;
    matchCount += matches;
  }
  if (matchCount === prevMatchCount) break;
}

console.log("Part2:", matchCount);
