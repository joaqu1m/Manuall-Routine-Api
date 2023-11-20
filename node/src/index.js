import "dotenv/config.js";

import crmConsumable from "./consumable/crm.js";
import pipefyConsumable from "./consumable/pipefy.js";
import crm from "./controller/crm.js";

crm()

console.log(process.env.ROUTINE_AUTH);
