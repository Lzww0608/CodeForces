impl Solution {
    pub fn find_valid_pair(s: String) -> String {
        let bytes = s.as_bytes();
        let mut cnt = vec![0usize; 10];

        for &b in bytes {
            let x = (b - b'0') as usize;
            cnt[x] += 1;
        }

        for (i, w) in bytes.windows(2).enumerate() {
            let a = (w[0] - b'0') as usize;
            let b = (w[1] - b'0') as usize;
            if a != b && cnt[a] == a && cnt[b] == b {
                return s[i..i + 2].to_string()
            }
        }

        String::new()
    }
}