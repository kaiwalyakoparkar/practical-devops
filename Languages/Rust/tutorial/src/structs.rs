struct Color {
    red: i32,
    green: i32,
    blue: i32,
}

//Tuple struct
struct Color_tuple(i32, i32, i32);

//Structure with member function
struct Person {
    f_name: String, 
    l_name: String,
}

impl Person {
    //Construct person
    fn new(f: &str, l: &str) -> Person {
        Person {
            f_name: f.to_string(),
            l_name: l.to_string(),
        }
    }

    //Get Full name
    fn full_name(&self) -> String {
        format!("{} {}", self.f_name, self.l_name)
    }

    //Set Last name
    fn set_last_name(&mut self, l: &str) {
        self.l_name = l.to_string();
    }

    //Name to tuple
    fn to_tuple(self) -> (String, String) {
        (self.f_name, self.l_name)
    }
}

pub fn run() {

    //Struct 1
    let mut c = Color {
        red:255,
        blue: 0,
        green: 0, 
    };

    c.red = 200;
    println!("Color: {} {} {}", c.red, c.blue, c.green);

    //Struct 2
    let mut c2 = Color_tuple(100, 100, 100);
    println!("Color Tuple: {} {} {}", c2.0, c2.1, c2.2);

    //Struct 3
    let mut c3 = Person::new("Marry", "Doe");
    // println!("Person: {} {}", c3.f_name, c3.l_name );
    println!("Person: {}", c3.full_name());//Before last name change
    c3.set_last_name("Williams");
    println!("Person: {}", c3.full_name());//After last name change

    println!("Person: {:?}", c3.to_tuple());//After last name change
}