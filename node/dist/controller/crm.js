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
const crm_js_1 = __importDefault(require("../external/crm.js"));
const crm_js_2 = __importDefault(require("../consumable/crm.js"));
const iteration = () => __awaiter(void 0, void 0, void 0, function* () {
    const crmChatTypes = ["ociosos", "heavy", "recentes"];
    for (let i = 0; i < crmChatTypes.length; i++) {
        try {
            const { data } = yield crm_js_1.default.get(`/${crmChatTypes[i]}`);
            for (let j = 0; j < data.length; j++) {
                try {
                    yield crm_js_1.default.post(`/${crmChatTypes[i]}/iniciarCrm/${data[j].id}`);
                }
                catch (err) {
                    console.log("Erro no início do CRM do usuário ", data[j].id);
                    console.log(err);
                }
                try {
                    (0, crm_js_2.default)(data[j].email);
                }
                catch (err) {
                    console.log("Erro no envio de email para o usuário ", data[j].id);
                    console.log(err);
                }
            }
        }
        catch (err) {
            console.log("Erro na iteração pelo CRM ", crmChatTypes[i]);
            console.log(err);
        }
    }
});
exports.default = iteration;
