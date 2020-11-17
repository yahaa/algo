struct SubrectangleQueries {
    rectangle: Vec<Vec<i32>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl SubrectangleQueries {
    fn new(rectangle: Vec<Vec<i32>>) -> Self {
        return SubrectangleQueries {
            rectangle: rectangle,
        };
    }
    fn update_subrectangle(&mut self, row1: i32, col1: i32, row2: i32, col2: i32, new_value: i32) {
        for i in row1..=row2 {
            for j in col1..=col2 {
                self.rectangle[i as usize][j as usize] = new_value;
            }
        }
    }
    fn get_value(&self, row: i32, col: i32) -> i32 {
        self.rectangle[row as usize][col as usize]
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_subrectangle_queries() {
        let nums: Vec<Vec<i32>> = vec![vec![1, 2, 2, 3], vec![1, 2, 2, 3]];
        let mut rec = SubrectangleQueries::new(nums);
        assert_eq!(1, rec.get_value(0, 0));
        rec.update_subrectangle(0, 0, 0, 0, 100);
        assert_eq!(100, rec.get_value(0, 0));
    }
}
