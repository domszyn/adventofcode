const { readFileSync } = require("fs");

const [cardPublicKey, doorPublicKey] = readFileSync(
  `${__dirname}/input.txt`,
  "utf-8"
)
  .split("\n")
  .map((s) => +s);

const transformSubjectNumber = (subjectNumber, loopSize, knownValue) => {
  let value = 1;
  for (let i = 0; i < loopSize; i++) {
    value *= subjectNumber;
    value %= 20201227;

    if (value === knownValue) return i + 1;
  }

  return value;
};

const cardLoopSize = transformSubjectNumber(
  7,
  Number.MAX_SAFE_INTEGER,
  cardPublicKey
);

const encryptionKey = transformSubjectNumber(doorPublicKey, cardLoopSize);

console.log(encryptionKey);
