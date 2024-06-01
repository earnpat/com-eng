import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export const errorRate = new Rate("errors");

export const options = {
  // vus: 10,
  // vus: 40,
  vus: 70,
  // vus: 100,
  // duration: "1m",
  // duration: "5m",
  duration: "15m",
  // duration: "20m",
  // duration: "30m",
};

const basePath = "http://localhost:9001";

export default function () {
  const response = http.get(basePath + "/timestamp");
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );
}
