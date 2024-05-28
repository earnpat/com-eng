import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export const errorRate = new Rate("errors");

const basePath = "http://localhost:3000";

export default function () {
  const response = http.get(basePath + "/timestamp");
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );
}
