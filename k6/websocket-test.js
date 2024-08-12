import { check } from "k6";
import ws from "k6/ws";

export const options = {
  stages: [
    { duration: "5s", target: 10 }, // Ramp-up to 10 users over 30 seconds
    { duration: "1s", target: 10 }, // Stay at 10 users for 1 minute
    { duration: "4s", target: 0 }, // Ramp-down to 0 users over 30 seconds
  ],
};

export default function () {
  const url = "ws://localhost:9003/ws/test";
  const params = { tags: { my_tag: "hello" } };

  const response = ws.connect(url, params, function (socket) {
    socket.on("open", function () {
      console.log("Connected");
      socket.send(
        JSON.stringify({ type: "greeting", message: "Hello, server!" })
      );

      // socket.setInterval(function () {
      //   socket.send(JSON.stringify({ type: "ping" }));
      // }, 1000);
    });

    socket.on("message", function (message) {
      console.log("Received message:", message);
      const data = JSON.parse(message);
      check(data, { "message is valid": (msg) => msg.type === "response" });
    });

    socket.on("close", function () {
      console.log("Disconnected");
    });

    socket.on("error", function (e) {
      console.log("Error:", e.error());
    });

    socket.setTimeout(function () {
      console.log("10 seconds passed, closing the socket");
      socket.close();
    }, 10000);
  });

  check(response, { "status is 101": (r) => r && r.status === 101 });
  // sleep(1);
}
