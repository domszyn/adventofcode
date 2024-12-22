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
let allDiffs = new Set();
for (let i = 0; i < secretNumbers.length; i++) {
    buyers[i] = {
        sn: secretNumbers[i],
        prices: [],
    };

    const bids = new Map();
    for (let j = 0; j <= 2000; j++) {
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
            if (bids.has(diffs)) {
                continue;
            }
            allDiffs.add(diffs);
            bids.set(diffs, price);
        }
    }
    buyers[i].prices = bids;
}

let maxBid = 0;
let checked = 0;
for (const d of allDiffs) {
    let b = 0n;
    for (let i = 0; i < buyers.length; i++) {
        if (buyers[i].prices.has(d)) {
            b += buyers[i].prices.get(d);
        }
    }

    if (b > maxBid) {
        maxBid = b;
    }
    console.log(++checked, allDiffs.size, maxBid);
}

console.log("Part 1", buyers.map(b => b.sn).sumBigInt());
console.log("Part 2", maxBid);