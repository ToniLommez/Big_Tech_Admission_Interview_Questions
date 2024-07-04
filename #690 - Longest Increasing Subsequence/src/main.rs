use std::cmp::max;
use std::collections::HashMap;
use std::io::Write;
use std::{io, vec};

fn main() {
    let mut input = String::new();

    print!("Enter the sequence separated by commas (e.g., '0, 8, 4, 12'): ");
    io::stdout().flush().unwrap();

    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line");

    let numbers: Result<Vec<i32>, _> = input
        .split(',')
        .map(|s| match s.trim().parse::<i32>() {
            Ok(num) => Ok(num),
            Err(e) => Err(format!("invalid number: {}", e)),
        })
        .collect();

    match numbers {
        Ok(nums) => {
            let (a, b, size, table) = longest_increasing_subsequence(nums);
            println!();
            println!("Longest Increasing Subsequence");
            println!("Size: {}", size);
            print_sequence(&a, &table);
            print_table(&a, &b, &table);
        }
        Err(e) => {
            println!("Error: {}", e)
        }
    }
}

/// Computes the Longest Increasing Subsequence (LIS) of a given vector of integers.
///
/// This function calculates the LIS using dynamic programming and memoization by doing
/// an encapsulate call of LCS using both the original array and it sorted, so it can find
/// the longest only increasing subsequence.
///
/// # Arguments
///
/// * `a` - A vector of integers for which the Longest Increasing Subsequence is to be found.
///
/// # Returns
///
/// - `Vec<i32>`: The original vector.
/// - `Vec<i32>`: The sorted version of the original vector.
/// - `usize`: The length of the Longest Increasing Subsequence.
/// - `Vec<Vec<i32>>`: The dynamic programming table used for computing the LIS.
fn longest_increasing_subsequence(a: Vec<i32>) -> (Vec<i32>, Vec<i32>, usize, Vec<Vec<i32>>) {
    let mut memo = HashMap::new();
    let mut b = a.clone();
    let n = a.len() as isize;
    let mut table: Vec<Vec<i32>> = vec![vec![-1; (n + 1) as usize]; (n + 1) as usize];

    b.sort();
    let res = lcs(&a, &b, n - 1, n - 1, &mut memo, &mut table);

    (a, b, res, table)
}

/// Recursive part of Longest Common Subsequences (LCS) for a vector of integers.
fn lcs(
    a: &Vec<i32>,
    b: &Vec<i32>,
    i: isize,
    j: isize,
    memo: &mut HashMap<(isize, isize), usize>,
    table: &mut Vec<Vec<i32>>,
) -> usize {
    if let Some(&result) = memo.get(&(i, j)) {
        return result;
    }

    let result = {
        if i == -1 || j == -1 {
            0
        } else if a[i as usize] == b[j as usize] {
            lcs(a, b, i - 1, j - 1, memo, table) + 1
        } else {
            max(
                lcs(a, b, i, j - 1, memo, table),
                lcs(a, b, i - 1, j, memo, table),
            )
        }
    };

    memo.insert((i, j), result);
    table[(i + 1) as usize][(j + 1) as usize] = result as i32;
    result
}

/// Recover and print the longest sequence given an dynamic matrix based on LCS.
///
/// # Arguments
///
/// * `a` - The original vector to retrieve the value.
/// * `table` - The dynamic table.
fn print_sequence(a: &Vec<i32>, table: &Vec<Vec<i32>>) {
    let mut seq: Vec<i32> = vec![];
    let mut i = table.len() - 1;
    let mut j = i;

    // Recover the original path
    while i != 0 && j != 0 {
        let x = table[i][j];
        if table[i - 1][j] == x {
            i -= 1;
        } else if table[i][j - 1] == x {
            j -= 1;
        } else {
            i -= 1;
            j -= 1;
            seq.push(a[i]);
        }
    }

    print!("Sequence: ");
    seq.reverse();
    for (i, num) in seq.iter().enumerate() {
        if i == seq.len() - 1 {
            print!("{}", num);
        } else {
            print!("{}, ", num);
        }
    }
    println!();
}

/// Given a dynamic table, print it.
///
/// # Arguments
///
/// * `a` - The first vector;
/// * `b` - The second vector;
/// * `table` - The table itself;
fn print_table(a: &Vec<i32>, b: &Vec<i32>, table: &Vec<Vec<i32>>) {
    println!("Table: ");
    print!("         ");
    for number in a {
        print!("{:3}", number);
    }
    println!();

    print!("---------");
    for _ in a {
        print!("---");
    }
    println!();

    let mut i = 0;
    for row in table {
        if i == 0 {
            print!("    | ");
        } else {
            print!("{:3} | ", b[i - 1]);
        }
        for value in row {
            if *value == -1 {
                print!("  -");
            } else {
                print!("{:3}", value);
            }
        }
        println!();
        i += 1;
    }
}
