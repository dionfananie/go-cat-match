import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "4s", target: 5 },
    { duration: "4s", target: 400 },
    { duration: "5s", target: 200 },
    { duration: "1s", target: 20 },
  ],
};

function randomIntFromInterval(min, max) {
  // min and max included
  return Math.floor(Math.random() * (max - min + 1) + min);
}

export default function () {
  const rndInt = randomIntFromInterval(1, 3);
  const url = `http://localhost:8080/v1/cat/${rndInt}`;

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ3NzA2MDEsImlhdCI6MTcxNDc0MTgwMSwidXNlcklkIjoyfQ.ZqOfjpmV12id8Pr9UdyM04l2Mhtst5GSaZ_vxwVloVI",
    },
  };

  let bodyContent = JSON.stringify({
    name: "Cat 4",
    race: "British Shorthair",
    ageInMonth: "10",
    sex: "male",
    description: "cat dog",
    imageUrls: [
      "https://www.catster.com/wp-content/uploads/2023/11/british-shorthair_FotoMirta_Shutterstock.jpg",
      "haha.com",
      "hihi.com",
    ],
  });

  const res = http.put(url, bodyContent, params);

  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
