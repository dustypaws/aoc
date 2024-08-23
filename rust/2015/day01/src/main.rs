use core::str;

fn main() {
    let input = include_str!("../input.txt");
    part_1(input);
    part_2(input);
}

fn part_2(input: &'static str) {
    let mut level = 1isize;
    for (i, x) in input.chars().enumerate() {
        match x {
            ')' => level -= 1,
            '(' => level += 1,
            _ => (),
        }

        if level < 0 {
            println!("Part 2: {}", i);
            break;
        }
    }
}

fn part_1(input: &'static str) {
    input.chars().fold(0, |acc, x| match x {
        ')' => acc - 1,
        '(' => acc + 1,
        '\n' => {
            println!("Part 1 - the hacky way...");
            println!("{}", acc);
            0
        }
        _ => unreachable!(),
    });
}
