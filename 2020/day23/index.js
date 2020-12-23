const input = Array.from("872495136").map((s) => +s);

const play = (input, padInput, moves, outputLength = 1) => {
  let cups = [];
  cups.length = padInput + 1;
  for (let i = 1; i <= padInput; i++) {
    const label = i > input.length ? i : input[i - 1];
    let next = i + 1 > input.length ? i + 1 : input[i];
    if (i === padInput) next = input[0];
    cups[label] = {
      label,
      next,
    };
  }

  let move = 1;
  let currentCup = cups[input[0]];

  do {
    // Pick 3 cups immediately after the current one
    let pickedCups = [cups[currentCup.next]];
    while (pickedCups.length < 3) {
      pickedCups.push(cups[pickedCups[pickedCups.length - 1].next]);
    }

    // Point current cup to the one after the last of the 3 picked cups
    currentCup.next = pickedCups[2].next;

    // Get lowest and highest available labels
    let minLabel = 1;
    let maxLabel = cups.length - 1;
    while (
      pickedCups[0].label === minLabel ||
      pickedCups[1].label === minLabel ||
      pickedCups[2].label === minLabel
    )
      minLabel++;
    while (
      pickedCups[0].label === maxLabel ||
      pickedCups[1].label === maxLabel ||
      pickedCups[2].label === maxLabel
    )
      maxLabel--;

    // Find destination cup
    let destinationLabel = currentCup.label - 1;
    for (;;) {
      if (destinationLabel < minLabel) destinationLabel = maxLabel;
      if (
        pickedCups[0].label !== destinationLabel &&
        pickedCups[1].label !== destinationLabel &&
        pickedCups[2].label !== destinationLabel
      )
        break;
      destinationLabel--;
    }
    let destinationCup = cups[destinationLabel];

    // Point last of the 3 picked cups to the cup that follows
    // the destination cup
    pickedCups[2].next = destinationCup.next;

    // Point destination cup to the first of the picked cups
    destinationCup.next = pickedCups[0].label;

    // Proceed to the cup immediately after the current
    currentCup = cups[currentCup.next];
    move++;
  } while (move <= moves);

  let answers = [cups[1].next];
  while (answers.length < outputLength) {
    answers.push(cups[answers[answers.length - 1]].next);
  }

  return answers;
};

console.log("Part 1:", play(input, 9, 100, 8).join(""));
console.log(
  "Part 2:",
  play(input, 1000000, 10000000, 2).reduce((a, b) => a * b, 1)
);
