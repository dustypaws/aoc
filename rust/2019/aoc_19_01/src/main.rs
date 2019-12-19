use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Where's the friggin' file?!");
    let masses = input
        .trim()
        .split("\n")
        .filter_map(|s| s.parse::<u32>().ok())
        .collect::<Vec<u32>>();

    println!("Part 1: {:?}", part_1(&masses));
    println!("Part 2: {:?}", part_2(&masses));
}

fn part_1(m: &Vec<u32>) -> u32 {
    m.iter().map(|x| calc_fuel(x)).sum::<u32>()
}

fn part_2(m: &Vec<u32>) -> u32 {
    m.iter().map(|x| recursive_fuel_calc(x, 0)).sum::<u32>()
}

fn recursive_fuel_calc(i: &u32, mut acc: u32) -> u32 {
    if acc == 0 {
        acc = acc.checked_add(calc_fuel(i)).unwrap_or(0);
        println!("Initial: {} -> {}", i, acc);
    }

    if *i > 0 {
        let _tmp = calc_fuel(i);
        acc = acc.checked_add(calc_fuel(i)).unwrap_or(0);
        println!("Initialized with: {} || {} -> {}", i, _tmp, acc);
        recursive_fuel_calc(&_tmp, acc);
    }
    acc
}

fn calc_fuel(i: &u32) -> u32 {
    // i / 3 - 2
    i.checked_div(3).unwrap_or(0).checked_sub(2).unwrap_or(0)
}
