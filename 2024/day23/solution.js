import { input } from "./input.js";

let connected = new Set();
const connections = new Map();

for (const c of input.split('\n')) {
    const [c1, c2] = c.split('-');
    if (!connections.has(c1)) connections.set(c1, new Set());
    if (!connections.has(c2)) connections.set(c2, new Set());
    connections.set(c1, new Set([...connections.get(c1), c2]));
    connections.set(c2, new Set([...connections.get(c2), c1]));
    connected.add([c1, c2].sort().join());
}

const findConnected = (connected) => {
    const connectedComputers = [...connected.keys()].map(cs => cs.split(','));
    const moreConnected = new Set();
    for (const c of connections.keys()) {
        for (const cc of connectedComputers) {
            if (cc.every(ccc => connections.get(ccc).has(c))) {
                let newParty = [...cc, c];
                newParty.sort();
                moreConnected.add(newParty.join());
            }
        }
    }

    return moreConnected;
}

connected = findConnected(connected); // three connected
console.log("Part1", [...connected.keys()].filter(c => c.startsWith('t') || c.indexOf(',t') > 0).length);

while (true) {
    const moreConnected = findConnected(connected);
    if (moreConnected.size > 0) {
        connected = moreConnected;
    } else {
        break;
    }
}

console.log("Part2", [...connected.keys()][0]);