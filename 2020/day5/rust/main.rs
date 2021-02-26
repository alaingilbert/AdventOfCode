use std::fs::File;
use std::io::{prelude::*, BufReader};
use std::cmp;
use std::collections::HashSet;

fn seat_id_by_boarding_pass(boarding_pass: &str) -> i32 {
    let mut out = 0;
    for c in boarding_pass.chars() {
        out <<= 1;
        if c == 'B' || c == 'R' {
            out |= 1;
        }
    }
    out
}

fn part1() {
    let file = File::open("./input").unwrap();
    let reader = BufReader::new(file);
    let mut max_seat_id = 0;
    for line in reader.lines() {
        let seat_id = seat_id_by_boarding_pass(line.unwrap().as_str());
        max_seat_id = cmp::max(max_seat_id, seat_id);
    }
    println!("{}", max_seat_id);
}

fn part2() {
    let file = File::open("./input").unwrap();
    let reader = BufReader::new(file);
    let mut cache = HashSet::new();
    cache.
    for line in reader.lines() {
        let seat_id = seat_id_by_boarding_pass(line.unwrap().as_str());
        cache.insert(seat_id);
    }
    for i in 0..1024 {
        if !cache.contains(&i) &&
            cache.contains(&(i-1)) &&
            cache.contains(&(i+1)) {
            println!("{}", i);
        }
    }
}

fn main() {
    part1();
    part2();
}
