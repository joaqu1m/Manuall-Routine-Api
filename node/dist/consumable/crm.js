"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const nodemailer_1 = __importDefault(require("nodemailer"));
const { EMAIL_USER, EMAIL_PASS } = process.env;
console.log(EMAIL_USER, EMAIL_PASS);
const transporter = nodemailer_1.default.createTransport({
    service: "Outlook365",
    auth: {
        user: EMAIL_USER,
        pass: EMAIL_PASS,
    },
});
const enviarEmail = (email) => {
    transporter.sendMail({
        from: EMAIL_USER,
        to: email,
        subject: "MANUALL: Você recebeu uma mensagem!",
        text: "Nosso assistente virtual Manuel te enviou uma mensagem",
    }, (err, info) => {
        if (err) {
            console.log(err);
        }
        else {
            console.log("Email enviado: " + info.response);
        }
    });
};
exports.default = enviarEmail;
