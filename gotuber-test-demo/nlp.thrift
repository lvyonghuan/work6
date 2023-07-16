namespace go nlp

struct Request{
1:required i64 choice
2:required string requestMessage
}

struct Response{
1:required string responseMessage
}

service nlp{
Response handel(1:Request req)
}