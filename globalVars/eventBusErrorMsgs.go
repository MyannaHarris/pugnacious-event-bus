package globalVars

import (
	"fmt"
)

var MissingSubscriptionParamsErr = fmt.Errorf("Subscription Params are missing. Event Key and (SQS Queue or API URL) are required.")
var MissingEventParamsErr = fmt.Errorf("Event Params are missing. Event Key is required.")
