# LeetCode solutions

![](https://i.imgur.com/zZnUaf9.png)
```
       1                10  11                20  21                30  31                40
 000   ▣ ▣ ▣ ▣ ▣ ▣ ▣ ▣ ▣ ▣   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚
 040   ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚   ▣ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 080   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ▣ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 120   ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 160   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚
 200   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 240   ⬚ ▣ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 280   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ▣ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 320   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 360   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 400   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 440   ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 480   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚           ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 520   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚   ▣ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚     ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 560   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 600   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 640   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ▣ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 680   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 720   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 760   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 800   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ▣ ⬚
 840   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 880   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚   ⬚   ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚ ⬚
 920   ⬚ ⬚ ⬚ ⬚
```

|  # | Title                                          | Difficulty | Solutions  |
|----|------------------------------------------------|------------|------------|
|  1 | Two sum                                        | Easy       | Go, Python |
|  2 | Add two numbers                                | Medium     | Go         |
|  3 | Longest substring without repeating characters | Medium     | Go         |
|  4 | Median of two sorted arrays                    | Hard       | Go         |
|  5 | Longest palindromic substring                  | Medium     | Go         |
|  6 | ZigZag conversion                              | Medium     | Go         |
|  7 | Reverse integer                                | Easy       | Go         |
|  8 | String to integer (aoi)                        | Medium     | Go         |
|  9 | Palindrome number                              | Easy       | Go         |
| 10 | Regular expression matching                    | Hard       | Go         |

## How to run code
Install [leetcode-cli](https://github.com/skygragon/leetcode-cli)
```bash
npm i -g leetcode-cli
```

```bash
# Run help
leetcode help

# Login with your leetcode account
leetcode user -l

# List all questions
leetcode list

# Show stat
leetcode stat -g
```


### Todo
- [ ] Get list of algorithms from [LeetCode API](https://leetcode.com/api/problems/algorithms/)
- [ ] Generate the README file from source code
- [ ] Write solutions in Rust, Erlang and Clojure
