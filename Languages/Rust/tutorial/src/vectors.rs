use std::mem;

pub fn run() {
    let mut number: Vec<i32> = vec![1, 2, 3, 4, 5];

    //Print all the values
    println!("{:?}", number);

    //Single value
    println!("{}", number[0]);

    //Reassign a value
    number[2] = 8;
    println!("{}", number[2]);

    //Get vector length
    println!("Vector length: {}", number.len());

    //Size of vector
    println!("Vector occupies {} bytes", mem::size_of_val(&number));

    //Get slice
    // let slice: &[i32] = &number;
    let slice: &[i32] = &number[0..2];//zero to second index
    println!("{:?}", slice);

    //Add element to the vecot
    number.push(9);
    number.push(10);
    println!("After pushing {:?}", number);

    //Remove last value
    number.pop();
    println!("After removing {:?}", number);

    //Loop through vector
    for x in number.iter() {
        println!("{}", x);
    }

    println!("After mutating:");
    
    //Loop and mutate value
    for x in number.iter_mut() {
        *x *= 2;
        println!("{}", x);
    }
}