Stoplight - Firewall Workflow in Go
======================================

![Alt text](http://github.gapinc.dev/security-engineering/stoplight/raw/master/firewall_req.png "Re-imagined Firewall Request App")

Dependencies:
```
brew install mongodb
brew install go
```

Howto Clone:
```
git clone git@github.gapinc.dev:security-engineering/stoplight.git
```

-------

Fire up Mongo:
```
mongod --config /usr/local/etc/mongod.conf
```

Run the Go App!:
```
go run app.go
```

You can also build a binary:
```
go build app.go
```

Onece you fire up the App - It can be Viewed in your local Browser: [http://localhost:3000/](http://localhost:3000/)

