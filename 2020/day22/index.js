const { readFileSync } = require("fs");
const [player1, player2] = readFileSync(`${__dirname}/input.txt`, "utf-8")
  .split("\n\n")
  .map((block) =>
    block
      .split("\n")
      .slice(1)
      .map((s) => +s)
  );

const calculateScore = (deck) => {
  let score = 0;
  for (let i = 0; i < deck.length; i++) {
    score += deck[i] * (deck.length - i);
  }
  return score;
};

const combat = () => {
  const deck1 = [...player1];
  const deck2 = [...player2];

  do {
    const card1 = deck1.shift();
    const card2 = deck2.shift();

    if (card1 > card2) {
      deck1.push(card1, card2);
    } else if (card2 > card1) {
      deck2.push(card2, card1);
    }
  } while (deck1.length > 0 && deck2.length > 0);

  return Math.max(...[deck1, deck2].map(calculateScore));
};

const serializeDecks = (deck1, deck2) =>
  [deck1, deck2].map((d) => d.join()).join("#");

const winners = new Map();
const recursiveCombat = (deck1, deck2, level = 1) => {
  const prevRounds = new Set();
  do {
    const roundKey = serializeDecks(deck1, deck2);
    if (prevRounds.has(roundKey)) {
      return 1;
    }
    prevRounds.add(roundKey);

    const card1 = deck1.shift();
    const card2 = deck2.shift();

    if (deck1.length >= card1 && deck2.length >= card2) {
      let result;
      const subgameDeck1 = deck1.slice(0, card1);
      const subgameDeck2 = deck2.slice(0, card2);
      const subgameKey = serializeDecks(subgameDeck1, subgameDeck2);
      if (winners.has(subgameKey)) {
        result = winners.get(subgameKey);
      } else {
        result = recursiveCombat(subgameDeck1, subgameDeck2, level + 1);
        winners.set(subgameKey, result);
      }

      if (result === 1) {
        deck1.push(card1, card2);
      } else {
        deck2.push(card2, card1);
      }
    } else {
      if (card1 > card2) {
        deck1.push(card1, card2);
      } else if (card2 > card1) {
        deck2.push(card2, card1);
      }
    }
  } while (deck1.length > 0 && deck2.length > 0);

  if (level == 1) {
    return Math.max(...[deck1, deck2].map(calculateScore));
  } else if (deck1.length > 0) {
    return 1;
  } else {
    return 2;
  }
};

console.log("Part 1: ", combat());
console.log("Part 2:", recursiveCombat([...player1], [...player2]));
