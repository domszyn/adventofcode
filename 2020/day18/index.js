const { readFileSync } = require("fs");

const expressions = readFileSync(`${__dirname}/input.txt`, "utf-8").split("\n");

const parse = (expr, addFirst) => {
  for (;;) {
    let openParenthesisIdx = expr.indexOf("(");
    if (openParenthesisIdx < 0) break;
    let closeParenthesisIdx;
    let level = 1;
    for (let i = openParenthesisIdx + 1; i < expr.length; i++) {
      if (expr.charAt(i) === "(") level++;
      if (expr.charAt(i) === ")") level--;
      if (level === 0) {
        closeParenthesisIdx = i;
        break;
      }
    }

    expr =
      expr.substring(0, openParenthesisIdx) +
      parse(expr.substring(openParenthesisIdx + 1, closeParenthesisIdx), addFirst) +
      expr.substring(closeParenthesisIdx + 1);
  }

  if (addFirst) {
    let literals = expr.split(" ");

    while (literals.indexOf("+") > 0) {
      const plusIdx = literals.indexOf("+");
      literals = [
        ...literals.slice(0, plusIdx - 1),
        +literals[plusIdx - 1] + +literals[plusIdx + 1],
        ...literals.slice(plusIdx + 2),
      ];
    }

    return literals.filter((s) => s !== "*").reduce((a, b) => a * b, 1);
  } else {
    const add = (a, b) => a + b;
    const multiply = (a, b) => a * b;

    let nextFn;
    const literals = expr.split(" ");
    return literals.reduce((acc, literal) => {
      switch (literal) {
        case "*":
          nextFn = multiply;
          return +acc;
        case "+":
          nextFn = add;
          return +acc;
        default:
          return nextFn ? nextFn.apply(this, [acc, +literal]) : +literal;
      }
    });
  }
};

const part1Answer = [...expressions]
  .map((expr) => parse(expr, false))
  .reduce((a, b) => a + b, 0);
const part2Answer = [...expressions]
  .map((expr) => parse(expr, true))
  .reduce((a, b) => a + b, 0);

console.log("Part1:", part1Answer);
console.log("Part2:", part2Answer);
