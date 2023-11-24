import consumable from "../consumable/pipefy.js";
import { data } from "../mockData/pipefy.js";

const iteration = async () => {
  const req = await consumable(1);
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

      if (!fieldData) continue;
      // @ts-ignore
      insertion[fieldData.dbCol] =
        fieldData.opts?.[fieldName.value] ||
        fieldData.format?.(fieldName.value) ||
        fieldName.value;
    }
    return [...acc, insertion];
  }, []);

  return insertions;
};

export default iteration;
