"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const pipefy_js_1 = __importDefault(require("../consumable/pipefy.js"));
const pipefy_js_2 = require("../mockData/pipefy.js");
const iteration = () => __awaiter(void 0, void 0, void 0, function* () {
    const req = yield (0, pipefy_js_1.default)(1);
    if (!req)
        return;
    const insertions = req.data.data.allCards.edges.reduce((acc) => {
        var _a, _b;
        const node = req.data.data.allCards.edges[0].node;
        const mockData = pipefy_js_2.data;
        const phaseName = node.current_phase.name;
        const insertion = {
            id_cliente: node.id,
            status: {
                Lead: 1,
                Oportunidade: 2,
                "Cliente Adquirido": 3,
                "Cliente Perdido": 4,
            }[phaseName],
        };
        for (let i = 0; i < node.fields.length; i++) {
            const fieldName = node.fields[i];
            const fieldData = mockData[fieldName.name] || null;
            if (!fieldData)
                continue;
            // @ts-ignore
            insertion[fieldData.dbCol] =
                ((_a = fieldData.opts) === null || _a === void 0 ? void 0 : _a[fieldName.value]) ||
                    ((_b = fieldData.format) === null || _b === void 0 ? void 0 : _b.call(fieldData, fieldName.value)) ||
                    fieldName.value;
        }
        return [...acc, insertion];
    }, []);
    console.log(insertions);
});
exports.default = iteration;
