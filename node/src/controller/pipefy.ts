import consumable from "../consumable/pipefy.js";
import external from "../external/pipefy.js";
import { data } from "../mockData/pipefy.js";

const iteration = async () => {
  const prospects = [...(await getProspects(1)), ...(await getProspects(2))];

  try {
    external.post("", prospects);
  } catch (err) {
    console.log("Erro no post dos prospects");
    console.log(err);
  }
};

const getProspects = async (tipoUsuario: 1 | 2) => {
  const req = await consumable(tipoUsuario);
  if (!req) return;

  const insertions = req.data.data.allCards.edges.reduce((acc: any) => {
    const node: any = req.data.data.allCards.edges[0].node;
    const mockData: any = data;
    const phaseName:
      | "Lead"
      | "Oportunidade"
      | "Cliente Adquirido"
      | "Cliente Perdido" = node.current_phase.name;

    const insertion = {
      idCliente: node.id,
      status: {
        Lead: 1,
        Oportunidade: 2,
        "Cliente Adquirido": 3,
        "Cliente Perdido": 4,
      }[phaseName],
      tipoUsuario,
    };

    for (let i = 0; i < node.fields.length; i++) {
      const fieldName = node.fields[i];
      const fieldData = mockData[fieldName.name] || null;

      if (!fieldData) continue;

      const optsValue = fieldData.opts?.[fieldName.value];
      const formatValue = fieldData.format?.(fieldName.value);
      const fieldValue = fieldName.value;
      // @ts-ignore
      insertion[fieldData.dbCol] =
        optsValue === undefined
          ? formatValue === undefined
            ? fieldValue
            : formatValue
          : optsValue;
    }
    return [...acc, insertion];
  }, []);

  return insertions;
};

export default iteration;
