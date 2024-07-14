import { check } from "k6";
import http from "k6/http";
import { Rate, Trend } from "k6/metrics";

export const errorRate = new Rate("errors");

// export const options = {
//   vus: 10,
//   // vus: 40,
//   // vus: 70,
//   // vus: 100,
//   duration: "10s",
//   // duration: "30s",
//   // duration: "1m",
// };

// const vus = [10, 40, 70, 100];
// const durations = ["10s", "30s", "60s"];
const vus = [10, 40, 70, 100];
const durations = [10]; // s unit
let s = 0;

// const roundNo = "001";
// const refKey = `${roundNo}_${options.vus}_${options.duration}`;
const refKey = "start";

const keys = [];
vus.forEach((vu, idxV) => {
  durations.forEach((d, idxD) => {
    const duration = `${d}s`;
    const key = `_${String((idxV + 1) * (idxD + 1)).padStart(
      3,
      "0"
    )}_${vu}_${duration}`;

    keys.push({
      key,
      vus: vu,
      duration,
      startTime: `${s}s`,
    });

    s += d + 5;
  });
});

const scenarios = {};
keys.forEach((key) => {
  scenarios[key.key] = {
    executor: "constant-vus",
    vus: key.vus,
    duration: key.duration,
    startTime: key.startTime,
    env: { SCENARIO: key.key },
  };
});

export const options = {
  scenarios,
  summaryTrendStats: [
    "avg",
    "min",
    "med",
    "max",
    "p(90)",
    "p(95)",
    "p(99)",
    "count",
  ],
};

const basePath = "http://localhost:9000";
const url = `${basePath}/rest/${refKey}`;
// const url = `${basePath}/grpc/${refKey}`;
// const url = `${basePath}/websocket/${refKey}`;

let trends = {};
keys.forEach((key) => {
  trends[key.key] = new Trend(key.key);
});

export default function () {
  const scenario = __ENV.SCENARIO;
  let trend = null;

  keys.forEach((key) => {
    if (scenario === key.key) {
      trend = trends[key.key];
    }
  });

  const response = http.get(url);
  check(
    response,
    { "status is OK": (r) => r.status === 200 } || errorRate.add(1)
  );

  trend.add(response.timings.duration);
}
