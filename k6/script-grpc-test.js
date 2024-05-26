import { check } from "k6";
import grpc from "k6/net/grpc";

const client = new grpc.Client();
client.load(null, "hello.proto");

export const options = {
  vus: 30,
  // duration: "1m",
  duration: "5s",
};

export default () => {
  client.connect("grpcbin.test.k6.io:9001", {
    // plaintext: false
  });

  const response = client.invoke("hello.HelloService/SayHello", {
    greeting: "Bert",
  });
  check(
    response,
    { "status is OK": (r) => r.status === grpc.StatusOK } || errorRate.add(1)
  );

  console.log(JSON.stringify(response.message));

  client.close();
};
