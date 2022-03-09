/* 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

 

示例 1:

输入: [1,2,3,4,5]
输出: True
 

示例 2:

输入: [0,0,1,2,5]
输出: True */

func isStraight(nums []int) bool {
    //nums 转化为递增序列
    sort.Ints(nums)
    //记录每个数字之间差值的和
    sub:=0
    for i:=0;i<4;i++{
        if nums[i]==0{
         	//continue 不进行下面的代码，进入下一次循环
            continue
        }
        //如果扑克牌（非0数字）重复，不是顺子
        if nums[i]==nums[i+1]{
            return false
        }
        //差值记录
        sub+=nums[i+1]-nums[i]
    }
    //总的差值小于5（或sub<=4）则为顺子
    return sub<5
}