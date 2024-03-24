# algorithm-practice

此项目为Golang相关的算法、数据结构练习项目

## 1.排序算法

### 1.1 快速排序
#### 1.1.1 原数组地址改变
[答案](arithmetic/quick_sort.go)
#### 1.1.2 原地原数组地址不改变
[答案](arithmetic/quick_sort1.go)
### 1.2 插入排序
[答案](arithmetic/insert_sort.go)

### 1.3 归并排序
[答案](arithmetic/merge_sort.go)

### 1.4 选择排序
[答案](arithmetic/select_sort.go)

### 1.5 冒泡排序
[答案](arithmetic/bubble_sort.go)

## 2.查找算法

### 用递归和非递归实现二分法查找
[答案](arithmetic/binary_search.go)

## 3. 递归

### 3.1 杨辉三角根据某行某列求值
[答案](arithmetic/pascal_triangle.go)

### 3.2 菲波那切数列

1、1、2、3、5、8、13、21、34、……

给定一个数值，得到数据的内容

[答案](arithmetic/fibonacci.go)

## 4. 回溯

### 4.1 八皇后问题：我们有一个 8x8 的棋盘，希望往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子。求一共有几种摆放方法?并打印出来
[答案](arithmetic/eight_queens.go)

### 4.2 全排列问题：给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
[答案](arithmetic/permutation.go)

## 5. 链表

### 5.1 单链表反转

[答案](arithmetic/reverse_node.go)

## 6. 动态规划

### 6.1 挖金矿问题：你是金矿老板，你有20个矿工；当下有5做金矿，分别能采5吨、10吨、11吨、17吨金子，每个金矿所需人手5、13、11，17个人...求如何安排工人，才能取得最多的金矿。
要求写出通用的解决算法

[答案](arithmetic/permutation.go)

## 7. 其他

### 7.1 求一个数中所有的质数(prime)

ps:一个只能被1和自身整除的数称之为质数

[答案](arithmetic/prime.go)

### 7.2 利用redis写一个限速函数，规定1分钟内限访问次数1000次，若超出限制，则禁止访问10秒

[答案](arithmetic/limit_request.go)

### 7.3 给定一个int型数组array,求出从该数组中获取m(m>1)个数的所有集合

例如：array是{1,2,3,4,5,6,7,8,9}

### 7.4 将一个十进制数转化为二进制数并输出字符串

[答案](arithmetic/numeration.go)

### 7.5 两个协程交替打印

[答案](arithmetic/goroutine_exchange.go)







