# golang驱动剖析_索引

```go
    if err := storeColl.EnsureIndex(mgo.Index{
        Key:    []string{"store_name", "owner_id"},//联合索引
        Unique: true,
    }); err != nil {
        return err
    }
```

* mgo.Index剖析

```go
 type Index struct {
    Key        []string // Index key fields; prefix name with dash (-) for descending order

    //建立的索引是否唯一。指定为true创建唯一索引。默认值为false.
    Unique     bool     // Prevent two documents from having the same index key
    //3.0+版本已废弃。在建立唯一索引时是否删除重复记录,指定 true 创建唯一索引。默认值为 false.
    DropDups   bool     // Drop documents with the same index key as a previously indexed one

    //建索引过程会阻塞其它数据库操作，background可指定以后台方式创建索引，即增加 "background" 可选参数。 "background" 默认值为false。
    Background bool     // Build index in background and return immediately
    //对文档中不存在的字段数据不启用索引；这个参数需要特别注意，如果设置为true的话，在索引字段中不会查询出不包含对应字段的文档.。默认值为 false.
	Sparse     bool     // Only index documents containing the Key fields

    //指定一个以秒为单位的数值，完成 TTL设定，设定集合的生存时间。
	// If ExpireAfter is defined the server will periodically delete
	// documents with indexed time.Time older than the provided delta.
	ExpireAfter time.Duration

	// Name holds the stored index name. On creation if this field is unset it is
	// computed by EnsureIndex based on the index key.
	Name string

	// Properties for spatial indexes.
	//
	// Min and Max were improperly typed as int when they should have been
	// floats.  To preserve backwards compatibility they are still typed as
	// int and the following two fields enable reading and writing the same
	// fields as float numbers. In mgo.v3, these fields will be dropped and
	// Min/Max will become floats.
	Min, Max   int
	Minf, Maxf float64
	BucketSize float64
	Bits       int

    //对于文本索引，该参数决定了停用词及词干和词器的规则的列表。 默认为英语
	// Properties for text indexes.
    DefaultLanguage  string
    //对于文本索引，该参数指定了包含在文档中的字段名，语言覆盖默认的language，默认值为 language.
	LanguageOverride string

    //索引权重值，数值在 1 到 99,999 之间，表示该索引相对于其他索引字段的得分权重。
	// Weights defines the significance of provided fields relative to other
	// fields in a text index. The score for a given word in a document is derived
	// from the weighted sum of the frequency for each of the indexed fields in
	// that document. The default field weight is 1.
	Weights map[string]int

	// Collation defines the collation to use for the index.
	Collation *Collation
}
```
