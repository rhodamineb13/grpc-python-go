import logging
import grpc
import sys
from proto import user_pb2, user_pb2_grpc
import warnings

warnings.filterwarnings('ignore')

def main():
    with grpc.insecure_channel('localhost:50051') as channel:
        rpc_input = input("Press 1 to register or 2 to login: ")
        stub = user_pb2_grpc.UserFeaturesStub(channel)
        match rpc_input:
            case "1":
                request = user_pb2.Person(name = "Suparjo", email = "odinthor@yahoo.com", username = "supparzo", date_of_birth = "2006-11-11", password = "halodek21231")
                try:
                    response = stub.RegisterUser(request)
                except grpc.RpcError as rpc_error:
                    print(rpc_error.details())
            case "2":                
                request = user_pb2.LoginRequest(username = "supparzo", password = "halodek21231")
                response = stub.Login(request)
                print(response)

            case _:        
                print("error: input must be 1 or 2")
if __name__ == "__main__":
    main()

