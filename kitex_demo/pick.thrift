namespace go api

struct Request{
1:required i64 count
2:required string name
3:required i32 age
4:optional string nickName
}

struct Response{
1:string indentity
}

service Pick{
Response pick(1:Request req)
}