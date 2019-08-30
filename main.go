package main

import (
	"gopkg.in/mgo.v2"
)

type book struct {
	Name string
	Price int
}
func main() {
	im := book{"im",1}
	imm := book{"imm",10}
	imm := book{"imm",10}

	sess, err := mgo.Dial("mongodb://mongo/ew")
	if err != nil {
		panic(err)
	}
	sess.SetMode(mgo.Monotonic, true)

	TagColl := sess.DB("").C("t")

	TagColl.With(sess).Insert(&imm)

}

