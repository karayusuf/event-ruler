Golang implementation of https://github.com/aws/event-ruler

# Event Ruler

Event Ruler (called Ruler in rest of the doc for brevity) is a golang library that allows matching Rules to Events.
An event is a list of fields, which may be given as name/value pairs or as a JSON object.
A rule associates event field names with lists of possible values.
There are two reasons to use Ruler:

It's fast; the time it takes to match Events doesn't depend on the number of Rules.
Customers like the JSON "query language" for expressing rules.

# Ruler by Example

An Event is a JSON object. Here's an example:

```json
{
  "version": "0",
  "id": "ddddd4-aaaa-7777-4444-345dd43cc333",
  "detail-type": "EC2 Instance State-change Notification",
  "source": "aws.ec2",
  "account": "012345679012",
  "time": "2017-10-02T16:24:49Z",
  "region": "us-east-1",
  "resources": [
    "arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00000"
  ],
  "detail": {
    "c-count": 5,
    "d-count": 3,
    "x-limit": 301.8,
    "source-ip": "10.0.0.33",
    "instance-id": "i-000000aaaaaa00000",
    "state": "running"
  }
}
```

You can also see this as a set of name/value pairs.
For brevity, we present only a sampling.
Ruler has APIs for providing events both in JSON form and as name/value pairs:

```
    +--------------+------------------------------------------+
    | name         | value                                    |
    |--------------|------------------------------------------|
    | source       | "aws.ec2"                                |
    | detail-type  | "EC2 Instance State-change Notification" |
    | detail.state | "running"                                |
    +--------------+------------------------------------------+
```

Events in the JSON form may be provided in the form of a raw JSON String or byte slice.

## Simple matching

The rules in this section all match the sample event above:

{
  "detail-type": [ "EC2 Instance State-change Notification" ],
  "resources": [ "arn:aws:ec2:us-east-1:123456789012:instance/i-000000aaaaaa00000" ],
  "detail": {
    "state": [ "initializing", "running" ]
  }
}


This will match any event with the provided values for the resource, detail-type, and detail.state values, ignoring any other fields in the event. It would also match if the value of detail.state had been "initializing".

Values in rules are always provided as arrays, and match if the value in the event is one of the values provided in the array. The reference to resources shows that if the value in the event is also an array, the rule matches if the intersection between the event array and rule-array is non-empty.

