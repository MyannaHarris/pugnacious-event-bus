package globalVars

import (
	"fmt"
)

var MissingSubscriptionParamsErr = fmt.Errorf("Subscription Params are missing. Event Key and (SQS Queue or API URL) are required.")
var TooManySubscriptionParamsErr = fmt.Errorf("There are too many Subscription Params. Only one of SQS Queue and API URL are allowd, not both.")
var MissingEventParamsErr = fmt.Errorf("Event Params are missing. Event Key is required.")
