"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
require("dotenv/config.js");
const express_1 = __importDefault(require("express"));
const cors_1 = __importDefault(require("cors"));
const crm_1 = __importDefault(require("./route/crm"));
const pipefy_1 = __importDefault(require("./route/pipefy"));
const app = (0, express_1.default)();
app.use((0, cors_1.default)());
app.use("/crm", crm_1.default);
app.use("/pipefy", pipefy_1.default);
app.listen(3001);
