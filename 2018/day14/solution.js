import { numRecipes } from "./input.js";

let scoreboard = new Map([[0, 3], [1, 7]]);
let elves = [0, 1];

const printScoreboard = board => [...board.keys()].map(k => board.get(k)).join('');

const nr = numRecipes.toString();
let recipes = '';
const getRecipe = i => (scoreboard.has(i) ? scoreboard.get(i) : +recipes[i]);
while (true) {
    const newRecipes = (getRecipe(elves[0]) + getRecipe(elves[1])).toString();
    for (const r of newRecipes.split('').map(Number)) {
        scoreboard.set(scoreboard.size + recipes.length, r);
    }

    elves = elves.map(i => (i + 1 + getRecipe(i)) % (scoreboard.size + recipes.length));

    if (scoreboard.size >= 100000) {
        recipes += printScoreboard(scoreboard);
        scoreboard.clear();
        if (recipes.indexOf(nr) > 0) {
            break;
        }
    }
}

console.log("Part 1", recipes.slice(numRecipes, numRecipes + 10));
console.log("Part 2", recipes.indexOf(nr));