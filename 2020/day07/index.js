const fs = require("fs");

const rules = fs.readFileSync(__dirname + "/input.txt", "utf-8").split("\n");

const bags = new Map();

rules.forEach((rule) => {
  const [parentBagColor] = rule.match(/^\w+\s\w+/);
  const childrenBags = rule.match(/\d+\s\w+\s\w+/g);

  if (!bags.has(parentBagColor)) {
    bags.set(parentBagColor, {
      color: parentBagColor,
      quantity: 1,
      parents: new Set(),
      children: new Set(),
    });
  }

  if (childrenBags) {
    const parentBag = bags.get(parentBagColor);
    childrenBags.forEach((child) => {
      const color = child.substring(child.indexOf(" ") + 1);
      const quantity = +child.substring(0, child.indexOf(" "));

      if (!bags.has(color)) {
        bags.set(color, {
          color,
          quantity: 1,
          parents: new Set(),
          children: new Set(),
        });
      }

      const childBag = bags.get(color);
      childBag.parents.add(parentBagColor);
      parentBag.children.add({
        ...childBag,
        quantity,
      });
    });
  }
});

const findAllParents = (bagColor) => {
  const bag = bags.get(bagColor);
  const parents = Array.from(bag.parents);
  if (parents.length === 0) return parents;

  return parents
    .map((p) => findAllParents(p))
    .reduce((a, b) => Array.from(new Set([...a, ...b])), parents);
};

const countChildren = (bagColor) => {
  const bag = bags.get(bagColor);
  if (!bag || bag.children.size === 0) return 0;

  return Array.from(bag.children).reduce(
    (a, { quantity, color }) => a + quantity * (countChildren(color) + 1),
    0
  );
};

console.log("Part1: ", findAllParents("shiny gold").length);
console.log("Part2: ", countChildren("shiny gold"));
