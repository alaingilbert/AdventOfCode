use std::fs::File;
use std::io::{prelude::*, BufReader};
use std::cmp;
use std::collections::HashSet;

fn seat_id_by_boarding_pass(boarding_pass: &str) -> i32 {
    let rows_str = &boarding_pass[..7];
    let cols_str = &boarding_pass[7..];
    let mut rows = (1 << 7) - 1; // 127
    let mut cols = (1 << 3) - 1; // 7
    let mut row = 0;
    let mut col = 0;
    for c in rows_str.chars() {
        rows >>= 1;
        if c == 'B' {
            row += rows + 1
        }
    }
    for c in cols_str.chars() {
        cols >>= 1;
        if c == 'R' {
            col += cols + 1
        }
    }
    row * 8 + col
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
