"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const crm_js_1 = __importDefault(require("../controller/crm.js"));
const router = (0, express_1.Router)();
let apiStatus = false;
let timeout = 15000;
let lastIntervalId = 0;
const start = (_, res) => {
    apiStatus = true;
    lastIntervalId = setInterval(crm_js_1.default, timeout);
    res === null || res === void 0 ? void 0 : res.status(200).send();
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
router.patch("/change/timeout", (req, res) => {
    res.status(200).send();
});
start(undefined, undefined);
exports.default = router;
