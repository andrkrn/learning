fn main() {
  use std::fmt;

  println!("hello {name}", name="andri");

  // https://doc.rust-lang.org/std/fmt/#formatting-traits
  println!("{} {:b}", 1, 3);
  println!("{:}", "andrkrn".to_ascii_uppercase());
  println!("My name is {0}, {1} {0}", "Bond", "James");

  #[derive(Debug)]
  struct Structure(i32);

  impl fmt::Display for Structure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "{}", self.0)
    }
  }

  let a = Structure(3);
  println!("Display: {}", a);
  println!("Debug: {:?}", a);

  #[derive(Debug)]
  struct Point2D(f32, f32);
  let p1 = Point2D(3.33, 4.44);

  impl fmt::Display for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "{}, {}", self.0, self.1)
    }
  }

  impl fmt::Binary for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "{:b}, {:b}", self.0.to_bits(), self.1.to_bits())
    }
  }

  println!("Display: {}", p1);
  println!("Display Binary: {:b}", p1);
  println!("Debug: {:?}", p1);



  struct List(Vec<i32>);

  impl fmt::Display for List {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      let vec = &self.0;

      write!(f, "[")?;
      for (c, v) in vec.iter().enumerate() {
        write!(f, " {}: {} ", c, v)?;
      }
      write!(f, "]") // no semicolon!!!
    }
  }
  let v = List(vec![1,2,3,4]);
  println!("Display: {}", v);
}
