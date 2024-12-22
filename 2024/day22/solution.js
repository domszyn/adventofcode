import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

let secretNumbers = parseInput(input, l => BigInt(l));

const transform = sn => {
    sn ^= sn << 6n;
    sn %= 16777216n;
    sn ^= sn >> 5n
    sn %= 16777216n;
    sn ^= sn << 11n;
    sn %= 16777216n

    return sn;
}

const getPrice = sn => {
    const snStr = sn.toString();
    return BigInt(snStr[snStr.length - 1]);
}

const getDiffs = ([p1, p2, p3, p4, p5]) => ([
    p2 - p1,
    p3 - p2,
    p4 - p3,
    p5 - p4
]);

let buyers = new Array(secretNumbers.length);
let allBids = new Map();
for (let i = 0; i < secretNumbers.length; i++) {
    buyers[i] = {
        sn: secretNumbers[i],
        prices: [],
    };

    const buyerBids = new Set();
    for (let j = 0; j < 2000; j++) {
        if (j > 0) {
            buyers[i].sn = transform(buyers[i].sn)
        }
        const price = getPrice(buyers[i].sn);
        buyers[i].prices.push(price);
        if (buyers[i].prices.length > 5) {
            buyers[i].prices = buyers[i].prices.slice(1);
        }
        if (buyers[i].prices.length == 5) {
            const diffs = getDiffs(buyers[i].prices).join();
            if (buyerBids.has(diffs)) {
                continue;
            }
            allBids.set(diffs, price + (allBids.get(diffs) ?? 0n));
            buyerBids.add(diffs);
        }
    }
    delete (buyers[i].prices);
}

console.log("Part 1", buyers.map(b => b.sn).sumBigInt().toString());
console.log("Part 2", [...allBids.values()].max().toString());