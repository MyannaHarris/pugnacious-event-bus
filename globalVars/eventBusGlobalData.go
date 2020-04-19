package globalVars

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/sqs"

	"pugnacious-event-bus/models"
)

// TODO: Must be a better way to make the hostname a constant. Perhaps, the recommended use would
// be to just call Hostname() every time, but want to skip the error.
var SOURCE_HOSTNAME, _ = os.Hostname()

var sess = session.Must(session.NewSession())
var SqsClient = sqs.New(sess, aws.NewConfig().WithRegion("us-east-2"))

var SubscriptionsMap = make(map[string]models.Subscription)
var EventsMap = make(map[string]models.Event)
