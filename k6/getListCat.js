import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "10s", target: 30 },
    { duration: "4s", target: 9000 },
    { duration: "5s", target: 800 },
    { duration: "1s", target: 20 },
  ],
};

export default function () {
  const url = "http://localhost:8080/v1/cat";

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ3NzA2MDEsImlhdCI6MTcxNDc0MTgwMSwidXNlcklkIjoyfQ.ZqOfjpmV12id8Pr9UdyM04l2Mhtst5GSaZ_vxwVloVI",
    },
  };

  const res = http.get(url, params);

  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
