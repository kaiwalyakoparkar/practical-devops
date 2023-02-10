// Can contain values of different types
// Max 12 elements
pub fn run() {
    let person: (&str, &str, i8) = ("Kaiwalya", "India", 19);

    println!("{} is from {} and has age {}", person.0, person.1, person.2);
}