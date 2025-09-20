## es的向量查询
https://www.elastic.co/docs/solutions/search/vector/knn#approximate-knn-limitations

es8.0支持数据的向量存储和向量查询

数据类型：dense vector
查询算法：ANN  支持查询的是 KNN
```text
这里其实有两种向量查询的方式

方式一：（推荐） KNN
方式二： 基于全量查询的方式，script_scope 只适合于小量的数据集场景
    因为该方式是基于先全量（可条件）查询得到的结果集，然后将结果集作为向量再查询的，如果结果集大，那么就忙； 减少速度慢的方式，可以将查询结果集小化
```

相似度的算法： 
```text
cosine   default
l1_norm
l2_norm
```

es本身只负责向量的存储和向量的查询，不负责如何将数据向量化

### KNN的使用

#### filter场景
存在filter时，filter是post-filter，是在向量检索后再执行的filter -- 这样的结果是，可能得到的结果数量<= top k

#### combined query
存在query查询和向量查询时（可能是多个向量查询），需要进行多个向量查询拿到多个top结果，然后与query的结果进行联合，求top size 然后再返回出去





