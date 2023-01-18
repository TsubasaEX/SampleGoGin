package database

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var MgoConnect *mgo.Collection

func MongoDB() {

	/*
		reference : https://github.com/globalsign/mgo
		reference : https://github.com/go-mgo/mgo
		reference :https://help.aliyun.com/document_detail/398673.html
		mgo is known to work well on (and has integration tests against) MongoDB v3.0, 3.2, 3.4 and 3.6.

		MongoDB 4.0 is currently experimental - we would happily accept PRs to help improve support!

		MongoDB 5.0 -> 对saslStart以及saslContinue命令的参数进行严格校验，无法兼容mgo。saslContinue只需要conversationId和payload参数，而mgo提供了一个多余的参数 mechanism，更多信息，请参见mgo。
	*/
	/*
		測試以下連線到MongoDB 6.0.3會連不到, err: no reachable servers
		連線到MongoDB 6.0.3會連不到, err: no reachable servers
		連線到MongoDB 3.6.23, PASS
	*/
	/* MongoDB 3.6.23
	1. Create folder /data/db
	2. Run Powershell as admin
	3. .\mongod.exe --dbpath ../data/db
	*/

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	MgoConnect = session.DB("Demo").C("TestCollection")
	fmt.Println(MgoConnect)
}
