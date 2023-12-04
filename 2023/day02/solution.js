const input = ``;

var games = input.split('\n').map(s => {
    var [s1, s2] = s.split(": ");
    var [id] = s1.match(/\d+/);
    var bag = s2
        .split("; ")
        .map(t => t.split(", ")
            .map(c => c.split(" "))
            .map(([num, color]) => ({ [color]: parseInt(num) }))
            .reduce((a, b) => ({ ...a, ...b }), { blue: 0, green: 0, red: 0 })
        ).reduce((b1, b2) => ({
            blue: Math.max(b1.blue, b2.blue),
            green: Math.max(b1.green, b2.green),
            red: Math.max(b1.red, b2.red)
        }));

    return ({
        id: parseInt(id),
        bag
    });
})

const bag = {
    red: 12,
    green: 13,
    blue: 14
};

const possibleGames = games.filter(g => g.bag.blue <= bag.blue && g.bag.green <= bag.green && g.bag.red <= bag.red);
const sum = arr => arr.reduce((a, b) => a + b, 0);

console.log("Part 1", sum(possibleGames.map(({ id }) => id)));
console.log("Part 2", sum(games.map(({ bag: { blue, green, red } }) => blue * green * red)));