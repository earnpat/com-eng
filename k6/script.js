import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export const errorRate = new Rate("errors");

export const options = {
  vus: 1,
  // vus: 10,
  // vus: 40,
  // vus: 70,
  // vus: 100,
  duration: "1s",
  // duration: "10s",
  // duration: "30s",
  // duration: "1m",
};

// const roundNo = "048";
// const refKey = `${roundNo}_${options.vus}_${options.duration}`;
const basePath = "http://localhost:9000";
// const url = `${basePath}/rest/start`;
// const url = `${basePath}/rest/${refKey}`;
// const url = `${basePath}/grpc/start`;
// const url = `${basePath}/grpc/${refKey}`;
const url = `${basePath}/websocket/start`;
// const url = `${basePath}/websocket/${refKey}`;

export default function () {
  const response = http.get(url);
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );
}
