# golang驱动剖析_语法糖
> $sum  计算总和
> $avg  计算平均值
> $min  获取集合中所有文档对应值得最小值
> $max  获取集合中所有文档对应值得最大值。
> $push 在结果文档中插入值到一个数组中。
> $last 根据资源文档的排序获取最后一个文档数据
> $first  根据资源文档的排序获取第一个文档数据。
> $addToSet  在结果文档中插入值到一个数组中，但不创建副本。
> $project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
> $match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。
> $limit：用来限制MongoDB聚合管道返回的文档数。
> $skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
> $unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
> $group：将集合中的文档分组，可用于统计结果。
> $sort：将输入文档排序后输出。
> $geoNear：输出接近某一地理位置的有序文档。
> $gt  greater than >
> $gte gt equal  >=
> $lt  less than <
> $lte lt equal  <=
> $ne  not equal !=
> $eq  equal     =
> $ref：集合名称
> $id：引用的id
> $db:数据库名称，可选参数
* 原子操作常用命令
> $set   用来指定一个键并更新键值，若键不存在并创建。
> $unset 用来删除一个键。
> $inc   可以对文档的某个值为数字型（只能为满足要求的数字）的键进行增减的操作。
> $push  把value追加到field里面去，field一定要是数组类型才行，如果field不存在，会新增一个数组类型加进去。
> $pushAll 同$push,只是一次可以追加多个值到一个数组字段内。
> $pull  从数组内删除一个等于value值。
> $addToSet 增加一个值到数组内，而且只有当这个值不在数组内才增加。
> $pop 删除数组的第一个或最后一个元素
> $rename 修改字段名称
> $bit 位操作，integer类型 
> $regex 操作符来设置匹配字符串的正则表达式。不区分大小写的正则表达式 db.posts.find({post_text:{$regex:"runoob",$options:"$i"}})

## 聚合操作

```go
func (db *vpDB) vpOriginNumList(cond, having interface{}, pag *PagingBase) (orm.VpOriginNumInfoList, error) {
	pipeLine := []bson.M{
		{"$match": cond},
		{"$group": bson.M{
			"_id": bson.M{
				"store_id":          "$store_id",
				"origin_number":     "$origin_number",
				"target_speaker_id": "$target_speaker_id",
			},
			"store_id":          bson.M{"$first": "$store_id"},
			"origin_number":     bson.M{"$first": "$origin_number"},
			"target_speaker_id": bson.M{"$first": "$target_speaker_id"},
			"call_time":         bson.M{"$max": "$call_time"},
			"vp_count":          bson.M{"$sum": 1},
			"call_serial_list":  bson.M{"$addToSet": "$call_serial"},
			"algo_engine_list":  bson.M{"$first": "$algo_engine_list"},
		}},
		{"$match": having},
	}

	if pag != nil {
		if len(pag.SortFields) > 0 {
			sortFile := bson.M{}
			for _, item := range pag.SortFields {
				if item != "" {
					if item[0] == '-' {
						sortFile[item[1:]] = -1
					} else {
						sortFile[item] = 1
					}
				}
			}
			pipeLine = append(pipeLine, bson.M{"$sort": sortFile})
		}

		if pag.Limit > 0 && pag.Offset > 0 {
			pipeLine = append(pipeLine, bson.M{"$limit": pag.Limit},
				bson.M{"$skip": pag.Offset})
		}

	}
	list := orm.VpOriginNumInfoList{}
	err := db.vpColl.With(db.sess).Pipe(pipeLine).All(&list)
	return list, err
} 
```
## 管道

### 管道的概念
> 管道在Unix和Linux中一般用于将当前命令的输出结果作为下一个命令的参数。
> MongoDB的聚合管道将MongoDB文档在一个管道处理完毕后将结果传递给下一个管道处理。管道操作是可以重复的。
> 表达式：处理输入文档并输出。表达式是无状态的，只能用于计算当前聚合管道的文档，不能处理其它的文档。

