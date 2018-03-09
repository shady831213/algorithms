# 扩展与收缩均摊分析

------------------------------------

## 势能函数：
考虑扩展完的势能为0， 到0.75的capacity时为capacity
所以 ：
当size >= 3/8 capacity时： (8/3) * size - capacity
当size < 3/8 capacity时： (3/8)capacity - size

至分析扩容和收缩的情况

## 扩容
