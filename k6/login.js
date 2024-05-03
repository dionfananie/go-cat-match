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
  const url = "http://localhost:8080/v1/user/login";

  let payload = JSON.stringify({
    email: "hellob@asas.com",
    password: "123456789",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(url, payload, params);

  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
