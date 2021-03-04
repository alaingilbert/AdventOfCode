#![feature(iterator_fold_self)]

#[derive(Debug)]
struct Bus {
    id: i64,
    offset: i64,
}

impl Bus {
    fn next_departure(&self, cur_time: i64) -> i64 {
        self.id * (cur_time as f64 / self.id as f64).ceil() as i64
    }
}

fn get_buses(buses_str: &str) -> Vec<Bus> {
    let buses: Vec<Bus> = buses_str
        .split(',')
        .enumerate()
        .map(|(idx, s)| Bus {
            id: s.parse().unwrap_or(0),
            offset: idx as i64,
        })
        .filter(|b| b.id != 0)
        .collect();
    buses
}

fn part1(buses_str: &str) {
    let t = 1004098;
    let buses = get_buses(buses_str);
    let min_bus = buses
        .iter()
        .min_by(|x, y| x.next_departure(t).cmp(&y.next_departure(t)))
        .unwrap();
    let wait_time = min_bus.next_departure(t) - t;
    println!("{}", min_bus.id * wait_time);
}

fn solve(bus1: &Bus, bus2: &Bus) -> i64 {
    let mut out = 0;
    let mut a = bus2.offset - bus1.offset;
    loop {
        if a % bus2.id == 0 {
            return out;
        }
        a += bus1.id;
        out += 1;
    }
}

fn merge_buses(bus1: &Bus, bus2: &Bus) -> Bus {
    let x1 = solve(bus1, bus2);
    Bus {
        id: bus1.id * bus2.id,
        offset: bus1.offset - (bus1.id * x1),
    }
}

fn part2(buses_str: &str) {
    let buses = get_buses(buses_str);
    let res = buses.into_iter().fold_first(|acc, b| merge_buses(&acc, &b));
    println!("{:?}", -res.unwrap().offset);
    println!("done");
}

fn main() {
    //let buses_str = "7,13,x,x,59,x,31,19";
    let buses_str = "23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,509,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29,x,401,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,19";
    part1(buses_str);
    part2(buses_str);
}
