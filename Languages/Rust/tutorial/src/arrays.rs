use std::mem;

pub fn run() {
    let mut number: [i32; 5] = [1, 2, 3, 4, 5];

    //Print all the values
    println!("{:?}", number);

    //Single value
    println!("{}", number[0]);

    //Reassign a value
    number[2] = 8;
    println!("{}", number[2]);

    //Get array length
    println!("{}", number.len());

    //Arrays are stack allocated
    println!("Array occupies {} bytes", mem::size_of_val(&number));

    //Get slice
    // let slice: &[i32] = &number;
    let slice: &[i32] = &number[0..2];//zero to second index
    println!("{:?}", slice);
}