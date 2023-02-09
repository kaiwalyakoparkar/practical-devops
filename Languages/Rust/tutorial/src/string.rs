//Primitive string = Immutable fixed length string
// String = Heap-allocated data structures

pub fn run() {
    //Primitive type
    let msg1 = "Hello";
    println!("{}",msg1);

    //The second string type
    let mut msg2 = String::from("Hello 2 ");

    //Add character to the string
    msg2.push('W');

    //Add string to the string
    msg2.push_str("orld");

    //Check if string empty
    println!("Empty: {}", msg2.is_empty());

    //Check if the string contains
    println!("String contains 'World' {}", msg2.contains("World"));

    //Replace
    println!("Replace: {}", msg2.replace("World", "Kaiwalya"));

    println!("{}",msg2);

    //Get Length
    println!("Length is {}", msg2.len());

    //Pushing to string with capacity
    let mut s = String::with_capacity(10);
    s.push('a');
    s.push('b');

    //Assertion testing - It will output nothing if it is true
    assert_eq!(2, s.len());
}