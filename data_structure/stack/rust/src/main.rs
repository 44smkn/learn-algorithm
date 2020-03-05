mod stack;

use stack::SimpleStack;

fn main() {
    println!("Hello, world!");
}

fn calc(reverse_polish_notation: &str) -> i32 {
    let mut s = SimpleStack::initialize();

    for v in reverse_polish_notation.split_whitespace() {
        match v {
            "+" => {
                let (a, b) = pop_two_elements(&mut s);
                s.push(a + b);
            }
            "-" => {
                let (a, b) = pop_two_elements(&mut s);
                s.push(b - a);
            }
            "*" => {
                let (a, b) = pop_two_elements(&mut s);
                s.push(a * b);
            }
            _ => s.push(v.parse().unwrap()),
        }
    }
    s.pop()
}

fn pop_two_elements(s: &mut SimpleStack) -> (i32, i32) {
    let a = s.pop();
    let b = s.pop();
    (a, b)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn calculate_reverse_polish_notation() {
        let expression = "1 2 + 3 4 - *";
        assert_eq!(calc(expression), -3);
    }
}
