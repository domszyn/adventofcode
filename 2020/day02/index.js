const fs = require("fs");

const parsePassword = (definition) => {
  const [policy, password] = definition.split(": ");
  const [range, symbol] = policy.split(" ");
  const [min, max] = range.split("-");
  return {
    password,
    policy: {
      symbol,
      min: +min,
      max: +max,
    },
  };
};

const isValid = ({ password, policy }) => {
  let symbolCount = 0;
  for (let i = 0; i < password.length; i++) {
    symbolCount += password.charAt(i) === policy.symbol ? 1 : 0;
  }
  return symbolCount >= policy.min && symbolCount <= policy.max;
};

const isOfficialyValid = ({
  password,
  policy: { min, max, symbol }
}) => {
  let firstChar = password.charAt(min - 1);
  let secondChar = password.charAt(max - 1);
  return (
    (firstChar === symbol || secondChar === symbol) && firstChar !== secondChar
  );
};

const passwords = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n")
  .map(parsePassword);

console.log("Part1: ", passwords.filter(isValid).length);
console.log("Part1: ", passwords.filter(isOfficialyValid).length);
