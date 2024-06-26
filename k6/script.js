import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "30s", target: 20 },
    { duration: "30s", target: 10 },
    { duration: "20s", target: 0 },
  ],
};

export default function () {
  const res = http.get("http://localhost:8080/");
  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
