import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [{ duration: "1s", target: 1 }],
};

export default function () {
  const url = `http://localhost:8080/v1/cat`;

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ4MDcwODQsImlhdCI6MTcxNDc3ODI4NCwidXNlcklkIjozfQ.E5d49Ipc5NnEGeFIdDTX4q9pBbVcJN1tWkfhTKNGCek",
    },
  };

  let bodyContent = JSON.stringify({
    name: "Cat 24",
    sex: "male",
    race: "Maine Coon",
    ageInMonth: "5",
    description: "readiiii",
    imageUrls: ["goole.com", "haha.com", "hihi.com"],
  });

  const res = http.post(url, bodyContent, params);
  console.log(res.body);
  check(res, { "status was 201": (r) => r.status == 201 });
  // check(res, { "payload has id": (r) => Boolean(r.data.id) });
  sleep(1);
}
