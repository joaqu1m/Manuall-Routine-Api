import nodemailer from "nodemailer";

const { EMAIL_USER, EMAIL_PASS } = process.env;

const transporter = nodemailer.createTransport({
  service: "Outlook365",
  auth: {
    user: EMAIL_USER,
    pass: EMAIL_PASS,
  },
});

const enviarEmail = (email: string) => {
  transporter.sendMail(
    {
      from: EMAIL_USER,
      to: email,
      subject: "MANUALL: Você recebeu uma mensagem!",
      text: "Nosso assistente virtual Manuel te enviou uma mensagem",
    },
    (err) => {
      if (err) {
        console.log("Erro no envio de email para o usuário ", email);
        console.log(err);
      }
    }
  );
};

export default enviarEmail;
