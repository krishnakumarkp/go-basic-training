# Select
The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statements will be a channel operation. 


# Practical use of select

Let's assume we have a mission critical application and we need to return the output to the user as quickly as possible. The database for this application is replicated and stored in different servers across the world. Assume that the functions server1 and server2 are in fact communicating with 2 such servers. The response time of each server is dependant on the load on each and the network delay. We send the request to both the servers and then wait on the corresponding channels for the response using the select statement. The server which responds first is chosen by the select and the other response is ignored. This way we can send the same request to multiple servers and return the quickest response to the user.