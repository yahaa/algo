struct Student(i32, i32);
fn main() {
    let mut number: usize = 3;

    while number != 0 {
        println!("{}!", number);

        number = number - 1;
    }

    let a = Student(11, 22);

    println!("{} {}", a.0, a.1);
}
