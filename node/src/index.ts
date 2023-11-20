import "dotenv/config.js";
import express from "express";
import cors from "cors";

import crm from "./route/crm";
import pipefy from "./route/pipefy";

const app = express();
app.use(cors());

app.use("/crm", crm);
app.use("/pipefy", pipefy);

app.listen(3001);
