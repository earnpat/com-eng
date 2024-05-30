import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export const errorRate = new Rate("errors");

export const options = {
  vus: 30,
  duration: "30s",
};

const basePath = "http://localhost:9000";

export default function () {
  const response = http.get(basePath + "/timestamp");
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );
}

