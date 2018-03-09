# 扩展与收缩平摊分析

------------------------------------

## 势能函数：
考虑扩展完的势能为0， 到0.75的capacity时为capacity
所以 ：
当size >= 3/8 capacity时： (8/3) * size - capacity
当size < 3/8 capacity时： (3/8)capacity - size

至分析扩容和收缩的情况

## 扩容
操作消耗为size+1,size为插入前的大小，cap为插入前的容量，所以平摊代价为：

实际代价 + 当前势能 - 上次势能 = size +1 + ((8/3) * (size+1) - 2 * cap) - ((8/3) * size - cap)

                             = size + 1 + (8/3) * szie + 8/3 - 2cap - (8/3) * size + cap
                             
                             = 11/8 + size - cap = 11/8 + (3/4)cap -cap < 11/8
                             
O(1)

## 收缩
1/8容量时收缩，操作消耗为size (删除1， 拷贝size-1),size为删除前的大小，cap为删除前的容量

实际代价 + 当前势能 - 上次势能 = size + ((3/8 * 1/2)cap - (size-1)) - ((3/8)cap - size)

                             = size + (3/16)cap - size + 1 - (3/8)cap + size
                             
                             = size - (3/16)cap + 1 = 1 + (2/16) - (3/16)cap < 1
                             
O(1)
