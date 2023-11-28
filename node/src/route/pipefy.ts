import { Router } from "express";
import controller from "../controller/pipefy.js";

const router = Router();

let apiStatus = false;
let timeout = 15000;
let lastIntervalId: string | number | NodeJS.Timeout | undefined = 0;

const start = (_: any, res: any) => {
  apiStatus = true;
  lastIntervalId = setInterval(controller, timeout);
  res?.status(200).send();
};

router.post("/on", start);

router.post("/off", (_, res) => {
  apiStatus = false;
  clearInterval(lastIntervalId);
  res.status(200).send();
});

router.get("/check", (_, res) => {
  res.status(200).send({ status: apiStatus, timeout });
});

router.patch("/change/timeout", (_, res) => {
  res.status(200).send();
});

export default router;
