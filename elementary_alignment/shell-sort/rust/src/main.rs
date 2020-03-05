fn main() {
    println!("Hello, world!");
}

fn shell_sort(arr: &mut Vec<i32>, gs: Vec<usize>) {
    for g in gs {
        insertion_sort(arr, g);
    }
}

fn insertion_sort(arr: &mut Vec<i32>, g: usize) {
    for i in g..arr.len() {
        let key = arr[i];
        let mut key_index = i; // keyの値のindexを記録する
        let length = arr.len();

        // [0] [1] [2] [3] [4,0] [5,1] のような配列を作る
        for j in (0..(length / g))
            .map(|x| x * g + i % g)
            .filter(|&x| x < i)
            .rev()
        {
            if arr[j] > key {
                arr[j + g] = arr[j];
                key_index = j
            }
        }
        arr[key_index] = key;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sort_num_ascending() {
        let mut arr = vec![5, 1, 4, 3, 2, 7, 9, 8, 6, 10];
        let gs = vec![4, 1];
        shell_sort(&mut arr, gs);
        assert_eq!(arr, vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
    }
}
