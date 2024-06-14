import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export const errorRate = new Rate("errors");

export const options = {
  // vus: 10,
  // vus: 40,
  // vus: 70,
  vus: 100,
  // duration: "1s",
  // duration: "10s",
  // duration: "30s",
  duration: "1m",
};

const roundNo = "048";
const refKey = `${options.vus}_${options.duration}_${roundNo}`;
const basePath = "http://localhost:9003";
// const url = `${basePath}/rest/${refKey}`;
const url = `${basePath}/grpc/${refKey}`;

export default function () {
  const response = http.get(url);
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );
}
