Stoplight - Firewall Workflow in Go
======================================

![Alt text](https://github.com/f47h3r/stoplight/raw/master/release/images/firewall_req.png "Re-imagined Firewall Request App")

---------

![Alt text](https://github.com/f47h3r/stoplight/raw/master/release/images/firewall_request_makerequest.png "Re-imagined Firewall Request App")

---------

![Alt text](https://github.com/f47h3r/stoplight/raw/master/release/images/firewall_request_status.png "Re-imagined Firewall Request App")

---------

![Alt text](https://github.com/f47h3r/stoplight/raw/master/release/images/firewall_request_approval.png "Re-imagined Firewall Request App")

---------

![Alt text](https://github.com/f47h3r/stoplight/raw/master/release/images/firewall_request_audit.png "Re-imagined Firewall Request App")

---------

Dependencies:
```
brew install mongodb
brew install go
```

Howto Clone:
```
git clone git@github.com:f47h3r/stoplight.git
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

