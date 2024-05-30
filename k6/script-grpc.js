import { check } from "k6";
import { Rate } from "k6/metrics";
import grpc from "k6/net/grpc";

export const errorRate = new Rate("errors");

export const options = {
  vus: 50,
  duration: "30m",
};

const client = new grpc.Client();
client.load(null, "topic.proto");

export default () => {
  const basePath = "localhost:9002";
  client.connect(basePath, {
    plaintext: true,
  });

  const response = client.invoke("topic.TopicService/HelloTopic", {});
  check(
    response,
    { "status is OK": (r) => r.status === grpc.StatusOK } || errorRate.add(1)
  );

  console.log(JSON.stringify(response.message.message));

  client.close();
};
