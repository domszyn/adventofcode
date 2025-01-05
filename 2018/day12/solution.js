import { input } from "./input.js";

let [state, rules] = input.split("\n\n");
let startIdx = 0;
state = state.split(": ")[1].split('')
rules = new Map(rules.split("\n").map(s => s.split(" => ")));

for (let gen = 0; gen < 185; gen++) {
    // console.log('Gen', gen, state.length, startIdx, state.join(''));
    while (state.indexOf('#') < 5) {
        state = ['.', ...state];
        startIdx--;
    }
    while (state.indexOf('#') > 5) {
        state = state.slice(1);
        startIdx++
    }
    while (state.lastIndexOf('#') > state.length - 6) {
        state = [...state, '.'];
    }
    while (state.lastIndexOf('#') < state.length - 6) {
        state = state.slice(0, state.length - 1);
    }
    
    let newState = new Array(state.length).fill('.');
    for (let i = 2; i < state.length - 2; i++) {
        const left = state.slice(i - 2, i).join('');
        const right = state.slice(i + 1, i + 3).join('');
        const pattern = left + state[i] + right;
        if (rules.has(pattern)) {
            newState[i] = rules.get(pattern);
        }
    }
    state = newState;
    if (gen == 19) {
        let part1 = 0;
        for (let i = 0; i < state.length; i++) {
            if (state[i] == '#') {
                part1 += i + startIdx;
            }
        }
        console.log(part1);
    }
}

let part2 = 0;
for (let i = 0; i < state.length; i++) {
    if (state[i] == '#') {
        part2 += i + startIdx + 50000000000 - 185;
    }
}
console.log(part2);