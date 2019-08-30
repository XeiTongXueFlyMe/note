# golang驱动剖析_连接
## 正确打开mongoDB方式（分布式架构下）

```go
var TagColl *mgo.Collection

func mongoInit() {
    sess, err := mgo.Dial("mongodb://mongo/config_service")
    //err.....
    
    sess.SetMode(mgo.Monotonic, true)
    sess.SetPoolLimit(10)
    
    go func() {
        for {
            time.Sleep(10 * time.Second)
            sess.Refresh()
        }
    }()
    
    TagColl = sess.DB("").C("TAG_COLL")
    _, err := TagColl.With(db.sess).UpsertId("", "")
    //err.....
}

```

## 补充知识

session设置的模式分别为:

* Strong

session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。

* Monotonic

session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。

* Eventual
session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。session 的写操作总是向主服务器发起，但是可能使用不同的连接，也就是写操作也不一定有序。



