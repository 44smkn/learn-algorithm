fn main() {
    let mut arr: Vec<i64> = vec![5, 2, 4, 6, 1, 3];
    sort(&mut arr);
    println!("ソート後の配列：{:?}", arr)
}

fn sort(arr: &mut Vec<i64>) {
    // iは未ソート部分列の先頭インデックス
    for i in 1..arr.len() {
        let key = arr[i];
        let mut key_index = 0;
        // jはソート済み部分列からkeyを挿入するためのループ変数。i-1から0までをループする。
        for j in (0..i).rev() {
            if arr[j] < key {
                // keyの方が大きい場合には、今のjのインデックスの右隣となるため+1している
                key_index = j + 1;
                break;
            }
            arr[j + 1] = arr[j];
        }
        // keyがソート済配列のどの要素よりも小さかった場合には、breakに入らず、インデックスが0となり、最初の要素にkeyが入る
        arr[key_index] = key;
    }
}
