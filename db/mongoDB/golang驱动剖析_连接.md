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



