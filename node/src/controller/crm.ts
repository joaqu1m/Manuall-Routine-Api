import consumable from "../consumable/crm.js";
import external from "../external/crm.js";

const iteration = async () => {
  const crmChatTypes = ["ociosos", "heavy", "recentes"];

  for (let i = 0; i < crmChatTypes.length; i++) {
    try {
      const { data } = await external.get(`/${crmChatTypes[i]}`);
      for (let j = 0; j < data.length; j++) {
        try {
          await external.post(`/${crmChatTypes[i]}/iniciarCrm/${data[j].id}`);
        } catch (err) {
          console.log("Erro no início do CRM do usuário ", data[j].id);
          console.log(err);
        }
        try {
          consumable(data[j].email);
        } catch (err) {
          console.log("Erro no envio de email para o usuário ", data[j].id);
          console.log(err);
        }
      }
    } catch (err: any) {
      if (err.response.status === 404) continue;

      console.log("Erro na iteração pelo CRM ", crmChatTypes[i]);
      console.log(err);
    }
  }
};

export default iteration;
