# 关于Go中的 := 和 = 的区别
在go里面，我们声明一个变量都是需要以 var 或者 func 之类的开头的，但是在函数里面，为了简单，我们可以直接使用 := 来声明变量。但是需要注意的是，在全局的环境中，我们不能这么做，而且在函数中，我们不能重复声明同一个变量名。
这个时候，如果出现一下要赋值两个值的情况，其中一个还是已经声明的时候要怎么办？

答案如下：
```
func main () {
    k := 1 
    j,k := 2,3 
}
```