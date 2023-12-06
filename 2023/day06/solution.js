const [times, distances] = input
    .split('\n')
    .map(s => s.split(" ").filter(s => s != '').slice(1).map(s => parseInt(s)));
const races = Array(times.length);
for (let i = 0; i < times.length; i++) {
    races[i] = { time: times[i], distance: distances[i] };
}

const mul = arr => arr.reduce((a, b) => a * b, 1);

const numWins = r => {
    let wins = 0;
    for (let wait = 1; wait < r.time; wait++) {
        const distance = (r.time - wait) * wait;
        if (distance > r.distance) {
            wins++;
        }
    }
    return wins;
};

console.log("Part 1", mul(races.map(numWins)));

const [time, distance] = input
    .split('\n')
    .map(s => parseInt(s.split('').filter(c => c != ' ').join("").split(":")[1]));
console.log("Part 2", numWins({ time, distance }));
