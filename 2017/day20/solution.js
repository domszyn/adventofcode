import { input } from "./input.js";
import { parseInput } from "../utils.js";

class Point3D {
    constructor(x, y, z) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    add(point) {
        return new Point3D(this.x + point.x, this.y + point.y, this.z + point.z);
    }

    static parse(s) {
        const [x, y, z] = s.slice(3, s.length - 1).split(',').map(Number);
        return new Point3D(x, y, z);
    }

    distance() {
        return Math.abs(this.x) + Math.abs(this.y) + Math.abs(this.z);
    }

    equals(point) {
        return this.x == point.x && this.y == point.y && this.z == point.z;
    }
}

class Particle {
    constructor(id, position, velocity, acceleration) {
        this.id = id;
        this.position = position;
        this.velocity = velocity;
        this.acceleration = acceleration;
        this.distance = this.position.distance();
    }

    tick() {
        this.velocity = this.velocity.add(this.acceleration);
        this.position = this.position.add(this.velocity);
        this.distance = this.position.distance();
    }
}

let particles = parseInput(input, (l, i) => {
    const [p, v, a] = l.split(', ');
    return new Particle(i, Point3D.parse(p), Point3D.parse(v), Point3D.parse(a));
});

let closestToZero = [];

while (true) {
    for (const p of particles) {
        p.tick();
    }

    const closest = particles.toSorted((a, b) => a.distance - b.distance)[0].id;
    closestToZero.push(closest);
    if (closestToZero.length > 200) {
        closestToZero.shift();

        if (closestToZero.every(e => e == closestToZero[0])) {
            console.log(closestToZero[0]);
            break;
        }
    }
}

particles = parseInput(input, (l, i) => {
    const [p, v, a] = l.split(', ');
    return new Particle(i, Point3D.parse(p), Point3D.parse(v), Point3D.parse(a));
});

let collisionCounts = [];

while (true) {
    for (const p of particles) {
        p.tick();
    }

    let collisions = new Set();
    for (let i = 0; i < particles.length; i++) {
        for (let j = i + 1; j < particles.length; j++) {
            if (particles[i].position.equals(particles[j].position)) {
                collisions.add(i);
                collisions.add(j);
            }
        }
    }

    particles = particles.filter((_, i) => !collisions.has(i));
    collisionCounts.push(collisions.size);

    if (collisionCounts.length > 200) {
        collisionCounts.shift();

        if (collisionCounts.every(e => e == 0)) {
            console.log(particles.length);
            break;
        }
    }
}