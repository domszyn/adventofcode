const { readFileSync } = require("fs");

const parseIngredients = (food) => {
  const ingredients = food.substring(0, food.indexOf("(") - 1).split(" ");
  const allergens = food
    .substring(food.indexOf(" ", food.indexOf("(")) + 1, food.length - 1)
    .split(", ");
  return { ingredients, allergens };
};

const food = readFileSync(`${__dirname}/input.txt`, "utf-8")
  .split("\n")
  .map(parseIngredients);

const allAlergens = new Set();
const allIngredients = new Set();
let foodWithAllergens = new Map();

for (let i = 0; i < food.length; i++) {
  for (let j = 0; j < food[i].allergens.length; j++) {
    allAlergens.add(food[i].allergens[j]);
  }
  for (let j = 0; j < food[i].ingredients.length; j++) {
    const ingredient = food[i].ingredients[j];
    allIngredients.add(ingredient);
  }
}

for (let allergen of allAlergens) {
  let commonIngredients = Array.from(allIngredients);
  for (let i = 0; i < food.length; i++) {
    if (food[i].allergens.includes(allergen)) {
      commonIngredients = commonIngredients.filter((ci) =>
        food[i].ingredients.includes(ci)
      );
    }
  }
  foodWithAllergens.set(allergen, commonIngredients);
}

while (true) {
  for (let allergen of allAlergens) {
    if (foodWithAllergens.get(allergen).length === 1) {
      const singleAllergen = foodWithAllergens.get(allergen)[0];
      for (let [key, value] of foodWithAllergens) {
        if (key === allergen) continue;
        foodWithAllergens.set(
          key,
          value.filter((_) => _ !== singleAllergen)
        );
      }
    }
  }
  if ([...foodWithAllergens.values()].every((_) => _.length === 1)) break;
}

const foodWithAllergensArr = [...foodWithAllergens.values()].reduce(
  (a, b) => [...a, ...b],
  []
);

const safeIngredientsCount = food
  .map((f) => f.ingredients.filter((i) => !foodWithAllergensArr.includes(i)))
  .reduce((a, b) => a + b.length, 0);

console.log("Part 1:", safeIngredientsCount);

const canonicalDangerousIngredientList = [...foodWithAllergens.keys()]
  .sort()
  .map((i) => foodWithAllergens.get(i)[0])
  .join();

console.log("Part 2:", canonicalDangerousIngredientList);
