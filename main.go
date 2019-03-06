package main


import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
//	"time"
	//"time"
//	"os"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Subject      string
	Score     int
//	Timestamp time.Time
}

var (
	IsDrop = true
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("testt").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("testt").C("user")

	// Index
	index := mgo.Index{
		Key:        []string{"subject", "score"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(&Person{Subject: "History", Score: 88},
		&Person{Subject: "History", Score: 92},
		&Person{Subject: "History", Score: 79})

	if err != nil {
		panic(err)
	}

	// Query One


	// Query All
	var results []Person
	err = c.Find(bson.M{"subject": "History"}).Sort("srore").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	count, err := c.Find(bson.M{"score": bson.M{
		"$gt": 80,
	}}).Count()

	fmt.Println(count)

	mattch, err := c.Find(bson.M{"score": bson.M{
		"$eq": 79,
	}}).Count()

	fmt.Println(mattch)


	pipe := c.Pipe([]bson.M{{"$match": bson.M{"score":79}}})
	resp := []bson.M{}
	err = pipe.All(&resp)
	if err != nil {
			}
	fmt.Println(resp)



	result := Person{}
	err = c.Find(bson.M{"subject": "History"}).Select(bson.M{"score": 79}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Score:", result)

	/*	var result []struct{
			Text string `bson:"text"`
			Otherfield string `bson:"otherfield"`

		err := c.Find(nil).Select(bson.M{"text": 1, "otherfield": 1}).All(&result)
		if err != nil {
			// handle error
		}
		for _, v := range result {
			fmt.Println(v.Text)
		}*/
/*	err = c.Remove(bson.M{"name": "Foo Bar"})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}*/

	//c.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"name": "new Name"}}
fmt.Println("asdad")
}