### 猴子偷桃问题

问题描述：
有五只猴子摘了一些桃子，打算隔天早上起来分了吃。
晚上的时候，第一只猴子偷偷起来把桃子分成五堆，还多了一个，就把多了的那个吃掉，并拿走了一堆。
第二只猴子也偷偷起来将桃子分成了五堆，还是又多了一个，同样吃掉了这一颗桃子，并拿走了其中一堆。第三只、第四只、第五只猴子都做了同样的事情。请问这堆桃子最少有多少个？

解答：
每只猴子每次的动作都一样，因此这个问题是一个典型的递归问题，解答如下：

```
package main

import "fmt"

/*
 * 功能：根据初始桃子数计算能够完成分配任务的猴子数
 * 参数：peachNum-初始桃子数 monkeyNum-能够完成分配任务的猴子数
 * 返回值：能够完成分配任务的猴子数
 */
var targetMonkeyNum int = 5

func dfs(peachNum, monkeyNum int) int {
	if (peachNum-1)%targetMonkeyNum != 0 {
		return monkeyNum
	}
	peachNum = (peachNum - 1) / targetMonkeyNum * (targetMonkeyNum - 1)
	return dfs(peachNum, monkeyNum+1)
}

func main() {
	peachNum := 1
	monkeyNum := 0
	// 如果有targetMonkeyNum个猴子能完成分配任务，则满足题目要求
	for dfs(peachNum, monkeyNum) != targetMonkeyNum {
		peachNum++
	}
	fmt.Printf("ans=%d\n", peachNum)
}

```

验证：

(3121-1)/5=624
624*4=2496

(2496-1)/5=499
499*4=1996

(1996-1)/5=399
399*4=1596

(1596-1)/5=319
319*4=1276

(1276-1)/5=255

