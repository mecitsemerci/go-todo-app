package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectId primitive.ObjectID

func (o *ObjectId) String() string {
	return primitive.ObjectID(*o).Hex()
}

func (o *ObjectId) Set(str string) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil{
		*o = ObjectId(primitive.NilObjectID)
	} else {
		*o = ObjectId(oid)
	}
}
