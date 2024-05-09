# Dynamic-Programming Algorithms
动态规划主要分为两个核心部分，一是确定「DP 状态」，二是确定「DP 转移方程」。

### **DP 状态**

「DP 状态」的确定主要有两大原则：

1. 最优子结构
2. 无后效性

### **DP 转移方程**

有了「DP 状态」之后，我们只需要用「分类讨论」的思想来枚举所有小状态向大状态转移的可能性即可推出「DP 转移方程」。

### **动态规划问题类别**

讲述完 DP 问题的解题思路后，我们来大致列举一下 DP 问题的类别。

DP 问题主要分为两大类，第一大类是 DP 类型，第二大类是 DP 优化方法。

### 总结

最后我们来总结一下 DP 问题的解题思路：

- 确定「DP 状态」
    - 符合「最优子结构」原则：DP 状态最优值由更小规模的 DP 状态最优值推出
    - 符合「无后效性」原则：状态的得到方式，不会影响后续其它 DP 状态取值
- 确定「DP 转移方程」
    - 分类讨论，细心枚举
    

### 参考

https://www.zhihu.com/question/39948290/answer/1309260344
https://www.zhihu.com/question/291280715/answer/1007691283

书籍 http://web.mit.edu/dimitrib/www/dpchapter.pdf

### Climbing Strairs Problems
- [Leetcode 70. Climbing Stairs](climbing_stairs.go)
- [LCCI 08.01.三步问题](climbing_stairs_v2.go)

### Subsequence Problems
**Maximum Sum contiguous subsequence**
- [Leetcode 53. Maximum Subarray](maximum_subarray.go)
- [Leetcode 1800. Maximum Ascending Subarray Sum](max_ascending_sum.go)
- [Leetcode 918. Maximum Sum Circular Subarray](max_subarray_sum_circular.go)

**Palindromic Substring and Subsequence**
- [Leetcode 647. Palindromic Substrings](palindromic_substrings.go)
- [Leetcode 5. Longest Palindromic Substring](longest_palindromic_substring.go)
