const input = 872495136;

const play = (cups, moves) => {
  let currentCupIdx = 0;
  console.time("game");
  for (let i = 1; i <= moves; i++) {
    if (i % 10000 === 0) {
      console.log(i);
      console.timeLog("game")
    }
    let currentCup = cups[currentCupIdx];
    let pickedCups = cups.splice(currentCupIdx + 1, 3);
    if (pickedCups.length < 3) {
      pickedCups.push(...cups.splice(0, 3 - pickedCups.length));
    }
    // console.timeLog("game", "pick cups");
    let destinationCupIdx;
    let minLabel = 1;
    let maxLabel = cups.length;
    while (
      pickedCups[0] === minLabel ||
      pickedCups[1] === minLabel ||
      pickedCups[2] === minLabel
    )
      minLabel++;
    while (
      pickedCups[0] === maxLabel ||
      pickedCups[1] === maxLabel ||
      pickedCups[2] === maxLabel
    )
      maxLabel--;
    for (let offset = 1; ; offset++) {
      let label = currentCup - offset;
      if (label < minLabel) label = maxLabel;
      if (pickedCups.includes(label)) continue;
      destinationCupIdx = cups.indexOf(label);
      if (destinationCupIdx >= 0) break;
    }
    cups.splice(
      destinationCupIdx + 1,
      0,
      pickedCups[0],
      pickedCups[1],
      pickedCups[2]
    );
    const currentCupOffset = cups.indexOf(currentCup) - currentCupIdx;

    if (currentCupOffset > 0) {
      cups.push(...cups.splice(0, currentCupOffset));
    }
    currentCupIdx = (currentCupIdx + 1) % cups.length;
  }
  while (cups[0] !== 1) {
    cups.push(cups.shift());
  }

  return cups.slice(1);
};

let cups = Array.from(input.toString()).map((s) => +s);
console.log("Part 1:", play(cups, 100).join(""));

