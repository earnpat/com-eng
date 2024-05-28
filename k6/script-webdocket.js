import { check } from "k6";
import ws from "k6/ws";

export default function () {
  const basePath = "ws://localhost:9002/ws";
  const params = {};

  const response = ws.connect(basePath, params, function (socket) {
    socket.on("open", () => console.log("connected"));
    socket.on("message", (data) => console.log("Message received: ", data));
    socket.on("close", () => console.log("disconnected"));
  });

  check(response, { "status is 101": (r) => r && r.status === 101 });
}
