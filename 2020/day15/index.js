let input = [19, 0, 5, 1, 10, 13];

do {
  let i = input.length - 1;
  const lastSpoken = input[i];
  const prevIdx = input.lastIndexOf(lastSpoken, -2);

  if (prevIdx >= 0) {
    input.push(i - prevIdx);
  } else {
    input.push(0);
  }
} while (input.length < 2020);

console.log("Part1:", input[2019]);

input = [19, 0, 5, 1, 10, 13];
const spokenNumbers = [];
spokenNumbers.length = 30000000;
for (let i = 0; i < input.length - 1; i++) {
  spokenNumbers[input[i]] = i;
}

let lastSpoken = input[input.length - 1];

for (let i = input.length; i < 30000000; i++) {
  let lastSpokenIdx = spokenNumbers[lastSpoken];
  if (lastSpokenIdx === undefined) lastSpokenIdx = i - 1;

  const tmp = i - lastSpokenIdx - 1;
  spokenNumbers[lastSpoken] = i - 1;
  lastSpoken = tmp;
}

console.log("Part2:", lastSpoken);
