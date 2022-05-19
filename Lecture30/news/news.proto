syntax="proto3";
package news;

option go_package = "./news";

message Message{
   bytes Title =1;
}
message MessageClient{
   string Title =1;
}

service NewsService{
rpc SayHello(MessageClient) returns(Message){}
}