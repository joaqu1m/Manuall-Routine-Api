import external from "../external/crm.js";

const iteration = async () => {
  const crmChatTypes = ["ociosos", "heavy", "recentes"];

  for (let i = 0; i < crmChatTypes.length; i++) {
    try {
      const { data } = await external.get(`/${crmChatTypes[i]}`);
      for (let j = 0; j < data.length; j++) {
        console.log(data[j]);
        try {
          await external.post(`/${crmChatTypes[i]}/iniciarCrm/${data[j].id}`);
        } catch (err) {
          console.log(err);
        }
      }
    } catch (err) {
      console.log(err);
    }
  }
};

export default iteration;
