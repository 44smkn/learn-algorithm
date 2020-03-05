pub struct SimpleStack {
    s: Vec<i32>,
    top: usize,
}

impl SimpleStack {
    pub fn initialize() -> Self {
        Self {
            s: vec![0; 5],
            top: 0,
        }
    }

    pub fn push(&mut self, num: i32) {
        if self.is_full() {
            panic!("overflow");
        }
        self.top += 1;
        self.s[self.top] = num;
    }

    fn is_full(&self) -> bool{
        self.top >= self.s.capacity()
    }

    pub fn pop(&mut self) -> i32 {
        if self.is_empty() {
            panic!("underflow");
        }
        self.top -= 1;
        self.s[self.top+1] 
    }

    fn is_empty(&self) -> bool {
        self.top == 0
    }
}