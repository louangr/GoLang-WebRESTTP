package resources

var InternalErrorJson = "{ \"message\": \"An internal error occurred\", \"code\": 500 }"

var NotFoundResourceJson = "{ \"message\": \"The resource is not found\", \"code\": 404 }"

var MarshalingErrorJson = "{ \"message\": \"An error occurred while marshaling the request body in json format\", \"code\": 400 }"

var MarshalingError = "An error occurred while marshaling in json format: "

var SuccessfulAdditionJson = "{ \"message\": \"The resource was successfully added\", \"code\": 201 }"

var UnsuccessfulAdditionJson = "{ \"message\": \"The resource was not added\", \"code\": 400 }"

var SuccessfulUpdateJson = "{ \"message\": \"The resource was successfully updated\", \"code\": 200 }"

var UnsuccessfulUpdateJson = "{ \"message\": \"The resource was not updated\", \"code\": 400 }"

var SuccessfulDeletionJson = "{ \"message\": \"The resource was successfully deleted\", \"code\": 200 }"

var UnsuccessfulDeletionJson = "{ \"message\": \"The resource was not deleted\", \"code\": 400 }"
